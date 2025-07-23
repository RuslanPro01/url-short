package handler

import (
	"github.com/go-chi/chi/v5"
	"github.com/go-playground/validator/v10"
	"io"
	"mime"
	"net/http"
	"url-shortner/internal/service"
)

type UrlHandler struct {
	urlService service.URLServiceProvider
}

var validate *validator.Validate

func init() {
	validate = validator.New()
}

func NewUrlHandler(service service.URLServiceProvider) *UrlHandler {
	return &UrlHandler{
		urlService: service,
	}
}

func (handler *UrlHandler) RegisterRoutes(router *chi.Mux) {
	router.Post("/", handler.PostUrl)
	router.Get("/{shortCode}", handler.GetOrigin)
}

func (handler *UrlHandler) PostUrl(writer http.ResponseWriter, request *http.Request) {
	if request.Method != http.MethodPost {
		writer.WriteHeader(http.StatusMethodNotAllowed)
		return
	}

	contentType := request.Header.Get("Content-Type")
	mediaType, _, err := mime.ParseMediaType(contentType)
	if err != nil {
		http.Error(writer, "Некорректный заголовок Content-Type", http.StatusBadRequest)
		return
	}

	if mediaType != "text/plain" {
		http.Error(writer, "Content-Type должен быть text/plain", http.StatusBadRequest)
		return
	}

	body, err := io.ReadAll(request.Body)
	if err != nil {
		http.Error(writer, "Не удалось прочитать тело запроса", http.StatusInternalServerError)
		return
	}

	url := string(body)

	errs := validate.Var(url, "required,url")
	if errs != nil {
		http.Error(writer, "Передан некорректный url для сокращения", http.StatusBadRequest)
		return
	}

	shortUrl, err := handler.urlService.CreateShortUrl(url)
	if err != nil {
		http.Error(writer, "Внутренняя ошибка сервера при создании URL", http.StatusInternalServerError)
		return
	}

	writer.Header().Set("Content-Type", "text/plain")
	writer.WriteHeader(http.StatusCreated)
	io.WriteString(writer, shortUrl)
}

func (handler *UrlHandler) GetOrigin(writer http.ResponseWriter, request *http.Request) {
	if request.Method != http.MethodGet {
		writer.WriteHeader(http.StatusMethodNotAllowed)
		return
	}
	shortCode := chi.URLParam(request, "shortCode")

	origin, isSuccess := handler.urlService.GetOriginalUrlByShortCode(shortCode)

	if !isSuccess {
		writer.WriteHeader(http.StatusNotFound)
		return
	}
	writer.Header().Set("Location", origin)
	writer.WriteHeader(http.StatusTemporaryRedirect)
}
