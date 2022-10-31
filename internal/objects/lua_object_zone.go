package objects

import (
	lua2 "RainbowRunner/internal/lua"
	"RainbowRunner/pkg/datatypes"
	lua "github.com/yuin/gopher-lua"
)

type ZoneLuaFunctions struct {
}

func (f ZoneLuaFunctions) GetName(state *lua.LState) int {
	z := lua2.CheckReferenceValue[Zone](state, 1)

	if state.GetTop() == 2 {
		state.ArgError(2, "variable is readonly")
		return 0
	}

	state.Push(lua.LString(z.Name))
	return 1
}

func (f ZoneLuaFunctions) SpawnNPC(state *lua.LState) int {
	z := lua2.CheckReferenceValue[Zone](state, 1)
	npc := lua2.CheckReferenceValue[NPC](state, 2)

	var position datatypes.Vector3Float32
	var rotation float32

	if state.GetTop() >= 3 {
		position = lua2.CheckValue[datatypes.Vector3Float32](state, 3)
	}

	if state.GetTop() >= 4 {
		rotation = float32(state.CheckNumber(4))
	}

	z.SpawnInit(npc, &position, &rotation)

	return 0
}

func (f ZoneLuaFunctions) LoadNPCFromConfig(s *lua.LState) int {
	z := lua2.CheckReferenceValue[Zone](s, 1)
	npcID := s.CheckString(2)

	npcConfig, ok := z.BaseConfig.NPCs[npcID]

	if !ok {
		s.Error(lua.LString("could not find NPC config with ID: "+npcID), -1)
	}

	npc := NewNPCFromConfig(npcConfig)

	ud := s.NewUserData()
	ud.Value = npc

	s.SetMetatable(ud, s.GetTypeMetatable(luaNPCTypeName))
	s.Push(ud)
	return 1
}

var zoneLuaFunctions ZoneLuaFunctions

const luaZoneTypeName = "Zone"

func registerLuaZone(s *lua.LState) {
	mt := s.NewTypeMetatable(luaZoneTypeName)
	s.SetGlobal("Zone", mt)
	s.SetField(mt, "__index", s.SetFuncs(s.NewTable(), zoneMethods))
}

func AddZoneToState(L *lua.LState, z *Zone) {
	ud := L.NewUserData()
	ud.Value = z
	L.SetMetatable(ud, L.GetTypeMetatable(luaZoneTypeName))
	L.SetGlobal("currentZone", ud)
}

var zoneMethods = map[string]lua.LGFunction{
	"name":              zoneLuaFunctions.GetName,
	"spawnNPC":          zoneLuaFunctions.SpawnNPC,
	"loadNPCFromConfig": zoneLuaFunctions.LoadNPCFromConfig,
}

//// Constructor
//func newPerson(L *lua.LState) int {
//	person := &Person{L.CheckString(1)}
//	ud := L.NewUserData()
//	ud.Value = person
//	L.SetMetatable(ud, L.GetTypeMetatable(luaZoneTypeName))
//	L.Push(ud)
//	return 1
//}

// Checks whether the first lua argument is a *LUserData with *Person and returns this *Person.
//func checkPerson(L *lua.LState) *Person {
//	ud := L.CheckUserData(1)
//	if v, ok := ud.Value.(*Person); ok {
//		return v
//	}
//	L.ArgError(1, "person expected")
//	return nil
//}

//var personMethods = map[string]lua.LGFunction{
//	"name": personGetSetName,
//}

//// Getter and setter for the Person#Name
//func personGetSetName(L *lua.LState) int {
//	p := checkPerson(L)
//	if L.GetTop() == 2 {
//		p.Name = L.CheckString(2)
//		return 0
//	}
//	L.Push(lua.LString(p.Name))
//	return 1
//}
