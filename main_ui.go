//go:build ui

package main

import (
	"embed"

	"qonvif/ui"
)

//go:embed all:web/dist
var assets embed.FS

func main() {
	ui.Run(assets)
}
