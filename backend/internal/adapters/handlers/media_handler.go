package handlers

import (
	"fmt"
	"io"
	"net/http"
	"os"
	"path/filepath"
	"time"

	"github.com/google/uuid"
	"github.com/labstack/echo/v4"
)

type mediaHandler struct {
	uploadDir string
}

func NewMediaHandler(uploadDir string) *mediaHandler {
	return &mediaHandler{uploadDir: uploadDir}
}

func (h *mediaHandler) Upload(c echo.Context) error {
	// Source
	file, err := c.FormFile("file")
	if err != nil {
		return err
	}
	src, err := file.Open()
	if err != nil {
		return err
	}
	defer src.Close()

	// Ensure upload directory exists
	if _, err := os.Stat(h.uploadDir); os.IsNotExist(err) {
		os.MkdirAll(h.uploadDir, os.ModePerm)
	}

	// Create unique filename
	ext := filepath.Ext(file.Filename)
	newFilename := fmt.Sprintf("%s-%d%s", uuid.New().String(), time.Now().Unix(), ext)
	dstPath := filepath.Join(h.uploadDir, newFilename)

	// Destination
	dst, err := os.Create(dstPath)
	if err != nil {
		return err
	}
	defer dst.Close()

	// Copy
	if _, err = io.Copy(dst, src); err != nil {
		return err
	}

	// Return relative URL
	url := fmt.Sprintf("/uploads/%s", newFilename)
	return c.JSON(http.StatusOK, map[string]string{
		"url": url,
	})
}
