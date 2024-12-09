package onvif

import (
	"qonvif/apis/common"
	"qonvif/configs"
	"qonvif/services/onvif"

	"github.com/gin-gonic/gin"
)

func ListDevices(c *gin.Context) {
	common.JSONHandler(c, 1, "devices", configs.Config.Devices)
}

func ListDeviceInfo(c *gin.Context) {
	name := c.DefaultQuery("name", "")
	if name == "" {
		common.JSONHandler(c, 0, "device name not found", []any{})
		return
	}

	client, err := onvif.NewClient(name)
	if err != nil {
		common.JSONHandler(c, 0, "new client error: "+err.Error(), []any{})
		return
	}

	data, err := client.GetDeviceData()
	if err != nil {
		common.JSONHandler(c, 0, "get device info error: "+err.Error(), []any{})
		return
	}

	common.JSONHandler(c, 1, "device info", data)
}

func ListDeviceProfile(c *gin.Context) {
	name := c.DefaultQuery("name", "")
	if name == "" {
		common.JSONHandler(c, 0, "device name not found", []any{})
		return
	}

	client, err := onvif.NewClient(name)
	if err != nil {
		common.JSONHandler(c, 0, "new client error: "+err.Error(), []any{})
		return
	}

	profile, err := client.GetProfiles()
	if err != nil {
		common.JSONHandler(c, 0, "new client error: "+err.Error(), []any{})
		return
	}

	common.JSONHandler(c, 1, "device profile", profile)
}

func ListDeviceStreamurl(c *gin.Context) {
	name := c.DefaultQuery("name", "")
	token := c.DefaultQuery("token", "")

	if name == "" || token == "" {
		common.JSONHandler(c, 0, "device name or token not found", []any{})
		return
	}

	client, err := onvif.NewClient(name)
	if err != nil {
		common.JSONHandler(c, 0, "new client error: "+err.Error(), []any{})
		return
	}

	streamData, err := client.GetStreamUri(token)
	if err != nil {
		common.JSONHandler(c, 0, "new client error: "+err.Error(), []any{})
		return
	}

	common.JSONHandler(c, 1, "device stream url", streamData)
}
