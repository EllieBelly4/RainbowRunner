package objects

import (
	"RainbowRunner/internal/lua"
	lua2 "github.com/yuin/gopher-lua"
)

const luaNPCTypeName = "NPC"

func registerLuaNPC(state *lua2.LState) {
	mt := state.NewTypeMetatable(luaNPCTypeName)
	state.SetGlobal("NPC", mt)
	state.SetField(mt, "new", state.NewFunction(newLuaNPC))
	state.SetField(mt, "__index", state.SetFuncs(state.NewTable(), entityLuaNPCMethods))
}

var entityLuaNPCMethods = map[string]lua2.LGFunction{
	"name": entityLuaNPCGetSetName,
}

func entityLuaNPCGetSetName(state *lua2.LState) int {
	npc := lua.CheckReferenceValue[NPC](state, 1)

	if state.GetTop() == 2 {
		npc.Name = state.CheckString(2)
		return 0
	}

	state.Push(lua2.LString(npc.Name))
	return 1
}

func newLuaNPC(l *lua2.LState) int {
	npc := NewNPCSimple(l.CheckString(1))

	ud := l.NewUserData()
	ud.Value = npc

	l.SetMetatable(ud, l.GetTypeMetatable(luaNPCTypeName))
	l.Push(ud)
	return 1
}
