package objects

import (
	"RainbowRunner/internal/types/configtypes"
	"RainbowRunner/internal/types/drobjecttypes"
)

type IActivatable interface {
	Activate(player *RRPlayer, u *UnitBehavior, id byte, seqID byte)
}

func EntityFromConfig(config *configtypes.EntityConfig) (drobjecttypes.DRObject, error) {
	return nil, nil
}
