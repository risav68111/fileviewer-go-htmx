package main

import (
	"encoding/json"
	"html/template"
	"log"
	"path/filepath"
	"strings"
	"os"

  "github.com/risav68111/filesviewer-go-htmx/internal/handlers"
	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()

	port := os.Getenv("PORT")
	if port == "" {
			port = "8080"
	}

	log.Println("Server starting on :" + port)
	r.Run(":" + port)


	// Add template functions
	r.SetFuncMap(template.FuncMap{

		"toJson": func(v interface{}) template.JS {
			a, _ := json.MarshalIndent(v, "", "  ")
			return template.JS(a)
		},
		"add": func(a, b int) int {
			return a + b
		},
		"even": func(i int) bool {
			return i%2 == 0
		},
		"pathJoin": func(elements ...string) string {
			return filepath.Join(elements...)
		},
		"getFileExtension": func(filename string) string {
			return strings.TrimPrefix(filepath.Ext(filename), ".")
		},
		"split": func(s string, sep string) []string {
			return strings.Split(s, sep)
		},
		"hasSuffix": func(s, suffix string) bool {
			return strings.HasSuffix(s, suffix)
		},
	})

	r.Static("/static", "./internal/static")
	r.LoadHTMLGlob("internal/templates/*.html")

	// Routes
	r.GET("/", handlers.Home)
	r.GET("/raw/:id", handlers.RawFile)
	r.GET("/excel/:id", handlers.ExcelTable) // This returns HTML table
	r.GET("/download/:id", handlers.Download)
	r.GET("/iframe/:id", handlers.IframePreview)

	log.Println("Server starting on :8080")
	log.Fatal(r.Run(":8080"))
}
