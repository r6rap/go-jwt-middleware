package utils

import (
	"log"
	"net/http"
	"strings"
)

func ExtractToken(r *http.Request) (string) {
	auth := r.Header.Get("Authorization")

	if !strings.HasPrefix(auth, "Bearer ") {
		log.Println("Token not found")
		return ""
	}

	token := strings.TrimPrefix(auth, "Bearer ")

	return  token
}