package objects

import (
	"RainbowRunner/internal/lua"
	lua2 "github.com/yuin/gopher-lua"
	"strings"
)

type ZoneLuaScripts struct {
	scriptGroup *lua.LuaScriptGroup
	State       *lua2.LState
}

func (s ZoneLuaScripts) ExecuteInit() error {
	script := s.scriptGroup.Get("init")

	if script == nil {
		return nil
	}

	err := script.Execute(s.State)

	if err != nil {
		return err
	}

	return nil
}

func NewZoneLuaScripts(z *Zone) *ZoneLuaScripts {
	luaState := lua2.NewState()
	RegisterLuaGlobals(luaState)

	luaState.SetGlobal("currentZone", z.ToLua(luaState))
	return &ZoneLuaScripts{
		State:       luaState,
		scriptGroup: lua.GetScriptGroup("zones." + strings.ToLower(z.Name)),
	}
}
