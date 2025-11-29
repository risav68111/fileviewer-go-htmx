package handlers

import (
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
)

func Home(c *gin.Context) {
	filesDir := "./files"
	
	// Read files from directory
	files, err := os.ReadDir(filesDir)
	if err != nil {
		c.HTML(http.StatusOK, "index.html", gin.H{
			"files": []string{},
		})
		return
	}

	var fileNames []string
	for _, file := range files {
		if !file.IsDir() {
			fileNames = append(fileNames, file.Name())
		}
	}

	c.HTML(http.StatusOK, "index.html", gin.H{
		"files": fileNames,
	})
}
