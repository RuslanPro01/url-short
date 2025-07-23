package middleware

import (
	"bytes"
	"io"
	"net/http"
)

func CheckBodyNotEmpty(next http.Handler) http.Handler {
	return http.HandlerFunc(func(writer http.ResponseWriter, request *http.Request) {
		if !(request.Method == http.MethodPost || request.Method == http.MethodPut || request.Method == http.MethodPatch) {
			next.ServeHTTP(writer, request)
			return
		}

		body, err := io.ReadAll(request.Body)
		if err != nil {
			http.Error(writer, "Не получилось прочитать тело запроса", http.StatusInternalServerError)
			return
		}
		request.Body.Close()

		if len(body) == 0 {
			http.Error(writer, "Тело не может быть пустым", http.StatusBadRequest)
			return
		}

		request.Body = io.NopCloser(bytes.NewReader(body))
		next.ServeHTTP(writer, request)
	})
}
