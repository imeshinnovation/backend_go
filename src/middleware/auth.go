package middleware

import (
	"context"
	"net/http"
	"netix/src/auth"
	"strings"
)

type contextKey string

func AuthMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// Excluir rutas públicas (login, register, etc.)
		if r.URL.Path == "/login" || r.URL.Path == "/users" && r.Method == "POST" {
			next.ServeHTTP(w, r)
			return
		}

		// Obtener el token del header Authorization
		authHeader := r.Header.Get("Authorization")
		if authHeader == "" {
			http.Error(w, "Authorization header required", http.StatusUnauthorized)
			return
		}

		// El formato debe ser "Bearer <token>"
		splitToken := strings.Split(authHeader, "Bearer ")
		if len(splitToken) != 2 {
			http.Error(w, "Invalid token format", http.StatusUnauthorized)
			return
		}

		tokenString := splitToken[1]

		// Validar el token
		claims, err := auth.ValidateToken(tokenString)
		if err != nil {
			http.Error(w, "Invalid token: "+err.Error(), http.StatusUnauthorized)
			return
		}

		// Añadir información del usuario al contexto
		ctx := context.WithValue(r.Context(), contextKey("userID"), claims.UserID)
		ctx = context.WithValue(ctx, contextKey("userEmail"), claims.Email)

		next.ServeHTTP(w, r.WithContext(ctx))
	})
}
