package onvif

import (
	"qonvif/apis/common"
	"qonvif/services/onvif"
	"qonvif/services/onvif/models"

	"github.com/gin-gonic/gin"
)

func handleMove(c *gin.Context, moveFunc func(client *onvif.OnvifBasic, x, y, z float64) (any, error), message string) {
	var ptzControl models.PtzControl

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

	status, err := moveFunc(client, ptzControl.Axes.X, ptzControl.Axes.Y, ptzControl.Axes.Z)
	if err != nil {
		common.JSONHandler(c, 0, "move error: "+err.Error(), []any{})
		return
	}

	common.JSONHandler(c, 1, message, status)
}

func Status(c *gin.Context) {
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

	status, err := client.PTZStatus()
	if err != nil {
		common.JSONHandler(c, 0, "move error: "+err.Error(), []any{})
		return
	}

	common.JSONHandler(c, 1, "control", status)
}

func RelativeMove(c *gin.Context) {
	handleMove(c, func(client *onvif.OnvifBasic, x, y, z float64) (any, error) {
		return client.PTZGoToAnyRelative(x, y, z)
	}, "relative move")
}

func AbsoluteMove(c *gin.Context) {
	handleMove(c, func(client *onvif.OnvifBasic, x, y, z float64) (any, error) {
		return client.PTZGoToAnyAbsolute(x, y, z)
	}, "absolute move")
}
