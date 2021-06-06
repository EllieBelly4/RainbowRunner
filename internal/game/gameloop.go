package game

import (
	"time"
)

func StartGameLoop(conn *RRConn) {
	ticker := time.NewTicker(33 * time.Millisecond)
	defer ticker.Stop()

	ticks := 0

	for conn.IsConnected {
		select {
		case <-ticker.C:
			if !conn.Player.IsSpawned {
				continue
			}

			conn.Player.Tick()

			//if conn.Player.IsMoving {
			//	conn.Player.SendPosition()
			//}else
			//if ticks % 100 == 0 {
			//	conn.Player.SendFollowClient()
			//}

			if conn.Player.TicksSinceLastUpdate >= 0x2D {
				conn.Player.SendPosition()
			}

			//mov := conn.Player.LastMovementRequest
			//SendMoveTo(conn, 0x05, mov.X, mov.Y)

			ticks++
		}
	}
}
