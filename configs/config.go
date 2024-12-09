package configs

import (
	"qonvif/configs/models"
	"qonvif/utils"

	"github.com/pelletier/go-toml/v2"
)

const configPath = "configs"

func loadConfig(filename string) (models.Config, error) {
	var config models.Config

	mainRoot, err := utils.FileSuite.RootPath(configPath)
	if err != nil {
		return config, err
	}

	configData, err := utils.FileSuite.GetFileData(mainRoot, filename)
	if err != nil {
		return config, err
	}

	err = toml.Unmarshal(configData, &config)
	if err != nil {
		return config, err
	}
	return config, nil
}
