package objects

import (
	lua2 "github.com/yuin/gopher-lua"
)

const luaGCObjectTypeName = "GCObject"

func registerLuaGCObject(state *lua2.LState) {
	mt := state.NewTypeMetatable(luaGCObjectTypeName)
	state.SetGlobal("GCObject", mt)
	state.SetField(mt, "new", state.NewFunction(newLuaGCObject))
	state.SetField(mt, "__index", state.SetFuncs(state.NewTable(),
		luaDRObjectExtendMethods(entityLuaGCObjectMethods)),
	)
}

var entityLuaGCObjectMethods = map[string]lua2.LGFunction{}

func newLuaGCObject(l *lua2.LState) int {
	gcObject := NewGCObject(l.CheckString(1))

	ud := l.NewUserData()
	ud.Value = gcObject

	l.SetMetatable(ud, l.GetTypeMetatable(luaGCObjectTypeName))
	l.Push(ud)
	return 1
}
