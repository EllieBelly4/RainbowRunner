package main

import (
	"RainbowRunner/internal/config"
	"RainbowRunner/internal/database"
	"RainbowRunner/internal/gosucks"
)

func main() {
	config.Load()
	database.LoadConfigFiles()

	zoneConfig, err := database.GetZoneConfig("town")

	if err != nil {
		panic(err)
	}

	gosucks.VAR(zoneConfig)
}
