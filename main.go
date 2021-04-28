package main

import (
	"RainbowRunner/internal/game"
	"RainbowRunner/internal/login"
)

var done = make(chan bool)

func main() {
	go login.StartLoginServer()
	go game.StartGameServer()

	for !<-done {

	}
}
