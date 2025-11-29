package handlers

import (
	"fmt"
	"github.com/gin-gonic/gin"
)

func RawFile(c *gin.Context) {
	id := c.Param("id")
	c.File("./files/" + id)
}

func Download(c *gin.Context) {
	id := c.Param("id")

	c.Header("X-Metadata", fmt.Sprintf("downloaded=%s", id))
	c.File("./files/" + id)
}

