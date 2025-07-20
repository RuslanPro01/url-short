package service

import (
	"crypto/rand"
	"encoding/hex"
)

type UrlService struct {
}

func NewUrlService() *UrlService {
	return &UrlService{}
}

func (s *UrlService) CreateShortUrl(originalUrl string) (string, error) {
	shortCode, err := generateShortCode(8)
	if err != nil {
		return "", err
	}

	return "localhost/" + shortCode, nil
}

func generateShortCode(length int) (string, error) {
	bytes := make([]byte, length/2)
	if _, err := rand.Read(bytes); err != nil {
		return "", err
	}
	return hex.EncodeToString(bytes), nil
}
