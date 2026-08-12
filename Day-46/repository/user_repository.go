package repository

import (
	"context"
	"database/sql"
	"encoding/json"
	"fmt"
	"log"
	"sync/atomic"
	"time"

	"github.com/redis/go-redis/v9"
)

// User represents a user entity in the database.
type User struct {
	ID        int       `json:"id"`
	Name      string    `json:"name"`
	Email     string    `json:"email"`
	Role      string    `json:"role"`
	CreatedAt time.Time `json:"created_at"`
}

// Stats metrics for cache hits and misses.
type CacheStats struct {
	Hits   int64 `json:"hits"`
	Misses int64 `json:"misses"`
}

type UserRepository interface {
	GetByID(ctx context.Context, id int) (*User, bool, error)
	GetAll(ctx context.Context) ([]User, error)
	Create(ctx context.Context, name, email, role string) (*User, error)
	GetStats() CacheStats
}

type PostgresRedisRepo struct {
	db     *sql.DB
	rdb    *redis.Client
	hits   int64
	misses int64
}

func NewUserRepository(db *sql.DB, rdb *redis.Client) *PostgresRedisRepo {
	return &PostgresRedisRepo{
		db:  db,
		rdb: rdb,
	}
}

func (r *PostgresRedisRepo) GetByID(ctx context.Context, id int) (*User, bool, error) {
	cacheKey := fmt.Sprintf("user:%d", id)

	// Try Redis Cache first if client is available
	if r.rdb != nil {
		val, err := r.rdb.Get(ctx, cacheKey).Result()
		if err == nil && val != "" {
			var u User
			if err := json.Unmarshal([]byte(val), &u); err == nil {
				atomic.AddInt64(&r.hits, 1)
				return &u, true, nil
			}
		}
	}

	atomic.AddInt64(&r.misses, 1)

	// Fetch from PostgreSQL
	query := `SELECT id, name, email, role, created_at FROM users WHERE id = $1`
	var u User
	err := r.db.QueryRowContext(ctx, query, id).Scan(&u.ID, &u.Name, &u.Email, &u.Role, &u.CreatedAt)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, false, nil
		}
		return nil, false, err
	}

	// Cache result in Redis for 60 seconds
	if r.rdb != nil {
		data, err := json.Marshal(u)
		if err == nil {
			if setErr := r.rdb.Set(ctx, cacheKey, data, 60*time.Second).Err(); setErr != nil {
				log.Printf("Failed to set cache for %s: %v", cacheKey, setErr)
			}
		}
	}

	return &u, false, nil
}

func (r *PostgresRedisRepo) GetAll(ctx context.Context) ([]User, error) {
	query := `SELECT id, name, email, role, created_at FROM users ORDER BY id ASC`
	rows, err := r.db.QueryContext(ctx, query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	users := make([]User, 0)
	for rows.Next() {
		var u User
		if err := rows.Scan(&u.ID, &u.Name, &u.Email, &u.Role, &u.CreatedAt); err != nil {
			return nil, err
		}
		users = append(users, u)
	}
	return users, rows.Err()
}

func (r *PostgresRedisRepo) Create(ctx context.Context, name, email, role string) (*User, error) {
	query := `INSERT INTO users (name, email, role, created_at) VALUES ($1, $2, $3, NOW()) RETURNING id, name, email, role, created_at`
	var u User
	err := r.db.QueryRowContext(ctx, query, name, email, role).Scan(&u.ID, &u.Name, &u.Email, &u.Role, &u.CreatedAt)
	if err != nil {
		return nil, err
	}

	return &u, nil
}

func (r *PostgresRedisRepo) GetStats() CacheStats {
	return CacheStats{
		Hits:   atomic.LoadInt64(&r.hits),
		Misses: atomic.LoadInt64(&r.misses),
	}
}
