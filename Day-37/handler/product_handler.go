package handler

import (
	"net/http"

	"day-37/domain"
	"day-37/usecase"

	"github.com/gin-gonic/gin"
)

type ProductHandler struct {
	productUC *usecase.ProductUseCase
	logger    domain.Logger
}

func NewProductHandler(productUC *usecase.ProductUseCase, logger domain.Logger) *ProductHandler {
	return &ProductHandler{
		productUC: productUC,
		logger:    logger,
	}
}

func (h *ProductHandler) CreateProduct(c *gin.Context) {
	var input domain.CreateProductInput
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request payload"})
		return
	}

	p, err := h.productUC.CreateProduct(c.Request.Context(), input)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, gin.H{"message": "Product created", "data": p})
}

func (h *ProductHandler) GetProductByID(c *gin.Context) {
	id := c.Param("id")
	p, err := h.productUC.GetProductByID(c.Request.Context(), id)
	if err != nil {
		if err == domain.ErrProductNotFound {
			c.JSON(http.StatusNotFound, gin.H{"error": "Product not found"})
		} else {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		}
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": p})
}

func (h *ProductHandler) UpdateProduct(c *gin.Context) {
	id := c.Param("id")
	var input domain.UpdateProductInput
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request payload"})
		return
	}

	p, err := h.productUC.UpdateProduct(c.Request.Context(), id, input)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Product updated and cache invalidated", "data": p})
}
