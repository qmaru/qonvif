package onvif

import (
	"qonvif/apis/common"

	"github.com/gin-gonic/gin"
)

func AuthCheck(c *gin.Context) {
	common.NoContentHandler(c)
}
