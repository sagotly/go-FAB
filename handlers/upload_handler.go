package handlers

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
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

		fileName := fmt.Sprintf("%s.txt", uuid.New().String())
		fullPath := "C:\\Users\\User\\Desktop\\go-fab\\data\\" + fileName

		err = c.SaveUploadedFile(file, fullPath)
		if err != nil {
			fmt.Printf("error has accured:")
			fmt.Print(err)
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}

		c.JSON(http.StatusOK, gin.H{"message": "Файл успешно загружен"})
		fmt.Println("konec")
		return
	}
}
