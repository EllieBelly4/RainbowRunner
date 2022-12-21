package game

import (
	"RainbowRunner/internal/global"
	"RainbowRunner/internal/objects"
	"time"
)

var i = 1

func StartGameLoop() {
	ticker := time.NewTicker(global.TickInterval * time.Millisecond)
	defer ticker.Stop()

	for {
		select {
		case <-ticker.C:
			for !global.JobQueue.IsEmpty() {
				job := global.JobQueue.Dequeue()
				job()
			}

			objects.Players.RLock()
			objects.Players.BeforeTick()

			for _, player := range objects.Players.Players {
				player.Conn.Client.Tick()
			}

			objects.Players.AfterTick()
			objects.Players.RUnlock()

			objects.Zones.Tick()

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
