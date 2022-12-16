// Code generated by scripts/generatelua DO NOT EDIT.
package actions

import (
	lua "RainbowRunner/internal/lua"
	byter "RainbowRunner/pkg/byter"
	lua2 "github.com/yuin/gopher-lua"
)

type IHide interface {
	GetHide() *Hide
}

func (h *Hide) GetHide() *Hide {
	return h
}

func registerLuaHide(state *lua2.LState) {
	// Ensure the import is referenced in code
	_ = lua.LuaScript{}

	mt := state.NewTypeMetatable("Hide")
	state.SetGlobal("Hide", mt)
	state.SetField(mt, "__index", state.SetFuncs(state.NewTable(),
		luaMethodsHide(),
	))
}

func luaMethodsHide() map[string]lua2.LGFunction {
	return lua.LuaMethodsExtend(map[string]lua2.LGFunction{
		"opCode": func(l *lua2.LState) int {
			objInterface := lua.CheckInterfaceValue[IHide](l, 1)
			obj := objInterface.GetHide()
			res0 := obj.OpCode()
			ud := l.NewUserData()
			ud.Value = res0
			l.SetMetatable(ud, l.GetTypeMetatable("BehaviourAction"))
			l.Push(ud)

			return 1
		},
		"init": func(l *lua2.LState) int {
			objInterface := lua.CheckInterfaceValue[IHide](l, 1)
			obj := objInterface.GetHide()
			obj.Init(
				lua.CheckReferenceValue[byter.Byter](l, 2),
			)

			return 0
		},
	})
}

func (h *Hide) ToLua(l *lua2.LState) lua2.LValue {
	ud := l.NewUserData()
	ud.Value = h

	l.SetMetatable(ud, l.GetTypeMetatable("Hide"))
	return ud
}