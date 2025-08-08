package handler

import (
	"database/sql"
	"encoding/json"
	"net/http"
	"sync"
	"time"

	"animal-api/pkg/utils"

	"golang.org/x/crypto/bcrypt"
)

type LoginRequest struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

type LoginResponse struct {
	Token string `json:"token"`
}

type loginAttempt struct {
	Count     int
	BlockedAt time.Time
}

var (
	loginAttempts = make(map[string]*loginAttempt)
	mu            sync.Mutex
	maxAttempts   = 5
	blockDuration = 10 * time.Minute
)

// LoginHandler menerima koneksi DB untuk query user
func LoginHandler(db *sql.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req LoginRequest
		if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
			http.Error(w, "Invalid request body", http.StatusBadRequest)
			return
		}

		// === Rate limit check ===
		mu.Lock()
		attempt, exists := loginAttempts[req.Username]
		if !exists {
			attempt = &loginAttempt{}
			loginAttempts[req.Username] = attempt
		}

		if attempt.Count >= maxAttempts {
			if time.Since(attempt.BlockedAt) < blockDuration {
				mu.Unlock()
				http.Error(w, "Terlalu banyak nyoba boy, lain kali lagi yak, NT.", http.StatusTooManyRequests)
				return
			}
			attempt.Count = 0 // reset setelah blokir selesai
		}
		mu.Unlock()

		// === Ambil user dari DB ===
		var hashedPassword string
		err := db.QueryRow(`SELECT password FROM admin WHERE username = ?`, req.Username).Scan(&hashedPassword)
		if err == sql.ErrNoRows {
			// Username tidak ada → hitung sebagai gagal
			mu.Lock()
			attempt.Count++
			if attempt.Count >= maxAttempts {
				attempt.BlockedAt = time.Now()
			}
			mu.Unlock()
			http.Error(w, "Invalid credentials", http.StatusUnauthorized)
			return
		} else if err != nil {
			http.Error(w, "Database error", http.StatusInternalServerError)
			return
		}

		// === Cek password ===
		if bcrypt.CompareHashAndPassword([]byte(hashedPassword), []byte(req.Password)) != nil {
			mu.Lock()
			attempt.Count++
			if attempt.Count >= maxAttempts {
				attempt.BlockedAt = time.Now()
			}
			mu.Unlock()
			http.Error(w, "Waduh salah cuy", http.StatusUnauthorized)
			return
		}

		// === Login sukses: reset percobaan ===
		mu.Lock()
		attempt.Count = 0
		mu.Unlock()

		token, err := utils.GenerateToken(req.Username)
		if err != nil {
			http.Error(w, "Failed to generate token", http.StatusInternalServerError)
			return
		}

		resp := LoginResponse{Token: token}
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(resp)
	}
}
