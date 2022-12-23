package global

import "time"

// 33 is ideal
const TickInterval = 33

var Tick = uint(0)
var ServerStartTime = time.Time{}

func GetTick() uint {
	return Tick
}

func GetDeltaTime() float64 {
	return TickInterval / 1000.0
}

func GetTimeSinceStartSeconds() float64 {
	return time.Since(ServerStartTime).Seconds()
}
