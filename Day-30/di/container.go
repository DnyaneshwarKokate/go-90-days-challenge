package di

import (
	"day-30/config"
	"day-30/domain"
	"day-30/handler"
	"day-30/logger"
	"day-30/middleware"
	"day-30/repository"
	"day-30/usecase"

	"github.com/gin-gonic/gin"
)

type Container struct {
	Config         config.Config
	Logger         domain.Logger
	UserRepo       domain.UserRepository
	StudentRepo    domain.StudentRepository
	AuthUseCase    *usecase.AuthUseCase
	StudentUseCase *usecase.StudentUseCase
	AuthHandler    *handler.AuthHandler
	StudentHandler *handler.StudentHandler
}

func NewContainer(cfg config.Config) (*Container, error) {
	appLogger, err := logger.NewZapLogger(cfg.Environment, cfg.LogFilePath)
	if err != nil {
		return nil, err
	}

	userRepo := repository.NewMemoryUserRepository(appLogger)
	studentRepo := repository.NewMemoryStudentRepository(appLogger)

	authUseCase := usecase.NewAuthUseCase(userRepo, cfg, appLogger)
	studentUseCase := usecase.NewStudentUseCase(studentRepo, appLogger)

	authHandler := handler.NewAuthHandler(authUseCase, appLogger)
	studentHandler := handler.NewStudentHandler(studentUseCase, appLogger)

	return &Container{
		Config:         cfg,
		Logger:         appLogger,
		UserRepo:       userRepo,
		StudentRepo:    studentRepo,
		AuthUseCase:    authUseCase,
		StudentUseCase: studentUseCase,
		AuthHandler:    authHandler,
		StudentHandler: studentHandler,
	}, nil
}

func (c *Container) SetupRouter() *gin.Engine {
	if c.Config.Environment == "production" {
		gin.SetMode(gin.ReleaseMode)
	} else {
		gin.SetMode(gin.DebugMode)
	}

	router := gin.New()
	router.Use(gin.Recovery())
	router.Use(middleware.RequestIDMiddleware())
	router.Use(middleware.StructuredLoggerMiddleware(c.Logger))

	v1 := router.Group("/api/v1")
	{
		authGroup := v1.Group("/auth")
		{
			authGroup.POST("/register", c.AuthHandler.Register)
			authGroup.POST("/login", c.AuthHandler.Login)
		}

		studentGroup := v1.Group("/students")
		studentGroup.Use(middleware.JWTAuthMiddleware(c.Config.JWTSecret, c.Logger))
		{
			// Read student endpoints accessible by ADMIN, TEACHER, STUDENT
			studentGroup.GET("", middleware.RequireRoleMiddleware(c.Logger, domain.RoleAdmin, domain.RoleTeacher, domain.RoleStudent), c.StudentHandler.ListStudents)
			studentGroup.GET("/:id", middleware.RequireRoleMiddleware(c.Logger, domain.RoleAdmin, domain.RoleTeacher, domain.RoleStudent), c.StudentHandler.GetStudentByID)

			// Write student endpoints accessible by ADMIN, TEACHER
			studentGroup.POST("", middleware.RequireRoleMiddleware(c.Logger, domain.RoleAdmin, domain.RoleTeacher), c.StudentHandler.CreateStudent)
			studentGroup.PUT("/:id", middleware.RequireRoleMiddleware(c.Logger, domain.RoleAdmin, domain.RoleTeacher), c.StudentHandler.UpdateStudent)

			// Delete student endpoint accessible ONLY by ADMIN
			studentGroup.DELETE("/:id", middleware.RequireRoleMiddleware(c.Logger, domain.RoleAdmin), c.StudentHandler.DeleteStudent)
		}
	}

	return router
}
