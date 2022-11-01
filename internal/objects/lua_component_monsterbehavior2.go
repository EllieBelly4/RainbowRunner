package objects

import (
	lua2 "github.com/yuin/gopher-lua"
)

const luaMonsterBehavior2TypeName = "MonsterBehavior2"

func registerLuaMonsterBehavior2(state *lua2.LState) {
	mt := state.NewTypeMetatable(luaMonsterBehavior2TypeName)
	state.SetGlobal("MonsterBehavior2", mt)
	state.SetField(mt, "new", state.NewFunction(newLuaMonsterBehavior2))
	state.SetField(mt, "__index", state.SetFuncs(state.NewTable(), luaMethodsMonsterBehavior2()))
}

func luaMethodsMonsterBehavior2() map[string]lua2.LGFunction {
	return luaMethodsExtend(
		entityLuaMonsterBehavior2Methods,
		luaMethodsComponent,
	)
}

var entityLuaMonsterBehavior2Methods = map[string]lua2.LGFunction{}

func newLuaMonsterBehavior2(l *lua2.LState) int {
	MonsterBehavior2 := NewMonsterBehavior2(l.CheckString(1))

	ud := l.NewUserData()
	ud.Value = MonsterBehavior2

	l.SetMetatable(ud, l.GetTypeMetatable(luaMonsterBehavior2TypeName))
	l.Push(ud)
	return 1
}
