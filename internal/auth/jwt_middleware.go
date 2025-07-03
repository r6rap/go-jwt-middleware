package auth

import (
	"context"
	"net/http"

	"github.com/r6rap/go-jwt-middleware/internal/utils"

	"github.com/golang-jwt/jwt/v5"
	"github.com/gorilla/mux"
)

func Middleware(jwtKey []byte) mux.MiddlewareFunc {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			tokenStr := utils.ExtractToken(r)

			token, err :=jwt.Parse(tokenStr, func(t *jwt.Token) (interface{}, error) {
				return jwtKey, nil
			})

			if err != nil || !token.Valid {
				http.Error(w, "Unauthoized: invalid token", http.StatusUnauthorized)
				return
			}

			claims, ok := token.Claims.(jwt.MapClaims)
			if !ok {
				http.Error(w, "Unauthorized: invalid claims", http.StatusUnauthorized)
				return 
			}

			ctx := context.WithValue(r.Context(), ContextUserKey, claims)

			next.ServeHTTP(w, r.WithContext(ctx))

		})
	}
}