package main

import (
	"log"

	"qonvif/apis"
)

func main() {
	err := apis.Run()
	if err != nil {
		log.Fatal(err)
	}
}
