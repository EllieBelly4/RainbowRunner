package objects

import (
	"RainbowRunner/internal/lua"
	log "github.com/sirupsen/logrus"
	lua2 "github.com/yuin/gopher-lua"
	"strings"
)

type ZoneLuaScripts struct {
	scriptGroup   *lua.LuaScriptGroup
	State         *lua2.LState
	main          *lua.LuaScript
	eventHandlers map[string]*lua2.LFunction
	tick          *lua2.LFunction
}

func (s *ZoneLuaScripts) ExecuteInit() error {
	if s.main != nil {
		err := s.main.Execute(s.State)

		if err != nil {
			log.Error(err)
			goto aftermain
		}

		if s.State.GetTop() == 2 {
			mod := s.State.CheckTable(2)

			s.registerEventHandlers(mod)
		}
	}

aftermain:

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

func (s *ZoneLuaScripts) ExecuteTick() error {
	if s.tick == nil {
		return nil
	}

	err := s.State.CallByParam(lua2.P{
		Fn:      s.tick,
		NRet:    0,
		Protect: true,
	})

	return err
}

func (s *ZoneLuaScripts) registerEventHandlers(mod *lua2.LTable) {
	mod.ForEach(func(key lua2.LValue, value lua2.LValue) {
		if value.Type() != lua2.LTFunction {
			return
		}

		keyString := key.String()

		if !strings.HasPrefix(keyString, "__") {
			return
		}

		eventName := strings.TrimPrefix(keyString, "__")

		s.eventHandlers[eventName] = value.(*lua2.LFunction)
	})

	tick, ok := s.eventHandlers["tick"]

	if ok {
		s.tick = tick
	}
}

func NewZoneLuaScripts(z *Zone) *ZoneLuaScripts {
	luaState := lua2.NewState()
	RegisterLuaGlobals(luaState)

	luaState.SetGlobal("currentZone", z.ToLua(luaState))
	scriptGroup := lua.GetScriptGroup("zones." + strings.ToLower(z.Name))

	zoneLuaScripts := &ZoneLuaScripts{
		State:         luaState,
		scriptGroup:   scriptGroup,
		eventHandlers: map[string]*lua2.LFunction{},
	}

	main := scriptGroup.Get("main")

	if main != nil {
		zoneLuaScripts.main = main
	}

	return zoneLuaScripts
}
