package objects

import (
	lua2 "github.com/yuin/gopher-lua"
)

const luaComponentTypeName = "Component"

func registerLuaComponent(state *lua2.LState) {
	mt := state.NewTypeMetatable(luaComponentTypeName)
	state.SetGlobal("Component", mt)
	state.SetField(mt, "new", state.NewFunction(newLuaComponent))
	state.SetField(mt, "__index", state.SetFuncs(state.NewTable(),
		luaDRObjectExtendMethods(entityLuaComponentMethods)),
	)
}

var entityLuaComponentMethods = map[string]lua2.LGFunction{}

func newLuaComponent(l *lua2.LState) int {
	component := NewComponent(l.CheckString(1), l.CheckString(2))

	ud := l.NewUserData()
	ud.Value = component

	l.SetMetatable(ud, l.GetTypeMetatable(luaComponentTypeName))
	l.Push(ud)
	return 1
}
