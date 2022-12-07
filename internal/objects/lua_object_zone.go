package objects

import (
	lua2 "RainbowRunner/internal/lua"
	"RainbowRunner/pkg/datatypes"
	lua "github.com/yuin/gopher-lua"
	"strings"
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

func (f ZoneLuaFunctions) LoadNPCFromConfig(s *lua.LState) int {
	z := lua2.CheckReferenceValue[Zone](s, 1)
	npcID := s.CheckString(2)

	npcConfig, ok := z.BaseConfig.NPCs[strings.ToLower(npcID)]

	if !ok {
		s.RaiseError("could not find NPC config with ID: " + npcID)
	}

	npc := NewNPCFromConfig(npcConfig)

	ud := s.NewUserData()
	ud.Value = npc

	s.SetMetatable(ud, s.GetTypeMetatable(luaNPCTypeName))
	s.Push(ud)
	return 1
}

func (f ZoneLuaFunctions) Spawn(state *lua.LState) int {
	ud := state.CheckUserData(2)

	switch ud.Value.(type) {
	case *NPC:
		return LuaZoneSpawnEntity(state)
	case *WorldEntity:
		return LuaZoneSpawnEntity(state)
	default:
		state.RaiseError("cannot spawn given entity type ")
	}

	return 0
}

func LuaZoneSpawnEntity(state *lua.LState) int {
	z := lua2.CheckReferenceValue[Zone](state, 1)
	entity := lua2.CheckInterfaceValue[DRObject](state, 2)

	var position datatypes.Vector3Float32
	var rotation float32

	if state.GetTop() >= 3 {
		position = lua2.CheckValue[datatypes.Vector3Float32](state, 3)
	}

	if state.GetTop() >= 4 {
		rotation = float32(state.CheckNumber(4))
	}

	z.SpawnInit(entity, &position, &rotation)

	return 0
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
	"spawn":             zoneLuaFunctions.Spawn,
	"loadNPCFromConfig": zoneLuaFunctions.LoadNPCFromConfig,
}
