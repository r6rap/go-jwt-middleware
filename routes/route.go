package routes

import (
	"go-jwtx/config"
	"go-jwtx/handlers"
	"go-jwtx/internal/auth"

	"github.com/gorilla/mux"
)

func InitRoute() *mux.Router {
	r := mux.NewRouter()

	api := r.PathPrefix("/api").Subrouter()

	api.Use(auth.Middleware(config.JWTKey))

	api.HandleFunc("/dashboard", handlers.Dashboard)

	return r

}