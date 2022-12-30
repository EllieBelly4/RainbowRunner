package main

import (
	"RainbowRunner/internal/database"
	"RainbowRunner/internal/gosucks"
	"RainbowRunner/internal/objects"
	"RainbowRunner/internal/serverconfig"
)

func main() {
	serverconfig.Load()
	database.LoadConfigFiles()

	zoneConfig, err := database.GetZoneConfig("town")

	if err != nil {
		panic(err)
	}

	npc := objects.NewNPCFromConfig(zoneConfig.NPCs["snowman1"])

	gosucks.VAR(npc)
}
