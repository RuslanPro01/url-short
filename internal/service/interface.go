package service

type URLServiceProvider interface {
	CreateShortUrl(url string) (string, error)
	GetOriginalUrlByShortCode(shortCode string) (string, bool)
}
