package main

import (
	"RainbowRunner/internal/database"
	"RainbowRunner/internal/gosucks"
	"fmt"
	"github.com/pkg/profile"
)

func main() {
	defer profile.Start(profile.MemProfile, profile.ProfilePath("./tmp")).Stop()

	database.LoadConfigFiles()

	zone, err := database.GetZoneConfig("town")

	if err != nil {
		panic(err)
	}

	for name, npcConfig := range zone.NPCs {
		fmt.Printf("%s: %s\n", name, npcConfig.GetNPCConfig().Behaviour.Type)
	}

	gosucks.VAR(zone)
	//worldConfig := database.LoadWorldConfigs()
	//zonesConfig := database.LoadZoneConfigs()

	//gosucks.VAR(worldConfig)
}
