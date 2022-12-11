package objects

/**
 * This file is generated by scripts/generatelua/generatelua.go
 * DO NOT EDIT
 */

import (
	lua "RainbowRunner/internal/lua"
	"RainbowRunner/pkg/byter"
	"RainbowRunner/pkg/datatypes"
	lua2 "github.com/yuin/gopher-lua"
)

func registerLuaNPC(state *lua2.LState) {
	// Ensure the import is referenced in code
	_ = lua.LuaScript{}

	mt := state.NewTypeMetatable("NPC")
	state.SetGlobal("NPC", mt)
	state.SetField(mt, "new", state.NewFunction(newLuaNPC))
	state.SetField(mt, "__index", state.SetFuncs(state.NewTable(),
		luaMethodsNPC(),
	))
}

func luaMethodsNPC() map[string]lua2.LGFunction {
	return luaMethodsExtend(map[string]lua2.LGFunction{
		"name":  luaGenericGetSetString[NPC](func(v NPC) *string { return &v.Name }),
		"level": luaGenericGetSetNumber[NPC, int32](func(v NPC) *int32 { return &v.Level }),
		"writeInit": func(l *lua2.LState) int {
			obj := lua.CheckReferenceValue[NPC](l, 1)
			obj.WriteInit(
				lua.CheckReferenceValue[byter.Byter](l, 1),
			)

			return 0
		},
	}, luaMethodsUnit)
}
func newLuaNPC(l *lua2.LState) int {
	obj := NewNPC(string(l.CheckString(1)), string(l.CheckString(2)),
		lua.CheckValue[datatypes.Vector3Float32](l, 3), float32(l.CheckNumber(4)),
	)
	ud := l.NewUserData()
	ud.Value = obj

	l.SetMetatable(ud, l.GetTypeMetatable("NPC"))
	l.Push(ud)
	return 1
}
