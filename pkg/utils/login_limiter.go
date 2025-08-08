package utils

import (
	"sync"
	"time"
)

type LoginAttempt struct {
	Count      int
	LastFailed time.Time
	Blocked    bool
	BlockedAt  time.Time
}

var (
	loginAttempts = make(map[string]*LoginAttempt)
	mu            sync.Mutex
)

const (
	MaxLoginAttempts = 5
	BlockDuration    = 10 * time.Minute
)

// RegisterFailedLogin mencatat percobaan login gagal
func RegisterFailedLogin(identifier string) {
	mu.Lock()
	defer mu.Unlock()

	// identifier bisa berupa IP atau username/email
	la, exists := loginAttempts[identifier]
	if !exists {
		la = &LoginAttempt{}
		loginAttempts[identifier] = la
	}

	la.Count++
	la.LastFailed = time.Now()

	if la.Count >= MaxLoginAttempts {
		la.Blocked = true
		la.BlockedAt = time.Now()
	}
}

// ResetLoginAttempts menghapus catatan jika login sukses
func ResetLoginAttempts(identifier string) {
	mu.Lock()
	defer mu.Unlock()
	delete(loginAttempts, identifier)
}

// IsBlocked mengecek apakah identifier diblokir
func IsBlocked(identifier string) (bool, time.Duration) {
	mu.Lock()
	defer mu.Unlock()

	la, exists := loginAttempts[identifier]
	if !exists {
		return false, 0
	}

	if la.Blocked {
		if time.Since(la.BlockedAt) >= BlockDuration {
			// Reset setelah blokir berakhir
			delete(loginAttempts, identifier)
			return false, 0
		}
		// Hitung sisa waktu blokir
		remaining := BlockDuration - time.Since(la.BlockedAt)
		return true, remaining
	}

	return false, 0
}
