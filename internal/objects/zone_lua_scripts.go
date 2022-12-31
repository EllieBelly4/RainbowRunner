package objects

import (
	"RainbowRunner/internal/lua"
	"RainbowRunner/internal/script"
	drobjectypes "RainbowRunner/internal/types/drobjecttypes"
	log "github.com/sirupsen/logrus"
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

	if script != nil {
		err := script.Execute(s.State)

		if err != nil {
			log.Errorf("Error executing standalone init script for zone %s", err.Error())
		}
	}

	err := s.EntityScript.Init(entity)

	if err != nil {
		return err
	}

	return nil
}

func (s *ZoneLuaScripts) OnPlayerEnter(player *Player) {
	s.CallEventHandler("onPlayerEnter", player)
}

func NewZoneLuaScripts(z *Zone) *ZoneLuaScripts {
	luaState := lua2.NewState()
	RegisterLuaGlobals(luaState)

	luaState.SetGlobal("currentZone", z.ToLua(luaState))

	scriptGroup := lua.GetScriptGroup("zones." + strings.ToLower(z.Name))
	defaultScriptGroup := lua.GetScriptGroup("zones.__default")

	main := scriptGroup.Get("main")

	if main == nil {
		main = defaultScriptGroup.Get("main")
	}

	zoneLuaScripts := &ZoneLuaScripts{
		scriptGroup: scriptGroup,
	}

	zoneLuaScripts.EntityScript = script.NewEntityScript(
		main,
		luaState,
	)

	return zoneLuaScripts
}
