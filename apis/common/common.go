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
