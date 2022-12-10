package commands

import "RainbowRunner/internal/objects"

type ChatCommandHandler func(player *objects.RRPlayer, args []string)
