package repository

import (
	"context"
	"sort"
	"strings"
	"sync"
	"time"

	"day-32/domain"
)

type memoryArticleRepository struct {
	mu       sync.RWMutex
	articles map[string]*domain.Article
	logger   domain.Logger
}

func NewMemoryArticleRepository(logger domain.Logger) domain.ArticleRepository {
	return &memoryArticleRepository{
		articles: make(map[string]*domain.Article),
		logger:   logger,
	}
}

func (r *memoryArticleRepository) Save(ctx context.Context, article *domain.Article) error {
	r.mu.Lock()
	defer r.mu.Unlock()

	r.articles[article.ID] = article
	return nil
}

func (r *memoryArticleRepository) BulkSave(ctx context.Context, articles []*domain.Article) error {
	r.mu.Lock()
	defer r.mu.Unlock()

	for _, a := range articles {
		r.articles[a.ID] = a
	}
	r.logger.Info(ctx, "DB BulkSave articles completed", "count", len(articles))
	return nil
}

func (r *memoryArticleRepository) filterArticles(filters domain.FilterParams) []*domain.Article {
	filtered := make([]*domain.Article, 0)

	for _, a := range r.articles {
		if filters.Category != "" && !strings.EqualFold(a.Category, filters.Category) {
			continue
		}
		if filters.MinViews > 0 && a.Views < filters.MinViews {
			continue
		}
		if filters.Search != "" {
			term := strings.ToLower(filters.Search)
			if !strings.Contains(strings.ToLower(a.Title), term) && !strings.Contains(strings.ToLower(a.Author), term) {
				continue
			}
		}
		filtered = append(filtered, a)
	}

	return filtered
}

func (r *memoryArticleRepository) sortArticles(articles []*domain.Article, sortBy string, order string) {
	desc := strings.ToLower(order) == "desc"

	sort.Slice(articles, func(i, j int) bool {
		switch strings.ToLower(sortBy) {
		case "views":
			if desc {
				return articles[i].Views > articles[j].Views
			}
			return articles[i].Views < articles[j].Views
		case "published_at":
			if desc {
				return articles[i].PublishedAt.After(articles[j].PublishedAt)
			}
			return articles[i].PublishedAt.Before(articles[j].PublishedAt)
		case "title":
			if desc {
				return articles[i].Title > articles[j].Title
			}
			return articles[i].Title < articles[j].Title
		default: // created_at
			if desc {
				return articles[i].CreatedAt.After(articles[j].CreatedAt)
			}
			return articles[i].CreatedAt.Before(articles[j].CreatedAt)
		}
	})
}

func (r *memoryArticleRepository) FindWithOffset(ctx context.Context, filters domain.FilterParams, pag domain.PaginationQuery) ([]*domain.Article, int, error) {
	r.mu.RLock()
	defer r.mu.RUnlock()

	r.logger.Debug(ctx, "Executing Offset Pagination query", "page", pag.Page, "limit", pag.Limit, "sort_by", pag.SortBy, "order", pag.Order)

	filtered := r.filterArticles(filters)
	totalCount := len(filtered)

	r.sortArticles(filtered, pag.SortBy, pag.Order)

	offset := (pag.Page - 1) * pag.Limit
	if offset >= totalCount {
		return []*domain.Article{}, totalCount, nil
	}

	end := offset + pag.Limit
	if end > totalCount {
		end = totalCount
	}

	result := filtered[offset:end]
	r.logger.Info(ctx, "Offset Pagination query completed", "result_count", len(result), "total_count", totalCount)
	return result, totalCount, nil
}

func (r *memoryArticleRepository) FindWithCursor(ctx context.Context, filters domain.FilterParams, limit int, cursorID string, cursorTime *time.Time) ([]*domain.Article, bool, error) {
	r.mu.RLock()
	defer r.mu.RUnlock()

	r.logger.Debug(ctx, "Executing Cursor Keyset Pagination query", "limit", limit, "cursor_id", cursorID)

	filtered := r.filterArticles(filters)

	// Keyset sorting by PublishedAt DESC, then ID DESC for determinism
	sort.Slice(filtered, func(i, j int) bool {
		if filtered[i].PublishedAt.Equal(filtered[j].PublishedAt) {
			return filtered[i].ID > filtered[j].ID
		}
		return filtered[i].PublishedAt.After(filtered[j].PublishedAt)
	})

	startIndex := 0
	if cursorTime != nil && cursorID != "" {
		for i, a := range filtered {
			if a.PublishedAt.Before(*cursorTime) || (a.PublishedAt.Equal(*cursorTime) && a.ID < cursorID) {
				startIndex = i
				break
			}
			if i == len(filtered)-1 {
				startIndex = len(filtered)
			}
		}
	}

	if startIndex >= len(filtered) {
		return []*domain.Article{}, false, nil
	}

	// Fetch limit + 1 to accurately detect if hasNext is true
	end := startIndex + limit + 1
	hasNext := false

	if end > len(filtered) {
		end = len(filtered)
	} else {
		hasNext = true
	}

	slice := filtered[startIndex:end]
	if hasNext && len(slice) > limit {
		slice = slice[:limit]
	}

	r.logger.Info(ctx, "Cursor Pagination completed", "fetched_count", len(slice), "has_next", hasNext)
	return slice, hasNext, nil
}
