package onvif

import (
	"qonvif/apis/common"
	"qonvif/configs"
	"qonvif/services/onvif"

	"github.com/gin-gonic/gin"
)

type PtzAxes struct {
	X float64 `json:"x"`
	Y float64 `json:"y"`
}

func ListDevices(c *gin.Context) {
	data := make([]string, 0)
	for _, device := range configs.Config.Devices {
		data = append(data, device.Name)
	}

	common.JSONHandler(c, 1, "devices", data)
}

func ListDeviceInfo(c *gin.Context) {
	name := c.DefaultQuery("name", "")
	if name == "" {
		common.JSONHandler(c, 0, "name not found", []any{})
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
		common.JSONHandler(c, 0, "name not found", []any{})
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
		common.JSONHandler(c, 0, "name or token not found", []any{})
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

func DeviceControl(c *gin.Context) {
	var axes PtzAxes

	name := c.DefaultQuery("name", "")
	if name == "" {
		common.JSONHandler(c, 0, "name not found", []any{})
		return
	}

	err := c.ShouldBindBodyWithJSON(&axes)
	if err != nil {
		common.JSONHandler(c, 0, "x or y error: "+err.Error(), []any{})
		return
	}

	client, err := onvif.NewClient(name)
	if err != nil {
		common.JSONHandler(c, 0, "new client error: "+err.Error(), []any{})
		return
	}

	status, err := client.PTZGoToAny(axes.X, axes.Y)
	if err != nil {
		common.JSONHandler(c, 0, "move error: "+err.Error(), []any{})
		return
	}

	common.JSONHandler(c, 1, "control", status)
}
