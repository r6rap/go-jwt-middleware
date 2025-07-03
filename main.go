package main

import (
	"go-jwtx/routes"
	"log"
	"net/http"
)

func main() {
	r := routes.InitRoute()

	log.Println("Server running at :8080")
	http.ListenAndServe(":8080", r)
}