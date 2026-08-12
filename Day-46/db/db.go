package db

import (
	"context"
	"database/sql"
	"fmt"
	"log"
	"time"

	"github.com/dnyaneshwarkokate/go-90-days-challenge/Day-46/config"
	_ "github.com/lib/pq"
	"github.com/redis/go-redis/v9"
)

// InitPostgres attempts to connect to PostgreSQL with retries to handle container startup lag.
func InitPostgres(cfg *config.Config) (*sql.DB, error) {
	dsn := cfg.GetDSN()
	var db *sql.DB
	var err error

	maxRetries := 10
	for i := 1; i <= maxRetries; i++ {
		db, err = sql.Open("postgres", dsn)
		if err == nil {
			err = db.Ping()
			if err == nil {
				log.Printf("Successfully connected to PostgreSQL at %s:%s (attempt %d/%d)", cfg.DBHost, cfg.DBPort, i, maxRetries)
				db.SetMaxOpenConns(25)
				db.SetMaxIdleConns(5)
				db.SetConnMaxLifetime(5 * time.Minute)
				return db, nil
			}
		}
		log.Printf("PostgreSQL not ready yet (attempt %d/%d): %v. Retrying in 2 seconds...", i, maxRetries, err)
		time.Sleep(2 * time.Second)
	}

	return nil, fmt.Errorf("failed to connect to PostgreSQL after %d attempts: %w", maxRetries, err)
}

// InitRedis attempts to connect to Redis with retries.
func InitRedis(cfg *config.Config) (*redis.Client, error) {
	rdb := redis.NewClient(&redis.Options{
		Addr:     cfg.GetRedisAddr(),
		Password: cfg.RedisPass,
		DB:       0,
	})

	maxRetries := 10
	ctx, cancel := context.WithTimeout(context.Background(), 2*time.Second)
	defer cancel()

	for i := 1; i <= maxRetries; i++ {
		pingCtx, pingCancel := context.WithTimeout(context.Background(), 2*time.Second)
		_, err := rdb.Ping(pingCtx).Result()
		pingCancel()
		if err == nil {
			log.Printf("Successfully connected to Redis at %s (attempt %d/%d)", cfg.GetRedisAddr(), i, maxRetries)
			return rdb, nil
		}
		log.Printf("Redis not ready yet (attempt %d/%d): %v. Retrying in 2 seconds...", i, maxRetries, err)
		time.Sleep(2 * time.Second)
	}

	_ = ctx
	return nil, fmt.Errorf("failed to connect to Redis after %d attempts", maxRetries)
}
