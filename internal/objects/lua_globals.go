package objects

import (
	"RainbowRunner/pkg/datatypes"
	lua "github.com/yuin/gopher-lua"
)

func AddGlobals(state *lua.LState) {
	addVector3(state)
	addVector2(state)
}

func addVector3(s *lua.LState) {
	mt := s.NewTypeMetatable("vector3")
	s.SetGlobal("vector3", mt)
	s.SetField(mt, "new", s.NewFunction(newLuaVector3))
}

func newLuaVector3(state *lua.LState) int {
	v3 := datatypes.Vector3Float32{}

	if state.GetTop() == 3 {
		v3.X = float32(state.CheckNumber(1))
		v3.Y = float32(state.CheckNumber(2))
		v3.Z = float32(state.CheckNumber(3))
	}

	ud := state.NewUserData()
	ud.Value = v3
	state.SetMetatable(ud, state.GetTypeMetatable("vector3"))

	state.Push(ud)
	return 1
}

func addVector2(s *lua.LState) {
	mt := s.NewTypeMetatable("vector2")
	s.SetGlobal("vector2", mt)
	s.SetField(mt, "new", s.NewFunction(newLuaVector2))
}

func newLuaVector2(state *lua.LState) int {
	v2 := datatypes.Vector2Float32{}

	if state.GetTop() == 2 {
		v2.X = float32(state.CheckNumber(1))
		v2.Y = float32(state.CheckNumber(2))
	}

	ud := state.NewUserData()
	ud.Value = v2
	state.SetMetatable(ud, state.GetTypeMetatable("vector2"))

	state.Push(ud)
	return 1
}
