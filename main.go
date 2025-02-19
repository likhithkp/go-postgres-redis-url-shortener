package main

import (
	"go-postgres-redis-url-shortener/config"
	"go-postgres-redis-url-shortener/router"
	"log"
	"net/http"

	"github.com/joho/godotenv"
)

func main() {
	mux := http.NewServeMux()

	err := godotenv.Load()
	if err != nil {
		log.Println("Warning: No .env file found")
	}

	config.ConnectDB()
	defer config.CloseDB()

	router.UrlRouter(mux)
	http.ListenAndServe(":3000", mux)
}
