package main

import (
	"RainbowRunner/internal/config"
	"RainbowRunner/internal/lua"
	"RainbowRunner/internal/objects"
)

func main() {
	config.Load()
	//database.LoadConfigFiles()
	err := lua.LoadScripts("./lua")
	if err != nil {
		panic(err)
	}
	objects.Init()

	zone := objects.Zones.CreateZone("town")

	zone.Init()
}
