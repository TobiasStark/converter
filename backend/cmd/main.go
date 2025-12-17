package main

import (
	"fmt"
	"log"
	"os"
	"time"

	"github.com/TobiasStark/converter/backend/config"
	"github.com/TobiasStark/converter/backend/router"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func main() {
	if err := os.MkdirAll(config.UploadDir, config.DefaultDirPermissions); err != nil {
		fmt.Printf("Failed to create upload folder: %v\n", err)
		return
	}

	r := gin.Default()
	r.MaxMultipartMemory = config.MaxFileSize

	r.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"http://localhost:5173"},
		AllowMethods:     []string{"POST", "OPTIONS"},
		AllowHeaders:     []string{"Origin", "Content-Type"},
		ExposeHeaders:    []string{"Content-Length"},
		AllowCredentials: true,
		MaxAge:           12 * time.Hour,
	}))

	router.SetupRouter(r)
	port := ":8080"

	if err := r.Run(port); err != nil {
		log.Fatalf("Failed to start server: %v", err)
	}
}
