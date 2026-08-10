package handler

import (
	"net/http"

	"day-30/domain"
	"day-30/usecase"

	"github.com/gin-gonic/gin"
)

type StudentHandler struct {
	studentUseCase *usecase.StudentUseCase
	logger         domain.Logger
}

func NewStudentHandler(studentUseCase *usecase.StudentUseCase, logger domain.Logger) *StudentHandler {
	return &StudentHandler{
		studentUseCase: studentUseCase,
		logger:         logger,
	}
}

func (h *StudentHandler) CreateStudent(c *gin.Context) {
	var input domain.CreateStudentInput
	if err := c.ShouldBindJSON(&input); err != nil {
		h.logger.Warn(c.Request.Context(), "Invalid Create Student JSON body", "error", err.Error())
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request payload", "details": err.Error()})
		return
	}

	student, err := h.studentUseCase.CreateStudent(c.Request.Context(), input)
	if err != nil {
		switch err {
		case domain.ErrEmailExists:
			c.JSON(http.StatusConflict, gin.H{"error": err.Error()})
		case domain.ErrInvalidInput:
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		default:
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Internal server error"})
		}
		return
	}

	c.JSON(http.StatusCreated, gin.H{
		"message": "Student created successfully",
		"data":    student,
	})
}

func (h *StudentHandler) GetStudentByID(c *gin.Context) {
	id := c.Param("id")
	student, err := h.studentUseCase.GetStudentByID(c.Request.Context(), id)
	if err != nil {
		if err == domain.ErrStudentNotFound {
			c.JSON(http.StatusNotFound, gin.H{"error": "Student record not found"})
			return
		}
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Internal server error"})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"data": student,
	})
}

func (h *StudentHandler) ListStudents(c *gin.Context) {
	department := c.Query("department")
	status := c.Query("status")

	students, err := h.studentUseCase.ListStudents(c.Request.Context(), department, status)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Internal server error"})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"count": len(students),
		"data":  students,
	})
}

func (h *StudentHandler) UpdateStudent(c *gin.Context) {
	id := c.Param("id")
	var input domain.UpdateStudentInput
	if err := c.ShouldBindJSON(&input); err != nil {
		h.logger.Warn(c.Request.Context(), "Invalid Update Student JSON body", "error", err.Error())
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request payload", "details": err.Error()})
		return
	}

	student, err := h.studentUseCase.UpdateStudent(c.Request.Context(), id, input)
	if err != nil {
		if err == domain.ErrStudentNotFound {
			c.JSON(http.StatusNotFound, gin.H{"error": "Student record not found"})
			return
		}
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Internal server error"})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "Student updated successfully",
		"data":    student,
	})
}

func (h *StudentHandler) DeleteStudent(c *gin.Context) {
	id := c.Param("id")
	if err := h.studentUseCase.DeleteStudent(c.Request.Context(), id); err != nil {
		if err == domain.ErrStudentNotFound {
			c.JSON(http.StatusNotFound, gin.H{"error": "Student record not found"})
			return
		}
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Internal server error"})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "Student deleted successfully",
	})
}
