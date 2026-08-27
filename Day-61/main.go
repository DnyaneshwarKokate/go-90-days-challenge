package main

import (
	"log"
	"time"

	"day61/cache"
)

func main() {
	distCache := cache.NewDistributedCache()
	log.Println("Starting Day 61 Distributed Caching & Singleflight Demo...")

	distCache.Set("user:101", `{"name":"Alice","role":"admin"}`, 10*time.Second)

	val, err := distCache.Get("user:101")
	if err != nil {
		log.Fatalf("Cache get failed: %v", err)
	}

	log.Printf("Cache Hit for user:101 -> %s", val)
}
