package player

import (
	"os/exec"

	"qonvif/configs"
)

type PlayParas struct {
	Url    string `json:"url"`
	Width  string `json:"width"`
	Height string `json:"height"`
}

func Open(para *PlayParas) error {
	execPath := configs.Config.Player.Path
	basicCmd := []string{"-i", para.Url, "-x", para.Width, "-y", para.Height, "-autoexit"}
	cmd := exec.Command(execPath, basicCmd...)

	err := cmd.Run()
	if err != nil {
		return err
	}
	return nil
}
