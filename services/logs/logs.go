package logs

import (
	"path/filepath"

	"qonvif/configs"
	"qonvif/utils"

	"github.com/gin-gonic/gin"
	"github.com/qmaru/qlog/ginlog"
)

func genLogFile(logfile string) (string, error) {
	if configs.Config.Debug {
		return "", nil
	}

	logPath, err := utils.FileSuite.RootPath("logs")
	if err != nil {
		return "", err
	}
	logpath, err := utils.FileSuite.Mkdir(logPath)
	if err != nil {
		return "", err
	}
	accessPath := filepath.Join(logpath, logfile)
	return accessPath, nil
}

func NewGinLogger(filename string) (gin.HandlerFunc, error) {
	output, err := genLogFile(filename)
	if err != nil {
		return nil, err
	}
	return ginlog.Logger(output)
}
