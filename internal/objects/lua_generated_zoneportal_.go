package objects

/**
 * This file is generated by scripts/generatelua/generatelua.go
 * DO NOT EDIT
 */

import (
	lua "RainbowRunner/internal/lua"
	"RainbowRunner/pkg/byter"
	lua2 "github.com/yuin/gopher-lua"
)

func registerLuaZonePortal(state *lua2.LState) {
	// Ensure the import is referenced in code
	_ = lua.LuaScript{}

	mt := state.NewTypeMetatable("ZonePortal")
	state.SetGlobal("ZonePortal", mt)
	state.SetField(mt, "new", state.NewFunction(newLuaZonePortal))
	state.SetField(mt, "__index", state.SetFuncs(state.NewTable(),
		luaMethodsZonePortal(),
	))
}

func luaMethodsZonePortal() map[string]lua2.LGFunction {
	return luaMethodsExtend(map[string]lua2.LGFunction{
		"unk0":   luaGenericGetSetString[ZonePortal](func(v ZonePortal) *string { return &v.Unk0 }),
		"unk1":   luaGenericGetSetString[ZonePortal](func(v ZonePortal) *string { return &v.Unk1 }),
		"width":  luaGenericGetSetNumber[ZonePortal, uint16](func(v ZonePortal) *uint16 { return &v.Width }),
		"height": luaGenericGetSetNumber[ZonePortal, uint16](func(v ZonePortal) *uint16 { return &v.Height }),
		"unk4":   luaGenericGetSetNumber[ZonePortal, uint32](func(v ZonePortal) *uint32 { return &v.Unk4 }),
		"target": luaGenericGetSetString[ZonePortal](func(v ZonePortal) *string { return &v.Target }),
		"activate": func(l *lua2.LState) int {
			obj := lua.CheckReferenceValue[ZonePortal](l, 1)
			obj.Activate(
				lua.CheckReferenceValue[RRPlayer](l, 2),
				lua.CheckReferenceValue[UnitBehavior](l, 3),
				lua.CheckValue[byte](l, 4),
			)

			return 0
		},
		"writeInit": func(l *lua2.LState) int {
			obj := lua.CheckReferenceValue[ZonePortal](l, 1)
			obj.WriteInit(
				lua.CheckReferenceValue[byter.Byter](l, 2),
			)

			return 0
		},
		"toLua": func(l *lua2.LState) int {
			obj := lua.CheckReferenceValue[ZonePortal](l, 1)
			res0 := obj.ToLua(
				lua.CheckReferenceValue[lua2.LState](l, 2),
			)
			ud := l.NewUserData()
			ud.Value = res0
			l.SetMetatable(ud, l.GetTypeMetatable("lua2.LValue"))
			l.Push(ud)

			return 1
		},
		"WorldEntity": func(l *lua2.LState) int {
			obj := lua.CheckReferenceValue[ZonePortal](l, 1)
			l.Push(obj.WorldEntity.ToLua(l))
			return 1
		},
	})
}
func newLuaZonePortal(l *lua2.LState) int {
	obj := NewZonePortal(string(l.CheckString(1)), string(l.CheckString(2)))
	ud := l.NewUserData()
	ud.Value = obj

	l.SetMetatable(ud, l.GetTypeMetatable("ZonePortal"))
	l.Push(ud)
	return 1
}

func (z *ZonePortal) ToLua(l *lua2.LState) lua2.LValue {
	ud := l.NewUserData()
	ud.Value = z

	l.SetMetatable(ud, l.GetTypeMetatable("ZonePortal"))
	return ud
}
