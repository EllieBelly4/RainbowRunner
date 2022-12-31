// Code generated by scripts/generatelua DO NOT EDIT.
package actions

import (
	lua "RainbowRunner/internal/lua"
	byter "RainbowRunner/pkg/byter"
	lua2 "github.com/yuin/gopher-lua"
)

type IActionAttackTarget2 interface {
	GetActionAttackTarget2() *ActionAttackTarget2
}

func (a *ActionAttackTarget2) GetActionAttackTarget2() *ActionAttackTarget2 {
	return a
}

func registerLuaActionAttackTarget2(state *lua2.LState) {
	// Ensure the import is referenced in code
	_ = lua.LuaScript{}

	mt := state.NewTypeMetatable("ActionAttackTarget2")
	state.SetGlobal("ActionAttackTarget2", mt)
	state.SetField(mt, "new", state.NewFunction(newLuaActionAttackTarget2))
	state.SetField(mt, "__index", state.SetFuncs(state.NewTable(),
		luaMethodsActionAttackTarget2(),
	))
}

func luaMethodsActionAttackTarget2() map[string]lua2.LGFunction {
	return lua.LuaMethodsExtend(map[string]lua2.LGFunction{

		"unk0": lua.LuaGenericGetSetValueAny[IActionAttackTarget2](func(v IActionAttackTarget2) *byte { return &v.GetActionAttackTarget2().Unk0 }),

		"targetID": lua.LuaGenericGetSetValueAny[IActionAttackTarget2](func(v IActionAttackTarget2) *uint16 { return &v.GetActionAttackTarget2().TargetID }),

		"opCode": func(l *lua2.LState) int {
			objInterface := lua.CheckInterfaceValue[IActionAttackTarget2](l, 1)
			obj := objInterface.GetActionAttackTarget2()
			res0 := obj.OpCode()
			ud := l.NewUserData()
			ud.Value = res0
			l.SetMetatable(ud, l.GetTypeMetatable("BehaviourAction"))
			l.Push(ud)

			return 1
		},

		"init": func(l *lua2.LState) int {
			objInterface := lua.CheckInterfaceValue[IActionAttackTarget2](l, 1)
			obj := objInterface.GetActionAttackTarget2()
			obj.Init(
				lua.CheckReferenceValue[byter.Byter](l, 2),
			)

			return 0
		},

		"getActionAttackTarget2": func(l *lua2.LState) int {
			objInterface := lua.CheckInterfaceValue[IActionAttackTarget2](l, 1)
			obj := objInterface.GetActionAttackTarget2()
			res0 := obj.GetActionAttackTarget2()
			if res0 != nil {
				l.Push(res0.ToLua(l))
			} else {
				l.Push(lua2.LNil)
			}

			return 1
		},
	})
}
func newLuaActionAttackTarget2(l *lua2.LState) int {
	obj := NewActionAttackTarget2()
	ud := l.NewUserData()
	ud.Value = obj

	l.SetMetatable(ud, l.GetTypeMetatable("ActionAttackTarget2"))
	l.Push(ud)
	return 1
}

func (a *ActionAttackTarget2) ToLua(l *lua2.LState) lua2.LValue {
	ud := l.NewUserData()
	ud.Value = a

	l.SetMetatable(ud, l.GetTypeMetatable("ActionAttackTarget2"))
	return ud
}
