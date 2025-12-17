package handler

import (
	"fmt"
	"net/http"
	"path/filepath"
	"time"

	"github.com/TobiasStark/converter/backend/config"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

func UploadFileHandler(c *gin.Context) {
	file, err := c.FormFile("file")
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": fmt.Sprintf("Failing to read file: %s", err.Error()),
		})
		return
	}

	if file.Size > config.MaxFileSize {
		c.JSON(http.StatusRequestEntityTooLarge, gin.H{
			"error": fmt.Sprintf("File too large. Max file size: %d MiB", config.MaxFileSize>>20),
		})
		return
	}

	jobID := uuid.New().String()

	fileExtension := filepath.Ext(file.Filename)
	uniqueFilename := fmt.Sprintf("%s%s", jobID, fileExtension)

	savePath := filepath.Join(config.UploadDir, uniqueFilename)

	if err := c.SaveUploadedFile(file, savePath); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": fmt.Sprintf("Failing to save file: %s", err.Error()),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message":    "Successfully uploaded and saved file",
		"job_id":     jobID,
		"filename":   file.Filename,
		"saved_as":   uniqueFilename,
		"size_bytes": file.Size,
		"timestamp":  time.Now(),
	})
}
