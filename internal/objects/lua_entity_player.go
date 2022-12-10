package objects

import (
	"RainbowRunner/internal/lua"
	log "github.com/sirupsen/logrus"
	lua2 "github.com/yuin/gopher-lua"
)

const luaPlayerTypeName = "Player"

func registerLuaPlayer(state *lua2.LState) {
	mt := state.NewTypeMetatable(luaPlayerTypeName)
	state.SetGlobal("Player", mt)
	state.SetField(mt, "new", state.NewFunction(newLuaPlayer))
	state.SetField(mt, "__index", state.SetFuncs(state.NewTable(),
		luaMethodsPlayer(),
	))
}

func luaMethodsPlayer() map[string]lua2.LGFunction {
	return luaMethodsExtend(map[string]lua2.LGFunction{
		"name":       luaGenericGetSetString[*Player](func(v *Player) *string { return &v.Name }),
		"changeZone": luaPlayerChangeZone,
	}, luaMethodsDRObject)
}

func luaPlayerChangeZone(state *lua2.LState) int {
	player := lua.CheckReferenceValue[Player](state, 1)

	zoneName := state.CheckString(2)

	rrPlayer := Players.GetPlayer(uint16(player.ID()))

	tZone := Zones.GetOrCreateZone(zoneName)

	if tZone == nil {
		log.Errorf("could not find zone %s", zoneName)
		return 0
	}

	rrPlayer.JoinZone(tZone)

	return 0
}

func LuaPlayerToCustomData(l *lua2.LState, player *Player) *lua2.LUserData {
	ud := l.NewUserData()
	ud.Value = player

	l.SetMetatable(ud, l.GetTypeMetatable(luaPlayerTypeName))
	return ud
}

func newLuaPlayer(l *lua2.LState) int {
	player := NewPlayer(l.CheckString(1))

	ud := LuaPlayerToCustomData(l, player)
	l.Push(ud)
	return 1
}
