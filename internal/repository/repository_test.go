package repository

import "testing"

func TestUrlRepository(t *testing.T) {
	repo := NewUrlRepository()
	originalURL := "https://practicum.yandex.ru"
	code := "testcode"

	_, found := repo.GetOriginalUrl(code)
	if found {
		t.Errorf("found url for code %s before it was added", code)
	}

	repo.PushShortedCode(originalURL, code)

	retrievedURL, found := repo.GetOriginalUrl(code)
	if !found {
		t.Errorf("could not find url for code %s", code)
	}

	if retrievedURL != originalURL {
		t.Errorf("retrieved url mismatch: got %s, want %s", retrievedURL, originalURL)
	}
}
