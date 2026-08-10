package handler

import (
	"net/http"

	"day-28/domain"
	"day-28/usecase"

	"github.com/gin-gonic/gin"
)

// OrderHandler handles HTTP requests for Order operations.
type OrderHandler struct {
	orderUC *usecase.OrderUseCase
	logger  domain.Logger
}

// NewOrderHandler creates an OrderHandler with injected dependencies.
func NewOrderHandler(orderUC *usecase.OrderUseCase, logger domain.Logger) *OrderHandler {
	return &OrderHandler{
		orderUC: orderUC,
		logger:  logger,
	}
}

// CreateOrder handles POST /api/v1/orders
func (h *OrderHandler) CreateOrder(c *gin.Context) {
	var input domain.CreateOrderInput
	if err := c.ShouldBindJSON(&input); err != nil {
		h.logger.Error("Invalid JSON request body: %v", err)
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request payload", "details": err.Error()})
		return
	}

	order, err := h.orderUC.CreateOrder(input)
	if err != nil {
		h.logger.Error("Order creation failed: %v", err)
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, gin.H{
		"message": "Order created successfully",
		"data":    order,
	})
}

// GetOrder handles GET /api/v1/orders/:id
func (h *OrderHandler) GetOrder(c *gin.Context) {
	id := c.Param("id")
	order, err := h.orderUC.GetOrder(id)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": order})
}

// ListOrders handles GET /api/v1/orders
func (h *OrderHandler) ListOrders(c *gin.Context) {
	orders, err := h.orderUC.ListOrders()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": orders, "count": len(orders)})
}
