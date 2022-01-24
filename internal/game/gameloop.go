package game

import (
	"RainbowRunner/internal/objects"
	"RainbowRunner/internal/state"
	"time"
)

func StartGameLoop() {
	ticker := time.NewTicker(33 * time.Millisecond)
	defer ticker.Stop()

	for {
		select {
		case <-ticker.C:
			objects.Players.RLock()
			objects.Players.BeforeTick()
			objects.Entities.Tick()

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

			state.Tick++
		}
	}
}
