package router

import (
	"go-postgres-redis-url-shortener/handlers"
	"net/http"
)

func UrlRouter(mux *http.ServeMux) {
	mux.HandleFunc("GET /{code}", handlers.GetFullUrl)
	mux.HandleFunc("POST /create", handlers.CreateShortUrl)
}
