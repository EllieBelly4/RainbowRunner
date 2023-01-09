package objects

import (
	"RainbowRunner/internal/actions"
	"RainbowRunner/internal/database"
	lua2 "RainbowRunner/internal/lua"
	"RainbowRunner/internal/script"
	"RainbowRunner/internal/types/configtypes"
	"RainbowRunner/pkg/datatypes"
	"RainbowRunner/pkg/datatypes/drfloat"
	lua "github.com/yuin/gopher-lua"
)

//go:generate go run ../../scripts/generateluaregistrations -includes=.,../actions,../database,../types/configtypes
func RegisterLuaGlobals(state *lua.LState) {
	script.RegisterAllTime(state)
	RegisterAllLuaFunctions(state)
	actions.RegisterAllLuaFunctions(state)
	database.RegisterAllLuaFunctions(state)
	configtypes.RegisterAllLuaFunctions(state)

	lua2.RegisterModules(state)
}

func registerLuaDRFloat(s *lua.LState) {
	mt := s.NewTypeMetatable("DRFloat")
	s.SetGlobal("DRFloat", mt)
	s.SetField(mt, "__index", s.SetFuncs(s.NewTable(),
		map[string]lua.LGFunction{},
	))
	s.SetField(mt, "new", s.NewFunction(func(state *lua.LState) int {
		num := state.CheckNumber(1)

		ud := state.NewUserData()
		ud.Value = drfloat.FromFloat32(float32(num))
		state.SetMetatable(ud, state.GetTypeMetatable("DRFloat"))

		state.Push(ud)
		return 1
	}))
}

func registerLuaVector3DRFloat(s *lua.LState) {
	mt := s.NewTypeMetatable("Vector3DRFloat")
	s.SetGlobal("Vector3DRFloat", mt)
	s.SetField(mt, "__index", s.SetFuncs(s.NewTable(),
		map[string]lua.LGFunction{
			"string": func(l *lua.LState) int {
				obj := lua2.CheckValue[datatypes.Vector3DRFloat](l, 1)
				l.Push(lua.LString(obj.String()))
				return 1
			},
			"x": lua2.LuaGenericGetSetValueAny[datatypes.Vector3DRFloat](func(v datatypes.Vector3DRFloat) *drfloat.DRFloat { return &v.X }),
			"y": lua2.LuaGenericGetSetValueAny[datatypes.Vector3DRFloat](func(v datatypes.Vector3DRFloat) *drfloat.DRFloat { return &v.Y }),
			"z": lua2.LuaGenericGetSetValueAny[datatypes.Vector3DRFloat](func(v datatypes.Vector3DRFloat) *drfloat.DRFloat { return &v.Z }),
		},
	))
	s.SetField(mt, "new", s.NewFunction(newLuaVector3))
}

func registerLuaVector3(s *lua.LState) {
	mt := s.NewTypeMetatable("Vector3")
	s.SetGlobal("Vector3", mt)
	s.SetField(mt, "__index", s.SetFuncs(s.NewTable(),
		map[string]lua.LGFunction{
			"string": func(l *lua.LState) int {
				obj := lua2.CheckValue[datatypes.Vector3Float32](l, 1)
				l.Push(lua.LString(obj.String()))
				return 1
			},
			"x": lua2.LuaGenericGetSetNumber[datatypes.Vector3Float32](func(v datatypes.Vector3Float32) *float32 { return &v.X }),
			"y": lua2.LuaGenericGetSetNumber[datatypes.Vector3Float32](func(v datatypes.Vector3Float32) *float32 { return &v.Y }),
			"z": lua2.LuaGenericGetSetNumber[datatypes.Vector3Float32](func(v datatypes.Vector3Float32) *float32 { return &v.Z }),
		},
	))
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
	state.SetMetatable(ud, state.GetTypeMetatable("Vector3"))

	state.Push(ud)
	return 1
}

func registerLuaVector2(s *lua.LState) {
	mt := s.NewTypeMetatable("Vector2")
	s.SetGlobal("Vector2", mt)
	s.SetField(mt, "__index", s.SetFuncs(s.NewTable(),
		map[string]lua.LGFunction{
			"string": func(l *lua.LState) int {
				obj := lua2.CheckValue[datatypes.Vector2Float32](l, 1)
				l.Push(lua.LString(obj.String()))
				return 1
			},
			"x": lua2.LuaGenericGetSetNumber[datatypes.Vector2Float32](func(v datatypes.Vector2Float32) *float32 { return &v.X }),
			"y": lua2.LuaGenericGetSetNumber[datatypes.Vector2Float32](func(v datatypes.Vector2Float32) *float32 { return &v.Y }),
		},
	))
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
	state.SetMetatable(ud, state.GetTypeMetatable("Vector2"))

	state.Push(ud)
	return 1
}
