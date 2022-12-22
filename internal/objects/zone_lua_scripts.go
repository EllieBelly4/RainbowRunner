package objects

import (
	"RainbowRunner/internal/lua"
	"RainbowRunner/internal/script"
	lua2 "github.com/yuin/gopher-lua"
	"strings"
)

type ZoneLuaScripts struct {
	*script.EntityScript
	scriptGroup *lua.LuaScriptGroup
	main        *lua.LuaScript
}

func (s *ZoneLuaScripts) Init() error {
	script := s.scriptGroup.Get("init")

	if script == nil {
		return nil
	}

	err := script.Execute(s.State)

	if err != nil {
		return err
	}

	err = s.EntityScript.Init()

	if err != nil {
		return err
	}

	return nil
}

func NewZoneLuaScripts(z *Zone) *ZoneLuaScripts {
	luaState := lua2.NewState()
	RegisterLuaGlobals(luaState)

	luaState.SetGlobal("currentZone", z.ToLua(luaState))
	scriptGroup := lua.GetScriptGroup("zones." + strings.ToLower(z.Name))

	main := scriptGroup.Get("main")

	zoneLuaScripts := &ZoneLuaScripts{
		EntityScript: script.NewEntityScript(
			main,
			luaState,
		),
		scriptGroup: scriptGroup,
	}

	return zoneLuaScripts
}
