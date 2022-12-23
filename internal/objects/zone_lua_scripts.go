package objects

import (
	"RainbowRunner/internal/lua"
	"RainbowRunner/internal/script"
	drobjectypes "RainbowRunner/internal/types/drobjecttypes"
	lua2 "github.com/yuin/gopher-lua"
	"strings"
)

type ZoneLuaScripts struct {
	*script.EntityScript
	scriptGroup *lua.LuaScriptGroup
	main        *lua.LuaScript
}

func (s *ZoneLuaScripts) Tick() error {
	if s.EntityScript == nil {
		return nil
	}

	return s.EntityScript.Tick()
}

func (s *ZoneLuaScripts) Load() error {
	if s.EntityScript == nil {
		return nil
	}

	return s.EntityScript.Load()
}

func (s *ZoneLuaScripts) Init(entity drobjectypes.DRObject) error {
	script := s.scriptGroup.Get("init")

	if script == nil {
		return nil
	}

	err := script.Execute(s.State)

	if err != nil {
		return err
	}

	err = s.EntityScript.Init(entity)

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
		scriptGroup: scriptGroup,
	}

	zoneLuaScripts.EntityScript = script.NewEntityScript(
		main,
		luaState,
	)

	return zoneLuaScripts
}
