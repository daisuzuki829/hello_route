package routes

import (
	"github.com/gin-gonic/gin"
)

// HelloJson...
func HelloJson(c *gin.Context) {
	// デフォルト値 HOGE
	name := c.DefaultQuery("name", "Michael")
	c.JSON(200, gin.H{
		"name": name,
	})
}

// HelloJsonName...
func HelloJsonName(c *gin.Context) {
	name := c.Param("name")
	c.JSON(200, gin.H{
		"name": name,
	})
}
