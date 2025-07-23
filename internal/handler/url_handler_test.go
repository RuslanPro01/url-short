package handler

import (
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
)

type mockUrlService struct {
	CreateShortUrlFunc            func(url string) (string, error)
	GetOriginalUrlByShortCodeFunc func(shortCode string) (string, bool)
}

func (m *mockUrlService) CreateShortUrl(url string) (string, error) {
	if m.CreateShortUrlFunc != nil {
		return m.CreateShortUrlFunc(url)
	}
	return "http://localhost:8080/mockurl", nil
}

func (m *mockUrlService) GetOriginalUrlByShortCode(shortCode string) (string, bool) {
	if m.GetOriginalUrlByShortCodeFunc != nil {
		return m.GetOriginalUrlByShortCodeFunc(shortCode)
	}
	return "", false
}

func TestUrlHandler_PostUrl(t *testing.T) {
	mockService := &mockUrlService{}
	handler := NewUrlHandler(mockService)

	testCases := []struct {
		name           string
		body           string
		wantStatusCode int
		wantBody       string
	}{
		{
			name:           "Valid request",
			body:           "https://example.com",
			wantStatusCode: http.StatusCreated,
			wantBody:       "http://localhost:8080/mockurl",
		},
		{
			name:           "Invalid URL in body",
			body:           "not-a-valid-url",
			wantStatusCode: http.StatusBadRequest,
			wantBody:       "Передан некорректный url для сокращения\n",
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			req := httptest.NewRequest(http.MethodPost, "/", strings.NewReader(tc.body))
			req.Header.Set("Content-Type", "text/plain")

			rr := httptest.NewRecorder()

			handler.PostUrl(rr, req)

			if rr.Code != tc.wantStatusCode {
				t.Errorf("wrong status code: got %d, want %d", rr.Code, tc.wantStatusCode)
			}

			if rr.Body.String() != tc.wantBody {
				t.Errorf("wrong response body: got %q, want %q", rr.Body.String(), tc.wantBody)
			}
		})
	}
}
