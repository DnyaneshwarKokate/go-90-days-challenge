package handler

import (
	"net/http"
	"strconv"

	"day-31/domain"
	"day-31/usecase"

	"github.com/gin-gonic/gin"
)

type ProductHandler struct {
	productUseCase *usecase.ProductUseCase
	logger         domain.Logger
}

func NewProductHandler(productUseCase *usecase.ProductUseCase, logger domain.Logger) *ProductHandler {
	return &ProductHandler{
		productUseCase: productUseCase,
		logger:         logger,
	}
}

func (h *ProductHandler) CreateProduct(c *gin.Context) {
	var input domain.CreateProductInput
	if err := c.ShouldBindJSON(&input); err != nil {
		h.logger.Warn(c.Request.Context(), "Invalid Create Product JSON body", "error", err.Error())
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request payload", "details": err.Error()})
		return
	}

	product, err := h.productUseCase.CreateProduct(c.Request.Context(), input)
	if err != nil {
		if err == domain.ErrSKUExists {
			c.JSON(http.StatusConflict, gin.H{"error": err.Error()})
			return
		}
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.Header("ETag", strconv.Itoa(product.Version))
	c.JSON(http.StatusCreated, gin.H{
		"message": "Product created successfully",
		"data":    product,
	})
}

func (h *ProductHandler) BulkCreateProducts(c *gin.Context) {
	var input domain.BulkCreateInput
	if err := c.ShouldBindJSON(&input); err != nil {
		h.logger.Warn(c.Request.Context(), "Invalid Bulk Create JSON payload", "error", err.Error())
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid bulk creation payload", "details": err.Error()})
		return
	}

	products, err := h.productUseCase.BulkCreateProducts(c.Request.Context(), input)
	if err != nil {
		if err == domain.ErrSKUExists {
			c.JSON(http.StatusConflict, gin.H{"error": err.Error()})
			return
		}
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, gin.H{
		"message": "Bulk product creation completed successfully",
		"count":   len(products),
		"data":    products,
	})
}

func (h *ProductHandler) GetProductByID(c *gin.Context) {
	id := c.Param("id")
	includeDeleted, _ := strconv.ParseBool(c.DefaultQuery("include_deleted", "false"))

	product, err := h.productUseCase.GetProductByID(c.Request.Context(), id, includeDeleted)
	if err != nil {
		if err == domain.ErrProductNotFound {
			c.JSON(http.StatusNotFound, gin.H{"error": "Product not found"})
			return
		}
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Internal server error"})
		return
	}

	c.Header("ETag", strconv.Itoa(product.Version))
	c.JSON(http.StatusOK, gin.H{"data": product})
}

func (h *ProductHandler) ListProducts(c *gin.Context) {
	includeDeleted, _ := strconv.ParseBool(c.DefaultQuery("include_deleted", "false"))

	products, err := h.productUseCase.ListProducts(c.Request.Context(), includeDeleted)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Internal server error"})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"count": len(products),
		"data":  products,
	})
}

func (h *ProductHandler) PatchUpdateProduct(c *gin.Context) {
	id := c.Param("id")
	var input domain.PatchProductInput

	if err := c.ShouldBindJSON(&input); err != nil {
		h.logger.Warn(c.Request.Context(), "Invalid PATCH JSON payload", "error", err.Error())
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request payload", "details": err.Error()})
		return
	}

	// Check optional If-Match header for Optimistic Concurrency Control
	if ifMatch := c.GetHeader("If-Match"); ifMatch != "" {
		if ver, err := strconv.Atoi(ifMatch); err == nil {
			input.ExpectedVersion = &ver
		}
	}

	product, err := h.productUseCase.PatchUpdateProduct(c.Request.Context(), id, input)
	if err != nil {
		switch err {
		case domain.ErrConcurrencyConflict:
			c.JSON(http.StatusConflict, gin.H{"error": err.Error()})
		case domain.ErrProductNotFound:
			c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
		default:
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Internal server error"})
		}
		return
	}

	c.Header("ETag", strconv.Itoa(product.Version))
	c.JSON(http.StatusOK, gin.H{
		"message": "Product patched successfully",
		"data":    product,
	})
}

func (h *ProductHandler) SoftDeleteProduct(c *gin.Context) {
	id := c.Param("id")

	if err := h.productUseCase.SoftDeleteProduct(c.Request.Context(), id); err != nil {
		switch err {
		case domain.ErrAlreadyDeleted:
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		case domain.ErrProductNotFound:
			c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
		default:
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Internal server error"})
		}
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Product soft-deleted successfully"})
}

func (h *ProductHandler) RestoreProduct(c *gin.Context) {
	id := c.Param("id")

	product, err := h.productUseCase.RestoreProduct(c.Request.Context(), id)
	if err != nil {
		switch err {
		case domain.ErrNotDeleted:
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		case domain.ErrProductNotFound:
			c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
		default:
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Internal server error"})
		}
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "Product restored successfully",
		"data":    product,
	})
}

func (h *ProductHandler) BulkSoftDeleteProducts(c *gin.Context) {
	var input domain.BulkDeleteInput
	if err := c.ShouldBindJSON(&input); err != nil {
		h.logger.Warn(c.Request.Context(), "Invalid Bulk Delete JSON payload", "error", err.Error())
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request payload", "details": err.Error()})
		return
	}

	count, err := h.productUseCase.BulkSoftDeleteProducts(c.Request.Context(), input.IDs)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message":       "Bulk soft delete operation completed",
		"deleted_count": count,
	})
}
