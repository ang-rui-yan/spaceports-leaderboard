package middleware

import (
	"net/http"
	"os"
	"strings"

	"github.com/dgrijalva/jwt-go"
)

// JWTMiddleware is a middleware function to check the JWT token
func JWTMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// Get the Authorization header from the request
		authHeader := r.Header.Get("Authorization")

		// Check if the Authorization header is present
		if authHeader == "" {
			http.Error(w, "Unauthorized", http.StatusUnauthorized)
			return
		}

		// Extract the token from the Authorization header
		tokenString := strings.Split(authHeader, " ")[1]

		// Parse the JWT token
		token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
			return []byte(os.Getenv("SECRET_KEY")), nil
		})

		// Check if there was an error parsing the token
		if err != nil {
			http.Error(w, "Unauthorized", http.StatusUnauthorized)
			return
		}

		// Check if the token is valid
		if !token.Valid {
			http.Error(w, "Unauthorized", http.StatusUnauthorized)
			return
		}

		// Access the claims of the token
		claims, ok := token.Claims.(jwt.MapClaims)
		if !ok {
			http.Error(w, "Unauthorized", http.StatusUnauthorized)
			return
		}

		// Access the specific claim value
		tokenInJWT := claims["token"].(string)
		if tokenInJWT != os.Getenv("TOKEN") {
			http.Error(w, "Unauthorized", http.StatusUnauthorized)
			return
		}

		next.ServeHTTP(w, r)
	})
}