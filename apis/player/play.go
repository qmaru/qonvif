package player

import (
	"qonvif/apis/common"
	"qonvif/services/player"

	"github.com/gin-gonic/gin"
)

func PlayStram(c *gin.Context) {
	var playParas player.PlayParas

	err := c.ShouldBindBodyWithJSON(&playParas)
	if err != nil {
		common.JSONHandler(c, 0, "paras error: "+err.Error(), []any{})
		return
	}

	if playParas.Url == "" {
		common.JSONHandler(c, 0, "url not found", []any{})
		return
	}

	go player.Open(&playParas)

	common.JSONHandler(c, 1, "Start", []any{})
}
