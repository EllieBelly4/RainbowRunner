package commands

import (
	"RainbowRunner/internal/lua"
	"RainbowRunner/internal/objects"
)

func ReloadLua(player *objects.RRPlayer, args []string) {
	err := lua.LoadScripts("./lua")

	if err != nil {
		panic(err)
	}
}
