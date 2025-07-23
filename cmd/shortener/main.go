package main

import (
	"fmt"
	"github.com/go-chi/chi/v5"
	chiMiddleware "github.com/go-chi/chi/v5/middleware"
	"log"
	"net/http"
	"url-shortner/config"
	urlhandler "url-shortner/internal/handler"
	"url-shortner/internal/middleware"
	"url-shortner/internal/repository"
	"url-shortner/internal/service"
)

func main() {
	cfg, err := config.ParseFlags()
	if err != nil {
		log.Fatalf("Ошибка при чтении конфигурации: %v", err)
	}

	repo := repository.NewUrlRepository()
	urlSvc := service.NewUrlService(repo, cfg)
	handler := urlhandler.NewUrlHandler(urlSvc)

	router := chi.NewRouter()
	router.Use(chiMiddleware.Logger)
	router.Use(chiMiddleware.Recoverer)
	router.Use(middleware.CheckBodyNotEmpty)

	handler.RegisterRoutes(router)

	fmt.Printf("Сервер запускается по адресу: %s\n", cfg.ServerAddress)
	fmt.Printf("Базовый URL для ссылок: %s\n", cfg.BaseUrl)

	startErr := http.ListenAndServe(cfg.ServerAddress, router)
	if startErr != nil {
		return
	}
}
