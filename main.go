package main

import (
	"RainbowRunner/internal/admin"
	"RainbowRunner/internal/database"
	"RainbowRunner/internal/game"
	"RainbowRunner/internal/logging"
	"RainbowRunner/internal/login"
	"RainbowRunner/internal/objects"
)

var done = make(chan bool)

func main() {
	logging.Init()

	database.LoadEquipmentFixtures()

	go login.StartLoginServer()
	go game.StartGameServer()
	go admin.StartAdminServer()
	objects.Init()

	for !<-done {

	}
}
