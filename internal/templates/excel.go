package handlers

import (
	"fmt"
	"net/http"
	"path/filepath"
	"strings"

	"filesviewer/internal/services"
	"github.com/gin-gonic/gin"
)

func ExcelTable(c *gin.Context) {
	filename := c.Param("id")
	filePath := "./files/" + filename
	
	// Check if client wants JSON format
	acceptHeader := c.GetHeader("Accept")
	wantsJSON := strings.Contains(acceptHeader, "application/json") || 
	             c.Query("format") == "json" ||
	             strings.Contains(c.GetHeader("HX-Request"), "true")

	if wantsJSON {
		// Return JSON format
		sheets, err := services.GetExcelTableData(filePath)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{
				"error": fmt.Sprintf("Error processing Excel file: %v", err),
			})
			return
		}
		c.JSON(http.StatusOK, gin.H{
			"filename": filename,
			"sheets":   sheets,
		})
	} else {
		// Return simple table format (backward compatibility)
		rows, err := services.XLSXToTable(filePath)
		if err != nil {
			c.HTML(http.StatusOK, "table.html", gin.H{
				"Rows": [][]string{},
			})
			return
		}

		c.HTML(http.StatusOK, "table.html", gin.H{
			"Rows": rows,
		})
	}
}

// New endpoint for enhanced JSON table view
func ExcelTableEnhanced(c *gin.Context) {
	filename := c.Param("id")
	filePath := "./files/" + filename

	sheets, err := services.GetExcelTableData(filePath)
	if err != nil {
		c.HTML(http.StatusOK, "excel_table.html", gin.H{
			"filename": filename,
			"error":    fmt.Sprintf("Error processing Excel file: %v", err),
			"sheets":   []services.ExcelSheet{},
		})
		return
	}

	c.HTML(http.StatusOK, "excel_table.html", gin.H{
		"filename": filename,
		"sheets":   sheets,
		"error":    "",
	})
}

// JSON endpoint
func ExcelJSON(c *gin.Context) {
	filename := c.Param("id")
	filePath := "./files/" + filename

	jsonData, err := services.XLSXToJSONString(filePath)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": fmt.Sprintf("Error converting Excel to JSON: %v", err),
		})
		return
	}

	c.Header("Content-Type", "application/json")
	c.String(http.StatusOK, jsonData)
}
