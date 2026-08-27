package urlshortener

import (
	"errors"
	"sync"
	"sync/atomic"
)

const base62Alphabet = "0123456789abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"

// Base62Encode converts an integer ID into a Base62 string key.
func Base62Encode(number uint64) string {
	if number == 0 {
		return "0"
	}

	result := ""
	base := uint64(len(base62Alphabet))

	for number > 0 {
		remainder := number % base
		result = string(base62Alphabet[remainder]) + result
		number /= base
	}
	return result
}

type ShortURLRecord struct {
	ShortKey string
	LongURL  string
	Clicks   int64
}

// URLShortener engine with cache and analytics tracking.
type URLShortener struct {
	mu         sync.RWMutex
	autoID     uint64
	urlMap     map[string]*ShortURLRecord // ShortKey -> Record
	longToKey  map[string]string         // LongURL -> ShortKey
}

// NewURLShortener initializes URL shortener service.
func NewURLShortener(startID uint64) *URLShortener {
	return &URLShortener{
		autoID:    startID,
		urlMap:    make(map[string]*ShortURLRecord),
		longToKey: make(map[string]string),
	}
}

// Shorten maps a long URL to a unique Base62 short key.
func (s *URLShortener) Shorten(longURL string) string {
	s.mu.Lock()
	defer s.mu.Unlock()

	// Return cached short key if long URL already exists
	if existingKey, exists := s.longToKey[longURL]; exists {
		return existingKey
	}

	s.autoID++
	shortKey := Base62Encode(s.autoID)

	rec := &ShortURLRecord{
		ShortKey: shortKey,
		LongURL:  longURL,
		Clicks:   0,
	}

	s.urlMap[shortKey] = rec
	s.longToKey[longURL] = shortKey
	return shortKey
}

// Resolve looks up the original long URL and increments click analytics.
func (s *URLShortener) Resolve(shortKey string) (string, error) {
	s.mu.RLock()
	rec, exists := s.urlMap[shortKey]
	s.mu.RUnlock()

	if !exists {
		return "", errors.New("short URL key not found")
	}

	atomic.AddInt64(&rec.Clicks, 1)
	return rec.LongURL, nil
}

// Clicks returns click analytics count for a short key.
func (s *URLShortener) Clicks(shortKey string) int64 {
	s.mu.RLock()
	defer s.mu.RUnlock()
	if rec, exists := s.urlMap[shortKey]; exists {
		return atomic.LoadInt64(&rec.Clicks)
	}
	return 0
}
