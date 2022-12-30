package main

import (
	"RainbowRunner/internal/database"
	"RainbowRunner/internal/gosucks"
	"RainbowRunner/internal/serverconfig"
)

func main() {
	serverconfig.Load()
	database.LoadConfigFiles()

	zoneConfig, err := database.GetZoneConfig("test")

	if err != nil {
		panic(err)
	}

	gosucks.VAR(zoneConfig)
}
