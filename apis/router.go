package apis

import (
	"fmt"
	"log"

	"qonvif/apis/middle"
	"qonvif/apis/onvif"
	"qonvif/apis/player"
	"qonvif/configs"
	"qonvif/services/logs"

	"github.com/gin-contrib/cors"
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

	listenAddr := fmt.Sprintf("%s:%d", configs.Config.Server.Host, configs.Config.Server.Port)
	log.Println("Listen: " + listenAddr)

	router := gin.New()
	router.SetTrustedProxies(nil)
	router.Use(gin.Recovery())
	router.Use(logger)

	corsConfig := cors.Config{
		AllowAllOrigins: true,
		AllowMethods:    cors.DefaultConfig().AllowMethods,
		AllowHeaders:    append(cors.DefaultConfig().AllowHeaders, "X-API-Key"),
	}
	router.Use(cors.New(corsConfig))

	api := router.Group("/api/onvif")
	{
		api.GET("/devices", middle.ApiKeyAuth(), onvif.ListDevices)
		api.GET("/device/info", middle.ApiKeyAuth(), onvif.ListDeviceInfo)
		api.GET("/device/profile", middle.ApiKeyAuth(), onvif.ListDeviceProfile)
		api.GET("/device/streamurl", middle.ApiKeyAuth(), onvif.ListDeviceStreamurl)
		api.POST("/device/ptz/control", middle.ApiKeyAuth(), onvif.DeviceControl)
		api.POST("/play", middle.ApiKeyAuth(), player.PlayStram)
	}

	return router.Run(listenAddr)
}
