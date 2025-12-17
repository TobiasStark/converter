package router

import (
	"github.com/TobiasStark/converter/backend/handler"
	"github.com/gin-gonic/gin"
)

func SetupRouter(r *gin.Engine) {
	v1 := r.Group("/api/v1")
	{
		v1.POST("/upload", handler.UploadFileHandler)
	}
}
