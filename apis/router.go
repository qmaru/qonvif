package apis

import (
	"fmt"
	"log"

	"qonvif/apis/middle"
	"qonvif/apis/onvif"
	"qonvif/configs"
	"qonvif/services/logs"

	"github.com/gin-gonic/gin"
)

func init() {
	if configs.Config.Debug {
		gin.SetMode(gin.DebugMode)
	} else {
		gin.SetMode(gin.ReleaseMode)
	}
}

func Run() error {
	logger, err := logs.NewGinLogger("access.log")
	if err != nil {
		return err
	}

	basicAuth, err := middle.BasicAuth()
	if err != nil {
		return err
	}

	listenAddr := fmt.Sprintf("%s:%d", configs.Config.Server.Host, configs.Config.Server.Port)
	log.Println("Listen: " + listenAddr)

	router := gin.New()
	router.SetTrustedProxies(nil)
	router.Use(gin.Recovery())
	router.Use(logger)
	router.Use(gin.Recovery())

	api := router.Group("/api/onvif")
	{
		api.GET("/devices", basicAuth, onvif.ListDevices)
		api.GET("/device/info", basicAuth, onvif.ListDeviceInfo)
		api.GET("/device/profile", basicAuth, onvif.ListDeviceProfile)
		api.GET("/device/streamurl", basicAuth, onvif.ListDeviceStreamurl)
		api.POST("/device/ptz/control", basicAuth, onvif.DeviceControl)
	}

	return router.Run(listenAddr)
}
