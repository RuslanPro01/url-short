package repository

type UrlRepository struct {
}

var shoredUrlMap = make(map[string]string, 10)

func NewUrlRepository() *UrlRepository {
	return &UrlRepository{}
}

func (repository *UrlRepository) PushShortedCode(originalUrl string, code string) {
	shoredUrlMap[code] = originalUrl
	return
}

func (repository *UrlRepository) GetOriginalUrl(code string) (string, bool) {
	shortUrl, ok := shoredUrlMap[code]
	if !ok {
		return "", false
	}
	return shortUrl, true
}
