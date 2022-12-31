package script

import (
	"RainbowRunner/internal/lua"
	"RainbowRunner/internal/types/drobjecttypes"
	log "github.com/sirupsen/logrus"
	lua2 "github.com/yuin/gopher-lua"
	"strings"
)

type IEntityScript interface {
	Init(entity drobjecttypes.DRObject) error
	Tick() error
	Load() error
}

type EntityScript struct {
	EventHandlers map[string]*lua2.LFunction
	State         *lua2.LState
	luaScript     *lua.LuaScript
	tick          *lua2.LFunction
}

func (e *EntityScript) RegisterEventHandlers(mod *lua2.LTable) {
	mod.ForEach(func(key lua2.LValue, value lua2.LValue) {
		if value.Type() != lua2.LTFunction {
			return
		}

		keyString := key.String()

		if !strings.HasPrefix(keyString, "__") {
			return
		}

		eventName := strings.TrimPrefix(keyString, "__")

		e.EventHandlers[eventName] = value.(*lua2.LFunction)
	})

	tick, ok := e.EventHandlers["tick"]

	if ok {
		e.tick = tick
	}
}

func (e *EntityScript) Load() error {
	if e == nil || e.luaScript == nil {
		return nil
	}

	preTop := e.State.GetTop()

	err := e.luaScript.Execute(e.State)

	if err != nil {
		return err
	}

	if e.State.GetTop() > preTop {
		mod := e.State.CheckTable(preTop + 1)

		e.RegisterEventHandlers(mod)
	}

	return nil
}

func (e *EntityScript) Init(entity drobjecttypes.DRObject) error {
	if e == nil || e.EventHandlers == nil {
		return nil
	}

	if init, ok := e.EventHandlers["init"]; ok {
		entityLua := lua2.LNil

		if entity != nil {
			entityLua = entity.ToLua(e.State)
		}

		err := e.State.CallByParam(lua2.P{
			Fn:      init,
			NRet:    0,
			Protect: true,
		}, entityLua)

		if err != nil {
			return err
		}
	}

	return nil
}

func (e *EntityScript) Tick() error {
	if e == nil || e.tick == nil {
		return nil
	}

	err := e.State.CallByParam(lua2.P{
		Fn:      e.tick,
		NRet:    0,
		Protect: true,
	})

	return err
}

func (e *EntityScript) CallEventHandler(eventHandlerName string, args ...interface{}) {
	if e == nil || e.EventHandlers == nil {
		return
	}

	if eventHandler, ok := e.EventHandlers[eventHandlerName]; ok {
		luaArgs := make([]lua2.LValue, len(args))

		for i, arg := range args {
			luaArgs[i] = lua.ValueToLValue(e.State, arg)
		}

		err := e.State.CallByParam(lua2.P{
			Fn:      eventHandler,
			NRet:    0,
			Protect: true,
		}, luaArgs...)

		if err != nil {
			log.Error(err)
		}
	} else {
		log.Warnf("Event handler %s not found", eventHandlerName)
	}
}

func NewEntityScript(luaScript *lua.LuaScript, state *lua2.LState) *EntityScript {
	return &EntityScript{
		luaScript:     luaScript,
		State:         state,
		EventHandlers: map[string]*lua2.LFunction{},
	}
}
