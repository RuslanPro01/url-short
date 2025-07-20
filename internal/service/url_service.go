package service

import (
	"crypto/rand"
	"encoding/hex"
	"url-shortner/internal/repository"
)

type UrlService struct {
	urlRepository *repository.UrlRepository
}

func NewUrlService() *UrlService {
	return &UrlService{
		urlRepository: repository.NewUrlRepository(),
	}
}

func (s *UrlService) CreateShortUrl(originalUrl string) (string, error) {
	shortCode, err := generateShortCode(8)
	s.urlRepository.PushShortedCode(originalUrl, shortCode)
	if err != nil {
		return "", err
	}

	return "localhost/" + shortCode, nil
}

func (s *UrlService) GetOriginalUrlByShortCode(shortCode string) (string, bool) {
	originalUrl, isSuccess := s.urlRepository.GetOriginalUrl(shortCode)
	return originalUrl, isSuccess
}

func generateShortCode(length int) (string, error) {
	bytes := make([]byte, length/2)
	if _, err := rand.Read(bytes); err != nil {
		return "", err
	}
	return hex.EncodeToString(bytes), nil
}
