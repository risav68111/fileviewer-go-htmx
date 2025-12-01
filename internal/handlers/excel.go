package handlers

import (
	"fmt"
	"net/http"

	"github.com/risav68111/filesviewer-go-htmx/internal/services"
	"github.com/gin-gonic/gin"
)

func ExcelTable(c *gin.Context) {
	filename := c.Param("id")
	filePath := "./files/" + filename

	// Convert Excel to JSON
	sheets, err := services.ExcelToJSON(filePath)
	if err != nil {
		c.HTML(http.StatusOK, "excel_table.html", gin.H{
			"filename": filename,
			"error":    fmt.Sprintf("Error processing Excel file: %v", err),
			"sheets":   []services.ExcelSheet{},
		})
		return
	}

	// Stream JSON data to table format
	c.HTML(http.StatusOK, "excel_table.html", gin.H{
		"filename": filename,
		"sheets":   sheets,
		"error":    "",
	})
}
