package repository

type UrlRepository struct {
	storage map[string]string
}

func NewUrlRepository() URLRepository {
	return &UrlRepository{
		storage: make(map[string]string, 10),
	}
}

func (repository *UrlRepository) PushShortedCode(originalUrl string, code string) {
	repository.storage[code] = originalUrl
	return
}

func (repository *UrlRepository) GetOriginalUrl(code string) (string, bool) {
	shortUrl, ok := repository.storage[code]
	if !ok {
		return "", false
	}
	return shortUrl, true
}
