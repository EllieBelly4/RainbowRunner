// Code generated by scripts/generatelua DO NOT EDIT.
package actions

import (
	lua "RainbowRunner/internal/lua"
	byter "RainbowRunner/pkg/byter"
	lua2 "github.com/yuin/gopher-lua"
)

type ISearchForAttack interface {
	GetSearchForAttack() *SearchForAttack
}

func (s *SearchForAttack) GetSearchForAttack() *SearchForAttack {
	return s
}

func registerLuaSearchForAttack(state *lua2.LState) {
	// Ensure the import is referenced in code
	_ = lua.LuaScript{}

	mt := state.NewTypeMetatable("SearchForAttack")
	state.SetGlobal("SearchForAttack", mt)
	state.SetField(mt, "__index", state.SetFuncs(state.NewTable(),
		luaMethodsSearchForAttack(),
	))
}

func luaMethodsSearchForAttack() map[string]lua2.LGFunction {
	return lua.LuaMethodsExtend(map[string]lua2.LGFunction{
		"opCode": func(l *lua2.LState) int {
			objInterface := lua.CheckInterfaceValue[ISearchForAttack](l, 1)
			obj := objInterface.GetSearchForAttack()
			res0 := obj.OpCode()
			ud := l.NewUserData()
			ud.Value = res0
			l.SetMetatable(ud, l.GetTypeMetatable("BehaviourAction"))
			l.Push(ud)

			return 1
		},
		"init": func(l *lua2.LState) int {
			objInterface := lua.CheckInterfaceValue[ISearchForAttack](l, 1)
			obj := objInterface.GetSearchForAttack()
			obj.Init(
				lua.CheckReferenceValue[byter.Byter](l, 2),
			)

			return 0
		},
	})
}

func (s *SearchForAttack) ToLua(l *lua2.LState) lua2.LValue {
	ud := l.NewUserData()
	ud.Value = s

	l.SetMetatable(ud, l.GetTypeMetatable("SearchForAttack"))
	return ud
}