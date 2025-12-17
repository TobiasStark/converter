package main

import (
	"fmt"
	"log"
	"os"

	"github.com/TobiasStark/converter/backend/config"
	"github.com/TobiasStark/converter/backend/router"
	"github.com/gin-gonic/gin"
)

func main() {
	if err := os.MkdirAll(config.UploadDir, config.DefaultDirPermissions); err != nil {
		fmt.Printf("Failed to create upload folder: %v\n", err)
		return
	}

	r := gin.Default()
	r.MaxMultipartMemory = config.MaxFileSize
	router.SetupRouter(r)
	port := ":8080"

	if err := r.Run(port); err != nil {
		log.Fatalf("Failed to start server: %v", err)
	}
}
