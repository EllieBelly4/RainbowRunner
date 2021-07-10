package game

import (
	"RainbowRunner/internal/managers"
	"time"
)

var Tick = uint(0)

func StartGameLoop(conn *RRConn) {
	ticker := time.NewTicker(33 * time.Millisecond)
	defer ticker.Stop()

	for conn.IsConnected {
		select {
		case <-ticker.C:
			managers.Entities.Tick()

			conn.Client.Tick()

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

			Tick++
		}
	}
}
