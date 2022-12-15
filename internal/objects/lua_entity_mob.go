package objects

import (
	"RainbowRunner/internal/lua"
	lua2 "github.com/yuin/gopher-lua"
)

const luaMOBTypeName = "MOB"

func registerLuaMOB(state *lua2.LState) {
	mt := state.NewTypeMetatable(luaMOBTypeName)
	state.SetGlobal("MOB", mt)
	state.SetField(mt, "new", state.NewFunction(newLuaMOB))
	state.SetField(mt, "__index", state.SetFuncs(state.NewTable(),
		luaMethodsMOB(),
	))
}

func luaMethodsMOB() map[string]lua2.LGFunction {
	return luaMethodsExtend(entityLuaMOBMethods, luaMethodsUnit)
}

var entityLuaMOBMethods = map[string]lua2.LGFunction{
	"name": entityLuaMOBGetSetName,
}

func entityLuaMOBGetSetName(state *lua2.LState) int {
	mob := lua.CheckReferenceValue[MOB](state, 1)

	if state.GetTop() == 2 {
		mob.Name = state.CheckString(2)
		return 0
	}

	state.Push(lua2.LString(mob.Name))
	return 1
}

func newLuaMOB(l *lua2.LState) int {
	mob := NewMOBSimple(l.CheckString(1))

	ud := l.NewUserData()
	ud.Value = mob

	l.SetMetatable(ud, l.GetTypeMetatable(luaMOBTypeName))
	l.Push(ud)
	return 1
}
