package configs

import (
	"log"

	"qonvif/configs/models"
)

var Config models.Config

func init() {
	cfg, err := loadConfig("config.toml")
	if err != nil {
		log.Fatal(err)
	}

	Config = cfg
}
