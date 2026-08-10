package service

import (
	"context"
	"fmt"
	"io"
	"mime/multipart"
	"os"
	"path/filepath"
	"strings"
	"time"

	"day-33/domain"

	"github.com/google/uuid"
)

type FileService struct {
	uploadDir string
	maxSize   int64
	logger    domain.Logger
}

func NewFileService(uploadDir string, maxSize int64, logger domain.Logger) (*FileService, error) {
	if err := os.MkdirAll(uploadDir, 0755); err != nil {
		return nil, fmt.Errorf("failed to create upload directory: %w", err)
	}
	return &FileService{
		uploadDir: uploadDir,
		maxSize:   maxSize,
		logger:    logger,
	}, nil
}

func (s *FileService) ValidateFile(fileHeader *multipart.FileHeader) error {
	if fileHeader.Size > s.maxSize {
		return domain.ErrFileTooLarge
	}

	ext := strings.ToLower(filepath.Ext(fileHeader.Filename))
	allowedExts := map[string]bool{
		".png":  true,
		".jpg":  true,
		".jpeg": true,
		".pdf":  true,
		".txt":  true,
	}

	if !allowedExts[ext] {
		return domain.ErrInvalidFileType
	}
	return nil
}

func (s *FileService) SaveUploadedFile(ctx context.Context, fileHeader *multipart.FileHeader) (*domain.FileMeta, error) {
	if err := s.ValidateFile(fileHeader); err != nil {
		s.logger.Warn(ctx, "File validation failed", "filename", fileHeader.Filename, "error", err)
		return nil, err
	}

	file, err := fileHeader.Open()
	if err != nil {
		return nil, fmt.Errorf("failed to open uploaded file: %w", err)
	}
	defer file.Close()

	ext := filepath.Ext(fileHeader.Filename)
	uniqueFilename := fmt.Sprintf("%s_%s%s", time.Now().Format("20060102_150405"), uuid.New().String()[:8], ext)
	dstPath := filepath.Join(s.uploadDir, uniqueFilename)

	dstFile, err := os.Create(dstPath)
	if err != nil {
		return nil, fmt.Errorf("failed to create destination file: %w", err)
	}
	defer dstFile.Close()

	writtenBytes, err := io.Copy(dstFile, file)
	if err != nil {
		return nil, fmt.Errorf("failed to write file to disk: %w", err)
	}

	fileMeta := &domain.FileMeta{
		ID:           uuid.New().String(),
		OriginalName: fileHeader.Filename,
		StoredName:   uniqueFilename,
		ContentType:  fileHeader.Header.Get("Content-Type"),
		SizeBytes:    writtenBytes,
		URL:          fmt.Sprintf("/api/v1/files/download/%s", uniqueFilename),
		UploadedAt:   time.Now(),
	}

	s.logger.Info(ctx, "File uploaded successfully", "stored_name", uniqueFilename, "bytes", writtenBytes)
	return fileMeta, nil
}

func (s *FileService) SaveMultipleFiles(ctx context.Context, fileHeaders []*multipart.FileHeader) ([]*domain.FileMeta, error) {
	results := make([]*domain.FileMeta, 0, len(fileHeaders))
	for _, fh := range fileHeaders {
		meta, err := s.SaveUploadedFile(ctx, fh)
		if err != nil {
			return nil, err
		}
		results = append(results, meta)
	}
	return results, nil
}

func (s *FileService) GetFilePath(filename string) string {
	return filepath.Join(s.uploadDir, filepath.Clean(filename))
}
