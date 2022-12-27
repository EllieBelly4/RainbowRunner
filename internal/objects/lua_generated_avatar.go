// Code generated by scripts/generatelua DO NOT EDIT.
package objects

import (
	lua "RainbowRunner/internal/lua"
	"RainbowRunner/internal/types/drobjecttypes"
	"RainbowRunner/pkg/byter"
	"RainbowRunner/pkg/datatypes"
	lua2 "github.com/yuin/gopher-lua"
)

type IAvatar interface {
	GetAvatar() *Avatar
}

func (a *Avatar) GetAvatar() *Avatar {
	return a
}

func registerLuaAvatar(state *lua2.LState) {
	// Ensure the import is referenced in code
	_ = lua.LuaScript{}

	mt := state.NewTypeMetatable("Avatar")
	state.SetGlobal("Avatar", mt)
	state.SetField(mt, "new", state.NewFunction(newLuaAvatar))
	state.SetField(mt, "__index", state.SetFuncs(state.NewTable(),
		luaMethodsAvatar(),
	))
}

func luaMethodsAvatar() map[string]lua2.LGFunction {
	return lua.LuaMethodsExtend(map[string]lua2.LGFunction{
		"isMoving":           lua.LuaGenericGetSetBool[IAvatar](func(v IAvatar) *bool { return &v.GetAvatar().IsMoving }),
		"rotation":           lua.LuaGenericGetSetNumber[IAvatar](func(v IAvatar) *int32 { return &v.GetAvatar().Rotation }),
		"clientUpdateNumber": lua.LuaGenericGetSetNumber[IAvatar](func(v IAvatar) *byte { return &v.GetAvatar().ClientUpdateNumber }),
		"moveUpdate":         lua.LuaGenericGetSetNumber[IAvatar](func(v IAvatar) *int { return &v.GetAvatar().MoveUpdate }),
		"isSpawned":          lua.LuaGenericGetSetBool[IAvatar](func(v IAvatar) *bool { return &v.GetAvatar().IsSpawned }),
		"addChild": func(l *lua2.LState) int {
			objInterface := lua.CheckInterfaceValue[IAvatar](l, 1)
			obj := objInterface.GetAvatar()
			obj.AddChild(
				lua.CheckValue[drobjecttypes.DRObject](l, 2),
			)

			return 0
		},
		"type": func(l *lua2.LState) int {
			objInterface := lua.CheckInterfaceValue[IAvatar](l, 1)
			obj := objInterface.GetAvatar()
			res0 := obj.Type()
			ud := l.NewUserData()
			ud.Value = res0
			l.SetMetatable(ud, l.GetTypeMetatable("drobjecttypes.DRObjectType"))
			l.Push(ud)

			return 1
		},
		"writeFullGCObject": func(l *lua2.LState) int {
			objInterface := lua.CheckInterfaceValue[IAvatar](l, 1)
			obj := objInterface.GetAvatar()
			obj.WriteFullGCObject(
				lua.CheckReferenceValue[byter.Byter](l, 2),
			)

			return 0
		},
		"writeInit": func(l *lua2.LState) int {
			objInterface := lua.CheckInterfaceValue[IAvatar](l, 1)
			obj := objInterface.GetAvatar()
			obj.WriteInit(
				lua.CheckReferenceValue[byter.Byter](l, 2),
			)

			return 0
		},
		"writeUpdate": func(l *lua2.LState) int {
			objInterface := lua.CheckInterfaceValue[IAvatar](l, 1)
			obj := objInterface.GetAvatar()
			obj.WriteUpdate(
				lua.CheckReferenceValue[byter.Byter](l, 2),
			)

			return 0
		},
		"tick": func(l *lua2.LState) int {
			objInterface := lua.CheckInterfaceValue[IAvatar](l, 1)
			obj := objInterface.GetAvatar()
			obj.Tick()

			return 0
		},
		"getUnitBehaviourID": func(l *lua2.LState) int {
			objInterface := lua.CheckInterfaceValue[IAvatar](l, 1)
			obj := objInterface.GetAvatar()
			res0 := obj.GetUnitBehaviourID()
			l.Push(lua2.LNumber(res0))

			return 1
		},
		"send": func(l *lua2.LState) int {
			objInterface := lua.CheckInterfaceValue[IAvatar](l, 1)
			obj := objInterface.GetAvatar()
			obj.Send(
				lua.CheckReferenceValue[byter.Byter](l, 2),
			)

			return 0
		},
		"sendFollowClient": func(l *lua2.LState) int {
			objInterface := lua.CheckInterfaceValue[IAvatar](l, 1)
			obj := objInterface.GetAvatar()
			obj.SendFollowClient()

			return 0
		},
		"sendStopFollowClient": func(l *lua2.LState) int {
			objInterface := lua.CheckInterfaceValue[IAvatar](l, 1)
			obj := objInterface.GetAvatar()
			obj.SendStopFollowClient()

			return 0
		},
		"warp": func(l *lua2.LState) int {
			objInterface := lua.CheckInterfaceValue[IAvatar](l, 1)
			obj := objInterface.GetAvatar()
			obj.Warp(float32(l.CheckNumber(2)), float32(l.CheckNumber(3)), float32(l.CheckNumber(4)))

			return 0
		},
		"sendMoveTo": func(l *lua2.LState) int {
			objInterface := lua.CheckInterfaceValue[IAvatar](l, 1)
			obj := objInterface.GetAvatar()
			obj.SendMoveTo(uint8(l.CheckNumber(2)), uint16(l.CheckNumber(3)), float32(l.CheckNumber(4)), float32(l.CheckNumber(5)))

			return 0
		},
		"getUnitContainer": func(l *lua2.LState) int {
			objInterface := lua.CheckInterfaceValue[IAvatar](l, 1)
			obj := objInterface.GetAvatar()
			res0 := obj.GetUnitContainer()
			if res0 != nil {
				l.Push(res0.ToLua(l))
			} else {
				l.Push(lua2.LNil)
			}

			return 1
		},
		"getManipulators": func(l *lua2.LState) int {
			objInterface := lua.CheckInterfaceValue[IAvatar](l, 1)
			obj := objInterface.GetAvatar()
			res0 := obj.GetManipulators()
			if res0 != nil {
				l.Push(res0.ToLua(l))
			} else {
				l.Push(lua2.LNil)
			}

			return 1
		},
		"getUnitBehaviour": func(l *lua2.LState) int {
			objInterface := lua.CheckInterfaceValue[IAvatar](l, 1)
			obj := objInterface.GetAvatar()
			res0 := obj.GetUnitBehaviour()
			if res0 != nil {
				l.Push(res0.ToLua(l))
			} else {
				l.Push(lua2.LNil)
			}

			return 1
		},
		"teleport": func(l *lua2.LState) int {
			objInterface := lua.CheckInterfaceValue[IAvatar](l, 1)
			obj := objInterface.GetAvatar()
			obj.Teleport(
				lua.CheckValue[datatypes.Vector3Float32](l, 2),
			)

			return 0
		},
		"getAvatar": func(l *lua2.LState) int {
			objInterface := lua.CheckInterfaceValue[IAvatar](l, 1)
			obj := objInterface.GetAvatar()
			res0 := obj.GetAvatar()
			if res0 != nil {
				l.Push(res0.ToLua(l))
			} else {
				l.Push(lua2.LNil)
			}

			return 1
		},
	}, luaMethodsHero)
}
func newLuaAvatar(l *lua2.LState) int {
	obj := NewAvatar(string(l.CheckString(1)))
	ud := l.NewUserData()
	ud.Value = obj

	l.SetMetatable(ud, l.GetTypeMetatable("Avatar"))
	l.Push(ud)
	return 1
}

func (a *Avatar) ToLua(l *lua2.LState) lua2.LValue {
	ud := l.NewUserData()
	ud.Value = a

	l.SetMetatable(ud, l.GetTypeMetatable("Avatar"))
	return ud
}
