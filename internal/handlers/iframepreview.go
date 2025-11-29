package handlers

import (
	"net/http"
	"path/filepath"
	"strings"

	"github.com/gin-gonic/gin"
)

func IframePreview(c *gin.Context) {
	filename := c.Param("id")
	
	// Determine file type
	ext := strings.ToLower(filepath.Ext(filename))
	isPDF := ext == ".pdf"
	isHTML := ext == ".html" || ext == ".htm"
	isExcel := ext == ".xlsx" || ext == ".xls"
	
	if isExcel {
		// For Excel files, redirect to the Excel table handler
		// This will load the table inside the iframe
		c.Redirect(http.StatusFound, "/excel/"+filename)
	} else if isPDF {
		c.HTML(http.StatusOK, "iframe.html", gin.H{
			"filename": filename,
			"isPDF":    true,
			"isHTML":   false,
			"isExcel":  false,
		})
	} else if isHTML {
		c.HTML(http.StatusOK, "iframe.html", gin.H{
			"filename": filename,
			"isPDF":    false,
			"isHTML":   true,
			"isExcel":  false,
		})
	} else {
		c.HTML(http.StatusOK, "iframe.html", gin.H{
			"filename": filename,
			"isPDF":    false,
			"isHTML":   false,
			"isExcel":  false,
		})
}
}
