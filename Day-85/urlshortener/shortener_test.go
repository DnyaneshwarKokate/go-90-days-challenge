package urlshortener_test

import (
	"testing"

	"day85/urlshortener"
)

func TestBase62EncodingAndURLShortening(t *testing.T) {
	encoded := urlshortener.Base62Encode(125000)
	if encoded == "" {
		t.Fatalf("Base62 encoding produced empty string")
	}

	shortener := urlshortener.NewURLShortener(1000000)
	longURL := "https://github.com/Dnyanesh0902/go-90-days-challenge"

	shortKey := shortener.Shorten(longURL)
	if shortKey == "" {
		t.Fatalf("Shorten returned empty key")
	}

	// Resolve URL
	resolved, err := shortener.Resolve(shortKey)
	if err != nil || resolved != longURL {
		t.Fatalf("Expected %s, got %s (err: %v)", longURL, resolved, err)
	}

	if shortener.Clicks(shortKey) != 1 {
		t.Fatalf("Expected 1 click count, got %d", shortener.Clicks(shortKey))
	}
}
