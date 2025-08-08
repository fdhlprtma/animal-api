package middleware

import (
	"net/http"
	"strings"

	"animal-api/internal/config"
	"animal-api/pkg/utils"
)

func JWTAuth(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		env := config.Config.AppEnv

		// Mode development → semua bebas
		if env == "development" {
			next.ServeHTTP(w, r)
			return
		}

		// Mode strict → GET bebas, selain itu wajib token
		if env == "strict" && r.Method == http.MethodGet {
			next.ServeHTTP(w, r)
			return
		}

		// Mode super_strict → semua wajib token
		authHeader := r.Header.Get("Authorization")
		if authHeader == "" {
			http.Error(w, "Missing Authorization header", http.StatusUnauthorized)
			return
		}

		parts := strings.Split(authHeader, " ")
		if len(parts) != 2 || parts[0] != "Bearer" {
			http.Error(w, "Invalid Authorization header", http.StatusUnauthorized)
			return
		}

		if !utils.VerifyToken(parts[1]) {
			http.Error(w, "Invalid or expired token", http.StatusUnauthorized)
			return
		}

		next.ServeHTTP(w, r)
	})
}
