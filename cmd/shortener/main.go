package main

import (
	"github.com/go-chi/chi/v5"
	chiMiddleware "github.com/go-chi/chi/v5/middleware"
	"net/http"
	urlhandler "url-shortner/internal/handler"
	"url-shortner/internal/middleware"
	"url-shortner/internal/repository"
	"url-shortner/internal/service"
)

func main() {
	repo := repository.NewUrlRepository()
	urlSvc := service.NewUrlService(repo)
	handler := urlhandler.NewUrlHandler(urlSvc)

	router := chi.NewRouter()
	router.Use(chiMiddleware.Logger)
	router.Use(chiMiddleware.Recoverer)
	router.Use(middleware.CheckBodyNotEmpty)

	handler.RegisterRoutes(router)

	err := http.ListenAndServe(":8080", router)
	if err != nil {
		return
	}
}
