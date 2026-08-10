package usecase

import (
	"context"
	"encoding/base64"
	"fmt"

	"math"
	"math/rand"
	"strconv"
	"strings"
	"time"

	"day-32/domain"

	"github.com/google/uuid"
)

type ArticleUseCase struct {
	repo   domain.ArticleRepository
	logger domain.Logger
}

func NewArticleUseCase(repo domain.ArticleRepository, logger domain.Logger) *ArticleUseCase {
	return &ArticleUseCase{
		repo:   repo,
		logger: logger,
	}
}

func (u *ArticleUseCase) GetArticlesWithOffset(ctx context.Context, filters domain.FilterParams, pag domain.PaginationQuery) (*domain.PaginatedResponse, error) {
	if pag.Page < 1 {
		pag.Page = 1
	}
	if pag.Limit < 1 {
		pag.Limit = 10
	}
	if pag.Limit > 100 {
		pag.Limit = 100
	}

	articles, totalRecords, err := u.repo.FindWithOffset(ctx, filters, pag)
	if err != nil {
		u.logger.Error(ctx, "Failed to fetch offset paginated articles", "error", err)
		return nil, err
	}

	totalPages := int(math.Ceil(float64(totalRecords) / float64(pag.Limit)))
	if totalPages == 0 {
		totalPages = 1
	}

	hasNext := pag.Page < totalPages
	hasPrev := pag.Page > 1

	meta := domain.PaginationMeta{
		Page:         pag.Page,
		Limit:        pag.Limit,
		TotalRecords: totalRecords,
		TotalPages:   totalPages,
		HasNext:      hasNext,
		HasPrev:      hasPrev,
	}

	return &domain.PaginatedResponse{
		Data: articles,
		Meta: meta,
	}, nil
}

func (u *ArticleUseCase) encodeCursor(publishedAt time.Time, id string) string {
	raw := fmt.Sprintf("%d:%s", publishedAt.UnixNano(), id)
	return base64.StdEncoding.EncodeToString([]byte(raw))
}

func (u *ArticleUseCase) decodeCursor(cursorStr string) (string, *time.Time, error) {
	if cursorStr == "" {
		return "", nil, nil
	}

	decodedBytes, err := base64.StdEncoding.DecodeString(cursorStr)
	if err != nil {
		return "", nil, fmt.Errorf("invalid cursor base64 string: %w", err)
	}

	parts := strings.Split(string(decodedBytes), ":")
	if len(parts) != 2 {
		return "", nil, fmt.Errorf("invalid cursor format")
	}

	nanoUnix, err := strconv.ParseInt(parts[0], 10, 64)
	if err != nil {
		return "", nil, fmt.Errorf("invalid cursor timestamp")
	}

	pubTime := time.Unix(0, nanoUnix)
	id := parts[1]
	return id, &pubTime, nil
}

func (u *ArticleUseCase) GetArticlesWithCursor(ctx context.Context, filters domain.FilterParams, cursorQuery domain.CursorQuery) (*domain.CursorPaginatedResponse, error) {
	limit := cursorQuery.Limit
	if limit < 1 {
		limit = 10
	}
	if limit > 100 {
		limit = 100
	}

	cursorID, cursorTime, err := u.decodeCursor(cursorQuery.Cursor)
	if err != nil {
		u.logger.Warn(ctx, "Failed to decode pagination cursor", "cursor", cursorQuery.Cursor, "error", err)
		return nil, domain.ErrInvalidInput
	}

	articles, hasNext, err := u.repo.FindWithCursor(ctx, filters, limit, cursorID, cursorTime)
	if err != nil {
		u.logger.Error(ctx, "Failed to fetch cursor paginated articles", "error", err)
		return nil, err
	}

	var nextCursor string
	if hasNext && len(articles) > 0 {
		lastItem := articles[len(articles)-1]
		nextCursor = u.encodeCursor(lastItem.PublishedAt, lastItem.ID)
	}

	meta := domain.CursorMeta{
		NextCursor: nextCursor,
		HasNext:    hasNext,
		Limit:      limit,
	}

	return &domain.CursorPaginatedResponse{
		Data: articles,
		Meta: meta,
	}, nil
}

func (u *ArticleUseCase) SeedArticles(ctx context.Context, count int) error {
	categories := []string{"TECH", "FINANCE", "HEALTH", "EDUCATION", "TRAVEL"}
	authors := []string{"Alice Johnson", "Bob Smith", "Charlie Brown", "Diana Prince"}

	articles := make([]*domain.Article, 0, count)
	now := time.Now()

	for i := 1; i <= count; i++ {
		pubTime := now.Add(-time.Duration(count-i) * time.Hour)
		article := &domain.Article{
			ID:          fmt.Sprintf("art_%s", uuid.New().String()[:8]),
			Title:       fmt.Sprintf("Go Backend Engineering Guide #%d", i),
			Category:    categories[rand.Intn(len(categories))],
			Author:      authors[rand.Intn(len(authors))],
			Views:       (i * 150) % 3000,
			PublishedAt: pubTime,
			CreatedAt:   pubTime,
		}
		articles = append(articles, article)
	}

	return u.repo.BulkSave(ctx, articles)
}


