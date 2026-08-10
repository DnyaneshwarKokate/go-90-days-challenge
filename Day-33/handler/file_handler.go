package handler

import (
	"net/http"
	"os"

	"day-33/domain"
	"day-33/service"

	"github.com/gin-gonic/gin"
)

type FileHandler struct {
	fileService *service.FileService
	logger      domain.Logger
}

func NewFileHandler(fileService *service.FileService, logger domain.Logger) *FileHandler {
	return &FileHandler{
		fileService: fileService,
		logger:      logger,
	}
}

func (h *FileHandler) UploadSingle(c *gin.Context) {
	fileHeader, err := c.FormFile("file")
	if err != nil {
		h.logger.Warn(c.Request.Context(), "Upload request missing file parameter")
		c.JSON(http.StatusBadRequest, gin.H{"error": domain.ErrNoFileProvided.Error()})
		return
	}

	meta, err := h.fileService.SaveUploadedFile(c.Request.Context(), fileHeader)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, gin.H{
		"message": "File uploaded successfully",
		"data":    meta,
	})
}

func (h *FileHandler) UploadMultiple(c *gin.Context) {
	form, err := c.MultipartForm()
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid multipart form payload"})
		return
	}

	files := form.File["files"]
	if len(files) == 0 {
		c.JSON(http.StatusBadRequest, gin.H{"error": domain.ErrNoFileProvided.Error()})
		return
	}

	metas, err := h.fileService.SaveMultipleFiles(c.Request.Context(), files)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, gin.H{
		"message": "Multiple files uploaded successfully",
		"count":   len(metas),
		"data":    metas,
	})
}

func (h *FileHandler) DownloadFile(c *gin.Context) {
	filename := c.Param("filename")
	path := h.fileService.GetFilePath(filename)

	if _, err := os.Stat(path); os.IsNotExist(err) {
		c.JSON(http.StatusNotFound, gin.H{"error": "File not found"})
		return
	}

	c.File(path)
}
