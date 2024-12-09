package cmd

import (
	"fmt"
	"os"

	"qonvif/utils"

	"github.com/jessevdk/go-flags"
)

type Option struct {
	Version func()        `short:"v" long:"version" description:"Show version"`
	Server  ServerCommand `command:"server"`
}

func Execute() {
	var opts Option

	parser := flags.NewParser(&opts, flags.Default)
	parser.Name = "qonvif"
	parser.LongDescription = "qonvif is a easy onvif client"

	if len(os.Args) == 1 {
		parser.WriteHelp(os.Stdout)
		return
	}

	opts.Version = func() {
		fmt.Printf("%s version %s\n", parser.Name, utils.Version)
		os.Exit(0)
	}

	_, err := parser.Parse()
	if err != nil {
		os.Exit(0)
	}
}
