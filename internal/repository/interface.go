package repository

type URLRepository interface {
	PushShortedCode(originalUrl string, code string)
	GetOriginalUrl(code string) (string, bool)
}
