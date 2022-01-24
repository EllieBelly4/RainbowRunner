package game

import (
	"RainbowRunner/internal/global"
	"RainbowRunner/internal/objects"
	"RainbowRunner/pkg"
	"RainbowRunner/pkg/math"
	"time"
)

var i = 1

func StartGameLoop() {
	ticker := time.NewTicker(33 * time.Millisecond)
	defer ticker.Stop()

	for {
		select {
		case <-ticker.C:
			objects.Players.RLock()
			objects.Players.BeforeTick()

			for !global.JobQueue.Empty() {
				job := global.JobQueue.Dequeue()
				job()
			}

			objects.Entities.Tick()

			if global.Tick%150 == 0 {
				for _, player := range objects.Players.Players {
					objects.CreateNPC(player, player.Zone, pkg.Transform{
						Position: pkg.Vector3{106342 + 2048*int32(i), -36000, 12778},
						Rotation: 180 * math.DRDegToRot,
					}, "npc.Avatar.Female.base.NPC_Amazon1_Base", "npc.Avatar.Female.base.NPC_Amazon1_Base.Behavior")
				}

				i++
			}

			for _, player := range objects.Players.Players {
				player.Conn.Client.Tick()
			}

			objects.Players.AfterTick()
			objects.Players.RUnlock()

			//if conn.Player.IsMoving {
			//	conn.Player.SendPosition()
			//}else
			//if ticks % 100 == 0 {
			//	conn.Player.SendFollowClient()
			//}

			//if conn.Client.TicksSinceLastUpdate >= 0x2D {
			//	conn.Client.SendPosition()
			//}

			//mov := conn.Player.LastMovementRequest
			//SendMoveTo(conn, 0x05, mov.X, mov.Y)

			global.Tick++
		}
	}
}
