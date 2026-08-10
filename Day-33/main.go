package main

import (
	"bytes"
	"context"
	"fmt"
	"mime/multipart"
	"net/http"
	"net/http/httptest"
	"os"

	"day-33/handler"
	"day-33/logger"
	"day-33/middleware"
	"day-33/service"

	"github.com/gin-gonic/gin"
)

func main() {
	fmt.Println("==========================================================")
	fmt.Println("🚀 Day 33: File Upload API in Go (Multipart, MIME, Storage)")
	fmt.Println("==========================================================")

	uploadDir := "./uploads"
	logFilePath := "file_app.log"
	zapLog, err := logger.NewZapLogger("development", logFilePath)
	if err != nil {
		fmt.Printf("Failed to initialize Zap Logger: %v\n", err)
		os.Exit(1)
	}
	defer zapLog.Sync()

	ctx := context.Background()
	zapLog.Info(ctx, "File Upload API Server starting", "upload_dir", uploadDir)

	fileSvc, err := service.NewFileService(uploadDir, 5*1024*1024, zapLog)
	if err != nil {
		fmt.Printf("Failed to initialize File Service: %v\n", err)
		os.Exit(1)
	}

	fileHandler := handler.NewFileHandler(fileSvc, zapLog)

	gin.SetMode(gin.ReleaseMode)
	router := gin.New()
	router.Use(gin.Recovery())
	router.Use(middleware.RequestIDMiddleware())
	router.Use(middleware.StructuredLoggerMiddleware(zapLog))

	v1 := router.Group("/api/v1/files")
	{
		v1.POST("/upload", fileHandler.UploadSingle)
		v1.POST("/upload-multiple", fileHandler.UploadMultiple)
		v1.GET("/download/:filename", fileHandler.DownloadFile)
	}

	fmt.Println("\n--- 1️⃣ Testing Single File Upload (.png) ---")
	uploadSingle(router, "sample_document.png", []byte("fake-image-binary-data"))

	fmt.Println("\n--- 2️⃣ Testing Invalid File Type Upload (.exe) ---")
	uploadSingle(router, "malicious_program.exe", []byte("binary-code"))

	fmt.Println("\n--- 3️⃣ Testing Multiple Files Upload ---")
	uploadMultiple(router, map[string][]byte{
		"report.pdf": []byte("PDF-report-content"),
		"photo.jpg":  []byte("JPEG-image-content"),
	})

	fmt.Println("\n✅ Day 33 File Upload API executed successfully! Check file_app.log for logs.")
}

func uploadSingle(router *gin.Engine, filename string, content []byte) {
	body := &bytes.Buffer{}
	writer := multipart.NewWriter(body)
	part, _ := writer.CreateFormFile("file", filename)
	part.Write(content)
	writer.Close()

	req, _ := http.NewRequest("POST", "/api/v1/files/upload", body)
	req.Header.Set("Content-Type", writer.FormDataContentType())

	w := httptest.NewRecorder()
	router.ServeHTTP(w, req)
	fmt.Printf("POST /upload (%s) -> Status %d | Body: %s\n", filename, w.Code, w.Body.String())
}

func uploadMultiple(router *gin.Engine, files map[string][]byte) {
	body := &bytes.Buffer{}
	writer := multipart.NewWriter(body)
	for fname, content := range files {
		part, _ := writer.CreateFormFile("files", fname)
		part.Write(content)
	}
	writer.Close()

	req, _ := http.NewRequest("POST", "/api/v1/files/upload-multiple", body)
	req.Header.Set("Content-Type", writer.FormDataContentType())

	w := httptest.NewRecorder()
	router.ServeHTTP(w, req)
	fmt.Printf("POST /upload-multiple -> Status %d | Body: %s\n", w.Code, w.Body.String())
}
