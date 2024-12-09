package common

import (
	"github.com/gin-gonic/gin"
)

func JSONHandler(c *gin.Context, status int, message string, data any) {
	c.JSON(200, gin.H{
		"status":  status,
		"message": message,
		"data":    data,
	})
}

func AuthAbortHandler(c *gin.Context) {
	c.AbortWithStatus(401)
}

func NoContentHandler(c *gin.Context) {
	c.Status(204)
}
