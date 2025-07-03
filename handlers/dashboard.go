package handlers

import (
	"go-jwtx/internal/auth"
	"net/http"
)

func Dashboard(w http.ResponseWriter, r *http.Request) {
	data, ok := auth.GetClaimString(r.Context(), "username")
	if !ok {
		http.Error(w, "Context not found", http.StatusNotFound)
		return
	}

	w.Write([]byte("Welcome, "+data+" this is Dashboard"))
}