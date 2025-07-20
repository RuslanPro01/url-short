package main

import (
	"github.com/go-chi/chi/v5"
	chiMiddleware "github.com/go-chi/chi/v5/middleware"
	"net/http"
	urlhandler "url-shortner/internal/handler"
	"url-shortner/internal/middleware"
)

func main() {
	router := chi.NewRouter()

	handler := urlhandler.NewUrlHandler()

	router.Use(chiMiddleware.Logger)
	router.Use(chiMiddleware.Recoverer)
	router.Use(middleware.CheckBodyNotEmpty)

	handler.RegisterRoutes(router)

	err := http.ListenAndServe(":8080", router)
	if err != nil {
		return
	}
}
