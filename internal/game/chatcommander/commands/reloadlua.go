package commands

import (
	"RainbowRunner/internal/lua"
	"RainbowRunner/internal/objects"
	log "github.com/sirupsen/logrus"
)

type scriptReloader interface {
	ReloadScripts() error
}

func ReloadLua(player *objects.RRPlayer, args []string) {
	err := lua.LoadScripts("./lua")

	if err != nil {
		panic(err)
	}

	for _, zone := range objects.Zones.GetZones() {
		if !zone.Initialised() {
			continue
		}

		err := zone.ReloadScripts()
		if err != nil {
			log.Error(err)
		}

		for _, entity := range zone.Entities() {
			if reloader, ok := entity.(scriptReloader); ok {
				err := reloader.ReloadScripts()
				if err != nil {
					log.Error(err)
				}
			}
		}
	}
}
