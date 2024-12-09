package onvif

import (
	"qonvif/apis/common"
	"qonvif/services/onvif"

	"github.com/gin-gonic/gin"
)

type PtzAxes struct {
	X float64 `json:"x"`
	Y float64 `json:"y"`
}

type PtzControl struct {
	Name string  `json:"name"`
	Axes PtzAxes `json:"axes"`
}

func DeviceControl(c *gin.Context) {
	var ptzControl PtzControl

	err := c.ShouldBindBodyWithJSON(&ptzControl)
	if err != nil {
		common.JSONHandler(c, 0, "control data error: "+err.Error(), []any{})
		return
	}

	if ptzControl.Name == "" {
		common.JSONHandler(c, 0, "device name not found", []any{})
		return
	}

	client, err := onvif.NewClient(ptzControl.Name)
	if err != nil {
		common.JSONHandler(c, 0, "new client error: "+err.Error(), []any{})
		return
	}

	status, err := client.PTZGoToAny(ptzControl.Axes.X, ptzControl.Axes.Y)
	if err != nil {
		common.JSONHandler(c, 0, "move error: "+err.Error(), []any{})
		return
	}

	common.JSONHandler(c, 1, "control", status)
}