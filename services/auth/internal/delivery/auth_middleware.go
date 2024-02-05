package middleware

import (
	"encoding/json"
	"log"
	"net/http"
	"strings"

	"github.com/adityasuryadi/go-shop/pkg/jwt"
	"github.com/adityasuryadi/go-shop/services/auth/internal/model"
)

func Verify(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		tokenString := r.Header.Get("Authorization")
		if len(tokenString) == 0 {
			w.Header().Set("Content-Type", "application/json")
			w.WriteHeader(http.StatusUnauthorized)
			json.NewEncoder(w).Encode(model.ErrorResponse[string]{Code: http.StatusUnauthorized, Status: "UNAUTHORIZE", Error: "Missing Authorization Header"})
			return
		}
		tokenString = strings.Replace(tokenString, "Bearer ", "", 1)
		claims, err := jwt.DecodeTokenString(tokenString)
		if err != nil {
			w.Header().Set("Content-Type", "application/json")
			w.WriteHeader(http.StatusUnauthorized)
			log.Println(claims)
			json.NewEncoder(w).Encode(model.ErrorResponse[string]{Code: http.StatusUnauthorized, Status: "UNAUTHORIZE", Error: "UNAUTHORIZE"})
			return
		}
		next.ServeHTTP(w, r)
	})
}
