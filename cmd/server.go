package cmd

import (
	"qonvif/apis"
)

type ServerCommand struct{}

func (c *ServerCommand) Execute(args []string) error {
	err := apis.Run()
	if err != nil {
		return err
	}
	return nil
}
