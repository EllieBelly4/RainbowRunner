package main

import (
	"RainbowRunner/internal/admin"
	"RainbowRunner/internal/api"
	"RainbowRunner/internal/config"
	"RainbowRunner/internal/database"
	"RainbowRunner/internal/game"
	"RainbowRunner/internal/logging"
	"RainbowRunner/internal/login"
	"RainbowRunner/internal/lua"
	"RainbowRunner/internal/objects"
	"github.com/pkg/profile"
)

var done = make(chan bool)

func main() {
	defer profile.Start(profile.ProfilePath("./tmp")).Stop()

	config.Load()
	logging.Init()
	err := lua.LoadScripts("./lua")

	if err != nil {
		panic(err)
	}

	database.LoadEquipmentFixtures()
	database.LoadConfigFiles()

	go login.StartLoginServer()
	go game.StartGameServer()
	go admin.StartAdminServer()
	go api.StartGraphqlAPI()

	objects.Init()

	for !<-done {

	}
}
