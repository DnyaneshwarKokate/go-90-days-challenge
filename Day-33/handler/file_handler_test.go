package handler_test

import (
	"bytes"
	"mime/multipart"
	"net/http"
	"net/http/httptest"
	"os"

	"testing"

	"day-33/handler"
	"day-33/logger"
	"day-33/middleware"
	"day-33/service"

	"github.com/gin-gonic/gin"
)

func setupTestRouter() *gin.Engine {
	gin.SetMode(gin.TestMode)
	zapLog, _ := logger.NewZapLogger("test", "")

	os.MkdirAll("./test_uploads", 0755)
	fileSvc, _ := service.NewFileService("./test_uploads", 5*1024*1024, zapLog)
	fileHandler := handler.NewFileHandler(fileSvc, zapLog)

	router := gin.New()
	router.Use(middleware.RequestIDMiddleware())
	router.Use(middleware.StructuredLoggerMiddleware(zapLog))

	v1 := router.Group("/api/v1/files")
	{
		v1.POST("/upload", fileHandler.UploadSingle)
		v1.POST("/upload-multiple", fileHandler.UploadMultiple)
		v1.GET("/download/:filename", fileHandler.DownloadFile)
	}

	return router
}

func TestFileUpload_SingleSuccess(t *testing.T) {
	router := setupTestRouter()

	body := &bytes.Buffer{}
	writer := multipart.NewWriter(body)
	part, _ := writer.CreateFormFile("file", "avatar.png")
	part.Write([]byte("fake-png-data"))
	writer.Close()

	req, _ := http.NewRequest("POST", "/api/v1/files/upload", body)
	req.Header.Set("Content-Type", writer.FormDataContentType())

	w := httptest.NewRecorder()
	router.ServeHTTP(w, req)

	if w.Code != http.StatusCreated {
		t.Fatalf("expected 201 Created, got %d. Body: %s", w.Code, w.Body.String())
	}
}

func TestFileUpload_InvalidType(t *testing.T) {
	router := setupTestRouter()

	body := &bytes.Buffer{}
	writer := multipart.NewWriter(body)
	part, _ := writer.CreateFormFile("file", "script.sh")
	part.Write([]byte("echo hello"))
	writer.Close()

	req, _ := http.NewRequest("POST", "/api/v1/files/upload", body)
	req.Header.Set("Content-Type", writer.FormDataContentType())

	w := httptest.NewRecorder()
	router.ServeHTTP(w, req)

	if w.Code != http.StatusBadRequest {
		t.Errorf("expected 400 Bad Request for .sh file, got %d", w.Code)
	}
}
