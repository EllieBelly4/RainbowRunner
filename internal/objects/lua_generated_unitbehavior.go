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

type IUnitBehavior interface {
	GetUnitBehavior() *UnitBehavior
}

func (u *UnitBehavior) GetUnitBehavior() *UnitBehavior {
	return u
}

func registerLuaUnitBehavior(state *lua2.LState) {
	// Ensure the import is referenced in code
	_ = lua.LuaScript{}

	mt := state.NewTypeMetatable("UnitBehavior")
	state.SetGlobal("UnitBehavior", mt)
	state.SetField(mt, "new", state.NewFunction(newLuaUnitBehavior))
	state.SetField(mt, "__index", state.SetFuncs(state.NewTable(),
		luaMethodsUnitBehavior(),
	))
}

func luaMethodsUnitBehavior() map[string]lua2.LGFunction {
	return luaMethodsExtend(map[string]lua2.LGFunction{
		"rotation":       luaGenericGetSetNumber[IUnitBehavior](func(v IUnitBehavior) *float32 { return &v.GetUnitBehavior().Rotation }),
		"unitMoverFlags": luaGenericGetSetNumber[IUnitBehavior](func(v IUnitBehavior) *byte { return &v.GetUnitBehavior().UnitMoverFlags }),
		"writeMoveUpdate": func(l *lua2.LState) int {
			objInterface := lua.CheckInterfaceValue[IUnitBehavior](l, 1)
			obj := objInterface.GetUnitBehavior()
			obj.WriteMoveUpdate(
				lua.CheckReferenceValue[byter.Byter](l, 2),
			)

			return 0
		},
		"writeInit": func(l *lua2.LState) int {
			objInterface := lua.CheckInterfaceValue[IUnitBehavior](l, 1)
			obj := objInterface.GetUnitBehavior()
			obj.WriteInit(
				lua.CheckReferenceValue[byter.Byter](l, 2),
			)

			return 0
		},
		"readUpdate": func(l *lua2.LState) int {
			objInterface := lua.CheckInterfaceValue[IUnitBehavior](l, 1)
			obj := objInterface.GetUnitBehavior()
			res0 := obj.ReadUpdate(
				lua.CheckReferenceValue[byter.Byter](l, 2),
			)
			ud := l.NewUserData()
			ud.Value = res0
			l.SetMetatable(ud, l.GetTypeMetatable("error"))
			l.Push(ud)

			return 1
		},
		"warp": func(l *lua2.LState) int {
			objInterface := lua.CheckInterfaceValue[IUnitBehavior](l, 1)
			obj := objInterface.GetUnitBehavior()
			obj.Warp(float32(l.CheckNumber(2)), float32(l.CheckNumber(3)), float32(l.CheckNumber(4)))

			return 0
		},
		"writeWarp": func(l *lua2.LState) int {
			objInterface := lua.CheckInterfaceValue[IUnitBehavior](l, 1)
			obj := objInterface.GetUnitBehavior()
			obj.WriteWarp(
				lua.CheckReferenceValue[ClientEntityWriter](l, 2),
			)

			return 0
		},
		"getUnitBehavior": func(l *lua2.LState) int {
			objInterface := lua.CheckInterfaceValue[IUnitBehavior](l, 1)
			obj := objInterface.GetUnitBehavior()
			res0 := obj.GetUnitBehavior()
			ud := l.NewUserData()
			ud.Value = res0
			l.SetMetatable(ud, l.GetTypeMetatable("UnitBehavior"))
			l.Push(ud)

			return 1
		},
	}, luaMethodsComponent)
}
func newLuaUnitBehavior(l *lua2.LState) int {
	obj := NewUnitBehavior(string(l.CheckString(1)))
	ud := l.NewUserData()
	ud.Value = obj

	l.SetMetatable(ud, l.GetTypeMetatable("UnitBehavior"))
	l.Push(ud)
	return 1
}

func (u *UnitBehavior) ToLua(l *lua2.LState) lua2.LValue {
	ud := l.NewUserData()
	ud.Value = u

	l.SetMetatable(ud, l.GetTypeMetatable("UnitBehavior"))
	return ud
}