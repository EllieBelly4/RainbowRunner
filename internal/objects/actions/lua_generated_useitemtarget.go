// Code generated by scripts/generatelua DO NOT EDIT.
package actions

import (
	lua "RainbowRunner/internal/lua"
	byter "RainbowRunner/pkg/byter"
	lua2 "github.com/yuin/gopher-lua"
)

type IUseItemTarget interface {
	GetUseItemTarget() *UseItemTarget
}

func (u *UseItemTarget) GetUseItemTarget() *UseItemTarget {
	return u
}

func registerLuaUseItemTarget(state *lua2.LState) {
	// Ensure the import is referenced in code
	_ = lua.LuaScript{}

	mt := state.NewTypeMetatable("UseItemTarget")
	state.SetGlobal("UseItemTarget", mt)
	state.SetField(mt, "__index", state.SetFuncs(state.NewTable(),
		luaMethodsUseItemTarget(),
	))
}

func luaMethodsUseItemTarget() map[string]lua2.LGFunction {
	return lua.LuaMethodsExtend(map[string]lua2.LGFunction{
		"opCode": func(l *lua2.LState) int {
			objInterface := lua.CheckInterfaceValue[IUseItemTarget](l, 1)
			obj := objInterface.GetUseItemTarget()
			res0 := obj.OpCode()
			ud := l.NewUserData()
			ud.Value = res0
			l.SetMetatable(ud, l.GetTypeMetatable("BehaviourAction"))
			l.Push(ud)

			return 1
		},
		"init": func(l *lua2.LState) int {
			objInterface := lua.CheckInterfaceValue[IUseItemTarget](l, 1)
			obj := objInterface.GetUseItemTarget()
			obj.Init(
				lua.CheckReferenceValue[byter.Byter](l, 2),
			)

			return 0
		},
	})
}

func (u *UseItemTarget) ToLua(l *lua2.LState) lua2.LValue {
	ud := l.NewUserData()
	ud.Value = u

	l.SetMetatable(ud, l.GetTypeMetatable("UseItemTarget"))
	return ud
}
