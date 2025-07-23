package service

import (
	"crypto/rand"
	"encoding/hex"
	"url-shortner/internal/repository"
)

type urlService struct {
	repo repository.URLRepository
}

func NewUrlService(repo repository.URLRepository) URLServiceProvider {
	return &urlService{
		repo: repo,
	}
}

func (s *urlService) CreateShortUrl(originalUrl string) (string, error) {
	shortCode, err := generateShortCode(8)
	s.repo.PushShortedCode(originalUrl, shortCode)
	if err != nil {
		return "", err
	}

	return "localhost/" + shortCode, nil
}

func (s *urlService) GetOriginalUrlByShortCode(shortCode string) (string, bool) {
	originalUrl, isSuccess := s.repo.GetOriginalUrl(shortCode)
	return originalUrl, isSuccess
}

func generateShortCode(length int) (string, error) {
	bytes := make([]byte, length/2)
	if _, err := rand.Read(bytes); err != nil {
		return "", err
	}
	return hex.EncodeToString(bytes), nil
}
