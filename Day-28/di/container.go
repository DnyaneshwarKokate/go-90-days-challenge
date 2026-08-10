package di

import (
	"day-28/domain"
	"day-28/handler"
	"day-28/logger"
	"day-28/notification"
	"day-28/repository"
	"day-28/usecase"

	"github.com/gin-gonic/gin"
)

// Container holds application singletons and dependencies.
type Container struct {
	Logger       domain.Logger
	OrderRepo    domain.OrderRepository
	Notification domain.NotificationService
	OrderUseCase *usecase.OrderUseCase
	OrderHandler *handler.OrderHandler
}

// Config allows customizing DI container behaviors (e.g., Environment mode).
type Config struct {
	Env string // "development", "production", "test"
}

// NewContainer performs central Dependency Injection wireup.
func NewContainer(cfg Config) *Container {
	var appLogger domain.Logger
	var notifier domain.NotificationService

	if cfg.Env == "test" {
		appLogger = logger.NewMockLogger()
		notifier = notification.NewMockNotificationService()
	} else {
		appLogger = logger.NewConsoleLogger("APP")
		notifier = notification.NewEmailNotificationService(appLogger)
	}

	// 1. Initialize Persistence Repository
	orderRepo := repository.NewMemoryOrderRepository()

	// 2. Initialize UseCase layer (Constructor Injection)
	orderUseCase := usecase.NewOrderUseCase(orderRepo, notifier, appLogger)

	// 3. Initialize HTTP Handlers
	orderHandler := handler.NewOrderHandler(orderUseCase, appLogger)

	return &Container{
		Logger:       appLogger,
		OrderRepo:    orderRepo,
		Notification: notifier,
		OrderUseCase: orderUseCase,
		OrderHandler: orderHandler,
	}
}

// SetupRoutes attaches container handlers to Gin router endpoints.
func (c *Container) SetupRoutes(router *gin.Engine) {
	api := router.Group("/api/v1")
	{
		api.POST("/orders", c.OrderHandler.CreateOrder)
		api.GET("/orders", c.OrderHandler.ListOrders)
		api.GET("/orders/:id", c.OrderHandler.GetOrder)
	}
}
