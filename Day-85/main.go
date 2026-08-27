package main

import (
	"fmt"

	"day85/urlshortener"
)

func main() {
	fmt.Println("=== Day 85: Distributed URL Shortener System Design (Bit.ly Architecture) ===")

	shortener := urlshortener.NewURLShortener(100000000)
	longURL := "https://golang.org/doc/effective_go#concurrency"

	fmt.Println("\n--- 1. Shortening Target Long URL ---")
	shortKey := shortener.Shorten(longURL)
	shortURL := fmt.Sprintf("https://short.link/%s", shortKey)
	fmt.Printf("Long URL:  %s\nShort URL: %s\n", longURL, shortURL)

	fmt.Println("\n--- 2. Simulating User Redirections ---")
	for i := 1; i <= 3; i++ {
		target, _ := shortener.Resolve(shortKey)
		fmt.Printf("Redirect #%d -> Destination: %s\n", i, target)
	}

	fmt.Printf("\n--- URL Analytics ---\nTotal Click Redirections for '%s': %d\n",
		shortKey, shortener.Clicks(shortKey))
}
