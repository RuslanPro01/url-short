package handler

import (
	"github.com/go-chi/chi/v5"
	"github.com/go-playground/validator/v10"
	"io"
	"net/http"
)

type UrlHandler struct {
}

var validate *validator.Validate

func init() {
	validate = validator.New()
}

func NewUrlHandler() *UrlHandler {
	return &UrlHandler{}
}

func (handler *UrlHandler) RegisterRoutes(router *chi.Mux) {
	router.Post("/", handler.PostUrl)
}

func (handler *UrlHandler) PostUrl(writer http.ResponseWriter, request *http.Request) {
	if request.Method != http.MethodPost {
		writer.WriteHeader(http.StatusMethodNotAllowed)
		return
	}

	if request.Header.Get("Content-Type") != "text/plain" {
		http.Error(writer, "Нужен url для сокращения в формате text/plain", http.StatusBadRequest)
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

	writer.Header().Set("Content-Type", "text/plain")
	writer.WriteHeader(http.StatusCreated)
	io.WriteString(writer, "localhost/testShortedUrl")
}
