package main

import (
	"RainbowRunner/internal/admin"
	"RainbowRunner/internal/api"
	"RainbowRunner/internal/config"
	"RainbowRunner/internal/database"
	"RainbowRunner/internal/game"
	"RainbowRunner/internal/logging"
	"RainbowRunner/internal/login"
	"RainbowRunner/internal/objects"
)

var done = make(chan bool)

func main() {
	config.Load()
	logging.Init()

	database.LoadEquipmentFixtures()

	go login.StartLoginServer()
	go game.StartGameServer()
	go admin.StartAdminServer()
	go api.StartGraphqlAPI()

	objects.Init()

	for !<-done {

	}
}
