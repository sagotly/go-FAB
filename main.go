package main

import (
	"go-fab/handlers"

	"github.com/gin-gonic/gin"
)

func main() {
	gin := gin.Default()
	// Настройка обработки статических файлов
	gin.LoadHTMLGlob("templates/*")
	gin.Static("/static", "./static")
	gin.GET("/upload", handlers.UploadHandler)
	gin.POST("/upload", handlers.UploadHandler)
	gin.Run(":4001")
}
