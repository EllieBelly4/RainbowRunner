package script

import (
	"RainbowRunner/internal/global"
	lua2 "github.com/yuin/gopher-lua"
	"time"
)

func RegisterAllTime(s *lua2.LState) {
	mt := s.NewTypeMetatable("Time")
	s.SetGlobal("Time", mt)
	s.SetFuncs(mt, map[string]lua2.LGFunction{
		"time":      luaGetTimeSinceStart,
		"unix":      luaGetUnixTime,
		"tick":      luaGetTick,
		"deltaTime": luaGetDeltaTime,
	})
}

func luaGetUnixTime(state *lua2.LState) int {
	state.Push(lua2.LNumber(time.Now().Unix()))
	return 1
}

func luaGetTick(state *lua2.LState) int {
	state.Push(lua2.LNumber(global.GetTick()))
	return 1
}

func luaGetDeltaTime(state *lua2.LState) int {
	state.Push(lua2.LNumber(global.GetDeltaTime()))
	return 1
}

func luaGetTimeSinceStart(state *lua2.LState) int {
	state.Push(lua2.LNumber(global.GetTimeSinceStartSeconds()))
	return 1
}
