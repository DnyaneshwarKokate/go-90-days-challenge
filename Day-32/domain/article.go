package domain

import (
	"context"
	"errors"
	"time"
)

type ContextKey string

const (
	RequestIDKey ContextKey = "X-Request-ID"
)

var (
	ErrInvalidInput = errors.New("invalid request input parameters")
)

type Article struct {
	ID          string    `json:"id"`
	Title       string    `json:"title"`
	Category    string    `json:"category"`
	Author      string    `json:"author"`
	Views       int       `json:"views"`
	PublishedAt time.Time `json:"published_at"`
	CreatedAt   time.Time `json:"created_at"`
}

type FilterParams struct {
	Category string `form:"category"`
	MinViews int    `form:"min_views"`
	Search   string `form:"search"`
}

type PaginationQuery struct {
	Page   int    `form:"page,default=1"`
	Limit  int    `form:"limit,default=10"`
	SortBy string `form:"sort_by,default=created_at"`
	Order  string `form:"order,default=desc"`
}

type CursorQuery struct {
	Cursor string `form:"cursor"`
	Limit  int    `form:"limit,default=10"`
}

type PaginationMeta struct {
	Page         int  `json:"page"`
	Limit        int  `json:"limit"`
	TotalRecords int  `json:"total_records"`
	TotalPages   int  `json:"total_pages"`
	HasNext      bool `json:"has_next"`
	HasPrev      bool `json:"has_prev"`
}

type CursorMeta struct {
	NextCursor string `json:"next_cursor"`
	HasNext    bool   `json:"has_next"`
	Limit      int    `json:"limit"`
}

type PaginatedResponse struct {
	Data []*Article     `json:"data"`
	Meta PaginationMeta `json:"meta"`
}

type CursorPaginatedResponse struct {
	Data []*Article `json:"data"`
	Meta CursorMeta `json:"meta"`
}

type Logger interface {
	Info(ctx context.Context, msg string, keysAndValues ...interface{})
	Warn(ctx context.Context, msg string, keysAndValues ...interface{})
	Error(ctx context.Context, msg string, keysAndValues ...interface{})
	Debug(ctx context.Context, msg string, keysAndValues ...interface{})
}

type ArticleRepository interface {
	Save(ctx context.Context, article *Article) error
	BulkSave(ctx context.Context, articles []*Article) error
	FindWithOffset(ctx context.Context, filters FilterParams, pag PaginationQuery) ([]*Article, int, error)
	FindWithCursor(ctx context.Context, filters FilterParams, limit int, cursorID string, cursorTime *time.Time) ([]*Article, bool, error)
}
