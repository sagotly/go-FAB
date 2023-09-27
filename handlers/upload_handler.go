package handlers

import (
	"fmt"
	"go-fab/utils"
	"net/http"

	"github.com/gin-gonic/gin"
)

func UploadHandler(c *gin.Context) {
	if c.Request.Method == "GET" {
		c.HTML(200, "upload.html", gin.H{})
	} else if c.Request.Method == "POST" {
		fmt.Print("getting Files...\n")

		file, err := c.FormFile("file")
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
		fmt.Print(file.Filename)
		err = c.SaveUploadedFile(file, utils.GenerateUniqueFileName(file.Filename))
		if err != nil {
			fmt.Printf("error has accured:%s", err)
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}

		c.JSON(http.StatusOK, gin.H{"message": "Файл успешно загружен"})
		fmt.Println("konec")
		return
	}
}
