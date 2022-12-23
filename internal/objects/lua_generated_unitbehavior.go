// Code generated by scripts/generatelua DO NOT EDIT.
package objects

import (
	actions2 "RainbowRunner/internal/actions"
	lua "RainbowRunner/internal/lua"
	"RainbowRunner/pkg/byter"
	"RainbowRunner/pkg/datatypes"
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
	return lua.LuaMethodsExtend(map[string]lua2.LGFunction{
		"speed":          lua.LuaGenericGetSetNumber[IUnitBehavior](func(v IUnitBehavior) *int { return &v.GetUnitBehavior().Speed }),
		"turnRate":       lua.LuaGenericGetSetNumber[IUnitBehavior](func(v IUnitBehavior) *int { return &v.GetUnitBehavior().TurnRate }),
		"lastPosition":   lua.LuaGenericGetSetValue[IUnitBehavior, datatypes.Vector3Float32](func(v IUnitBehavior) *datatypes.Vector3Float32 { return &v.GetUnitBehavior().LastPosition }),
		"position":       lua.LuaGenericGetSetValue[IUnitBehavior, datatypes.Vector3Float32](func(v IUnitBehavior) *datatypes.Vector3Float32 { return &v.GetUnitBehavior().Position }),
		"rotation":       lua.LuaGenericGetSetNumber[IUnitBehavior](func(v IUnitBehavior) *float32 { return &v.GetUnitBehavior().Rotation }),
		"unitMoverFlags": lua.LuaGenericGetSetNumber[IUnitBehavior](func(v IUnitBehavior) *byte { return &v.GetUnitBehavior().UnitMoverFlags }),
		// -------------------------------------------------------------------------------------------------------------
		// Unsupported field type Action1
		// -------------------------------------------------------------------------------------------------------------
		// -------------------------------------------------------------------------------------------------------------
		// Unsupported field type Action2
		// -------------------------------------------------------------------------------------------------------------
		"sessionID":        lua.LuaGenericGetSetNumber[IUnitBehavior](func(v IUnitBehavior) *byte { return &v.GetUnitBehavior().SessionID }),
		"unitMoverUnk0":    lua.LuaGenericGetSetNumber[IUnitBehavior](func(v IUnitBehavior) *byte { return &v.GetUnitBehavior().UnitMoverUnk0 }),
		"unitMoverUnk1":    lua.LuaGenericGetSetNumber[IUnitBehavior](func(v IUnitBehavior) *uint32 { return &v.GetUnitBehavior().UnitMoverUnk1 }),
		"unitMoverUnk2":    lua.LuaGenericGetSetNumber[IUnitBehavior](func(v IUnitBehavior) *uint32 { return &v.GetUnitBehavior().UnitMoverUnk2 }),
		"unitMoverUnk3":    lua.LuaGenericGetSetNumber[IUnitBehavior](func(v IUnitBehavior) *uint32 { return &v.GetUnitBehavior().UnitMoverUnk3 }),
		"unitMoverUnk4":    lua.LuaGenericGetSetNumber[IUnitBehavior](func(v IUnitBehavior) *uint32 { return &v.GetUnitBehavior().UnitMoverUnk4 }),
		"unitMoverUnk7":    lua.LuaGenericGetSetNumber[IUnitBehavior](func(v IUnitBehavior) *uint32 { return &v.GetUnitBehavior().UnitMoverUnk7 }),
		"unitBehaviorUnk0": lua.LuaGenericGetSetNumber[IUnitBehavior](func(v IUnitBehavior) *byte { return &v.GetUnitBehavior().UnitBehaviorUnk0 }),
		"unitBehaviorUnk1": lua.LuaGenericGetSetNumber[IUnitBehavior](func(v IUnitBehavior) *byte { return &v.GetUnitBehavior().UnitBehaviorUnk1 }),
		"unitBehaviorUnk2": lua.LuaGenericGetSetNumber[IUnitBehavior](func(v IUnitBehavior) *byte { return &v.GetUnitBehavior().UnitBehaviorUnk2 }),
		"isMoving":         lua.LuaGenericGetSetBool[IUnitBehavior](func(v IUnitBehavior) *bool { return &v.GetUnitBehavior().IsMoving }),
		"tick": func(l *lua2.LState) int {
			objInterface := lua.CheckInterfaceValue[IUnitBehavior](l, 1)
			obj := objInterface.GetUnitBehavior()
			obj.Tick()

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
		"writeMoveUpdate": func(l *lua2.LState) int {
			objInterface := lua.CheckInterfaceValue[IUnitBehavior](l, 1)
			obj := objInterface.GetUnitBehavior()
			obj.WriteMoveUpdate(
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
		"warpTo": func(l *lua2.LState) int {
			objInterface := lua.CheckInterfaceValue[IUnitBehavior](l, 1)
			obj := objInterface.GetUnitBehavior()
			obj.WarpTo(
				lua.CheckValue[datatypes.Vector3Float32](l, 2),
			)

			return 0
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
		"moveTo": func(l *lua2.LState) int {
			objInterface := lua.CheckInterfaceValue[IUnitBehavior](l, 1)
			obj := objInterface.GetUnitBehavior()
			obj.MoveTo(
				lua.CheckValue[datatypes.Vector2Float32](l, 2),
			)

			return 0
		},
		"moveToEntity": func(l *lua2.LState) int {
			objInterface := lua.CheckInterfaceValue[IUnitBehavior](l, 1)
			obj := objInterface.GetUnitBehavior()
			obj.MoveToEntity(
				lua.CheckValue[IWorldEntity](l, 2),
			)

			return 0
		},
		"executeAction": func(l *lua2.LState) int {
			objInterface := lua.CheckInterfaceValue[IUnitBehavior](l, 1)
			obj := objInterface.GetUnitBehavior()
			obj.ExecuteAction(
				lua.CheckValue[actions2.Action](l, 2),
			)

			return 0
		},
		"stopFollowClient": func(l *lua2.LState) int {
			objInterface := lua.CheckInterfaceValue[IUnitBehavior](l, 1)
			obj := objInterface.GetUnitBehavior()
			obj.StopFollowClient()

			return 0
		},
		"getUnitBehavior": func(l *lua2.LState) int {
			objInterface := lua.CheckInterfaceValue[IUnitBehavior](l, 1)
			obj := objInterface.GetUnitBehavior()
			res0 := obj.GetUnitBehavior()
			if res0 != nil {
				l.Push(res0.ToLua(l))
			} else {
				l.Push(lua2.LNil)
			}

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
