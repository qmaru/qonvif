package middle

import (
	"qonvif/apis/common"
	"qonvif/configs"

	"github.com/gin-gonic/gin"
)

var (
	ApiKey = configs.Config.Server.ApiKey
)

func ApiKeyAuth() gin.HandlerFunc {
	return func(c *gin.Context) {
		apiKey := c.GetHeader("X-API-Key")
		if apiKey == "" {
			common.AuthAbortHandler(c)
			return
		}

		if apiKey != ApiKey {
			common.AuthAbortHandler(c)
			return
		}

		c.Next()
	}
}
