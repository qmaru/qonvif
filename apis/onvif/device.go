package onvif

import (
	URL "net/url"
	"qonvif/apis/common"
	"qonvif/configs"
	"qonvif/services/onvif"

	"github.com/gin-gonic/gin"
)

func addAuthtoUrl(url, username, password string) string {
	if username == "" && password == "" {
		return url
	}

	urlParse, err := URL.Parse(url)
	if err != nil {
		return ""
	}

	urlParse.User = URL.UserPassword(username, password)
	return urlParse.String()
}

func ListDevices(c *gin.Context) {
	devices := make([]map[string]any, 0)
	for _, device := range configs.Config.Devices {
		client, err := onvif.NewClient(device.Name)
		if err != nil {
			common.JSONHandler(c, 0, "new client error: "+err.Error(), []any{})
			return
		}
		data, err := client.GetDeviceData()
		if err != nil {
			common.JSONHandler(c, 0, "get device info error: "+err.Error(), []any{})
			return
		}

		dev := map[string]any{
			"profile": device,
			"details": data,
		}
		devices = append(devices, dev)
	}

	common.JSONHandler(c, 1, "devices", devices)
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
	username := c.DefaultQuery("username", "")
	password := c.DefaultQuery("password", "")

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

	streamData.Url = addAuthtoUrl(streamData.Url, username, password)
	common.JSONHandler(c, 1, "device stream url", streamData)
}
