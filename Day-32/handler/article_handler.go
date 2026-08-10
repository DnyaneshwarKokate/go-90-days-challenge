package handler

import (
	"net/http"
	"strconv"

	"day-32/domain"
	"day-32/usecase"

	"github.com/gin-gonic/gin"
)

type ArticleHandler struct {
	articleUseCase *usecase.ArticleUseCase
	logger          domain.Logger
}

func NewArticleHandler(articleUseCase *usecase.ArticleUseCase, logger domain.Logger) *ArticleHandler {
	return &ArticleHandler{
		articleUseCase: articleUseCase,
		logger:         logger,
	}
}

func (h *ArticleHandler) GetArticlesOffset(c *gin.Context) {
	var filters domain.FilterParams
	if err := c.ShouldBindQuery(&filters); err != nil {
		h.logger.Warn(c.Request.Context(), "Invalid filter parameters", "error", err.Error())
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid filter query parameters"})
		return
	}

	var pag domain.PaginationQuery
	if err := c.ShouldBindQuery(&pag); err != nil {
		h.logger.Warn(c.Request.Context(), "Invalid pagination parameters", "error", err.Error())
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid pagination query parameters"})
		return
	}

	res, err := h.articleUseCase.GetArticlesWithOffset(c.Request.Context(), filters, pag)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Internal server error"})
		return
	}

	c.JSON(http.StatusOK, res)
}

func (h *ArticleHandler) GetArticlesCursor(c *gin.Context) {
	var filters domain.FilterParams
	if err := c.ShouldBindQuery(&filters); err != nil {
		h.logger.Warn(c.Request.Context(), "Invalid filter parameters", "error", err.Error())
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid filter query parameters"})
		return
	}

	var cursorQuery domain.CursorQuery
	if err := c.ShouldBindQuery(&cursorQuery); err != nil {
		h.logger.Warn(c.Request.Context(), "Invalid cursor query parameters", "error", err.Error())
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid cursor query parameters"})
		return
	}

	res, err := h.articleUseCase.GetArticlesWithCursor(c.Request.Context(), filters, cursorQuery)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid pagination request or cursor string"})
		return
	}

	c.JSON(http.StatusOK, res)
}

func (h *ArticleHandler) SeedArticles(c *gin.Context) {
	countStr := c.DefaultQuery("count", "25")
	count, err := strconv.Atoi(countStr)
	if err != nil || count <= 0 {
		count = 25
	}

	if err := h.articleUseCase.SeedArticles(c.Request.Context(), count); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to seed articles"})
		return
	}

	c.JSON(http.StatusCreated, gin.H{
		"message": "Articles seeded successfully",
		"count":   count,
	})
}
