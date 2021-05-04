package main

import (
	"RainbowRunner/internal/game"
	"RainbowRunner/internal/logging"
	"RainbowRunner/internal/login"
)

var done = make(chan bool)

func main() {
	logging.Init()

	go login.StartLoginServer()
	go game.StartGameServer()

	for !<-done {

	}
}
