package main

import (
	"dependencies/app"
	"dependencies/config"
	"fmt"
	"os"
)

const dev = "development"
const production = "production"

func main() {

	cfg, err := config.CrateConfig(dev)
	if err != nil {
		fmt.Println("Error reading config file", err)
		os.Exit(1)
	}

	app.Run(cfg)
}
