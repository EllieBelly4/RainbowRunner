package commands

import "RainbowRunner/internal/objects"

type ChatCommandAliasHandler func(player *objects.RRPlayer, args []string) []string

func AliasCustom(handler ChatCommandHandler, aliasHandler ChatCommandAliasHandler) ChatCommandHandler {
	return func(player *objects.RRPlayer, args []string) {
		newArgs := aliasHandler(player, args)
		handler(player, newArgs)
	}
}
