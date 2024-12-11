package ui

import (
	"context"

	"qonvif/configs"
	"qonvif/services/onvif"
	"qonvif/services/onvif/models"
	"qonvif/services/player"
)

// App struct
type App struct {
	ctx context.Context
}

// NewApp creates a new App application struct
func NewApp() *App {
	return &App{}
}

// startup is called when the app starts. The context is saved
// so we can call the runtime methods
func (a *App) Startup(ctx context.Context) {
	a.ctx = ctx
}

func (a *App) ApiAuthCheck(apikey string) JsonData {
	if apikey == configs.Config.Server.ApiKey {
		return JSONHandler(1, "ok", []any{})
	}
	return JSONHandler(0, "failed", []any{})
}

func (a *App) ApiOnvifDevices(apikey string) JsonData {
	devices := make([]map[string]any, 0)
	for _, device := range configs.Config.Devices {
		client, err := onvif.NewClient(device.Name)
		if err != nil {
			return JSONHandler(0, "new client error: "+err.Error(), []any{})
		}

		data, err := client.GetDeviceData()
		if err != nil {
			return JSONHandler(0, "get device info error: "+err.Error(), []any{})
		}

		dev := map[string]any{
			"profile": device,
			"details": data,
		}
		devices = append(devices, dev)
	}

	return JSONHandler(1, "devices", devices)
}

func (a *App) ApiOnvifDeviceProfile(apikey, name string) JsonData {
	if name == "" {
		return JSONHandler(0, "device name not found", []any{})
	}

	client, err := onvif.NewClient(name)
	if err != nil {
		return JSONHandler(0, "new client error: "+err.Error(), []any{})
	}

	profile, err := client.GetProfiles()
	if err != nil {
		return JSONHandler(0, "new client error: "+err.Error(), []any{})
	}

	return JSONHandler(1, "device profile", profile)
}

func (a *App) ApiOnvifDeviceStreamurl(apikey, name, token, username, password string) JsonData {
	if name == "" || token == "" {
		return JSONHandler(0, "device name or token not found", []any{})
	}

	client, err := onvif.NewClient(name)
	if err != nil {
		return JSONHandler(0, "new client error: "+err.Error(), []any{})
	}

	streamData, err := client.GetStreamUri(token)
	if err != nil {
		return JSONHandler(0, "new client error: "+err.Error(), []any{})
	}

	streamData.Url = addAuthtoUrl(streamData.Url, username, password)
	return JSONHandler(1, "device stream url", streamData)
}

func (a *App) ApiOnvifPlay(apikey string, playParas player.PlayParas) JsonData {
	if playParas.Url == "" {
		return JSONHandler(0, "url not found", []any{})
	}
	go player.Open(&playParas)
	return JSONHandler(1, "Start", []any{})
}

func (a *App) ApiOnvifDevicePtzStatus(apikey, name string) JsonData {
	if name == "" {
		return JSONHandler(0, "device name not found", []any{})
	}

	client, err := onvif.NewClient(name)
	if err != nil {
		return JSONHandler(0, "new client error: "+err.Error(), []any{})
	}

	status, err := client.PTZStatus()
	if err != nil {
		return JSONHandler(0, "move error: "+err.Error(), []any{})
	}

	return JSONHandler(1, "control", status)
}

func (a *App) ApiOnvifDevicePtzMoveRelative(apikey string, ptzControl models.PtzControl) JsonData {
	if ptzControl.Name == "" {
		return JSONHandler(0, "device name not found", []any{})
	}

	client, err := onvif.NewClient(ptzControl.Name)
	if err != nil {
		return JSONHandler(0, "new client error: "+err.Error(), []any{})
	}

	status, err := client.PTZGoToAnyRelative(ptzControl.Axes.X, ptzControl.Axes.Y, ptzControl.Axes.Z)
	if err != nil {
		return JSONHandler(0, "move error: "+err.Error(), []any{})
	}

	return JSONHandler(1, "relative move", status)
}

func (a *App) ApiOnvifDevicePtzMoveAbsolute(apikey string, ptzControl models.PtzControl) JsonData {
	if ptzControl.Name == "" {
		return JSONHandler(0, "device name not found", []any{})
	}

	client, err := onvif.NewClient(ptzControl.Name)
	if err != nil {
		return JSONHandler(0, "new client error: "+err.Error(), []any{})
	}

	status, err := client.PTZGoToAnyAbsolute(ptzControl.Axes.X, ptzControl.Axes.Y, ptzControl.Axes.Z)
	if err != nil {
		return JSONHandler(0, "move error: "+err.Error(), []any{})
	}

	return JSONHandler(1, "absolute move", status)
}
