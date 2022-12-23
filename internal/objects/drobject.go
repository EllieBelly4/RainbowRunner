package objects

import (
	"RainbowRunner/pkg/datatypes"
)

type DRItem interface {
	SetInventoryPosition(vector2 datatypes.Vector2)
}
