package middle

import (
	"fmt"

	"qonvif/configs"

	"github.com/gin-gonic/gin"
)

var (
	username = configs.Config.Server.Username
	password = configs.Config.Server.Password
)

func BasicAuth() (gin.HandlerFunc, error) {
	if username == "" || password == "" {
		return nil, fmt.Errorf("auth username or password is empty")
	}

	return gin.BasicAuth(gin.Accounts{
		username: password,
	}), nil
}
