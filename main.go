package main

import (
	"RainbowRunner/internal/admin"
	"RainbowRunner/internal/api"
	"RainbowRunner/internal/database"
	"RainbowRunner/internal/game"
	"RainbowRunner/internal/logging"
	"RainbowRunner/internal/login"
	"RainbowRunner/internal/lua"
	"RainbowRunner/internal/objects"
	"RainbowRunner/internal/serverconfig"
	"flag"
	"github.com/pkg/profile"
)

var done = make(chan bool)

var (
	profiledEnabled = flag.Bool("profile", false, "enable profiling")
)

func main() {
	if *profiledEnabled {
		defer profile.Start(profile.ProfilePath("./tmp")).Stop()
	}

	serverconfig.Load()
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
