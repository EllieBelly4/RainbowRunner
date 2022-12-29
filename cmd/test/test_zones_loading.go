package main

import (
	"RainbowRunner/internal/database"
	"RainbowRunner/internal/gosucks"
	"github.com/pkg/profile"
)

func main() {
	defer profile.Start(profile.ProfilePath("./tmp")).Stop()

	database.LoadConfigFiles()

	zone, err := database.GetZoneConfig("town")

	if err != nil {
		panic(err)
	}

	gosucks.VAR(zone)
	//worldConfig := database.LoadWorldConfigs()
	//zonesConfig := database.LoadZoneConfigs()

	//gosucks.VAR(worldConfig)
}
