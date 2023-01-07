package synchronisation

import (
	"RainbowRunner/internal/actions"
	"RainbowRunner/internal/objects"
	"RainbowRunner/internal/types"
	"RainbowRunner/pkg/byter"
	"RainbowRunner/pkg/datatypes"
)

var playerPosMap = make(map[int]datatypes.Vector3Float32)

func Tick() {
	for _, zone := range objects.Zones.GetZones() {
		synchronisePlayerMovement(zone)
	}
}

func synchronisePlayerMovement(zone *objects.Zone) {
	for _, player := range zone.Players() {
		oldPos, ok := playerPosMap[player.Conn.GetID()]
		behaviour := player.CurrentCharacter.GetPlayer().GetAvatar().GetUnitBehaviour()
		newPos := behaviour.Position

		if !ok {
			playerPosMap[player.Conn.GetID()] = newPos
			continue
		}

		if oldPos.X != newPos.X || oldPos.Y != newPos.Y || oldPos.Z != newPos.Z {
			moveTo := actions.NewActionMoveTo()

			moveTo.PosX = newPos.X
			moveTo.PosY = newPos.Y

			zone.NotifyPlayers(types.Pointer(behaviour.OwnerID()), func() *byter.Byter {
				CEWriter := objects.NewClientEntityWriterWithByter()
				CEWriter.BeginComponentUpdate(behaviour)
				CEWriter.CreateActionComplete(moveTo)
				CEWriter.EndComponentUpdate(behaviour)

				return CEWriter.Body
			})

			playerPosMap[player.Conn.GetID()] = newPos
		}
	}
}
