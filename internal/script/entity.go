package script

import (
	"RainbowRunner/internal/lua"
	lua2 "github.com/yuin/gopher-lua"
	"strings"
)

type IEntityScript interface {
	Init() error
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

func (e *EntityScript) Init() error {
	if init, ok := e.EventHandlers["init"]; ok {
		err := e.State.CallByParam(lua2.P{
			Fn:      init,
			NRet:    0,
			Protect: true,
		})

		if err != nil {
			return err
		}
	}

	return nil
}

func (e *EntityScript) Tick() error {
	if e.tick == nil {
		return nil
	}

	err := e.State.CallByParam(lua2.P{
		Fn:      e.tick,
		NRet:    0,
		Protect: true,
	})

	return err
}

func NewEntityScript(luaScript *lua.LuaScript, state *lua2.LState) *EntityScript {
	return &EntityScript{
		luaScript:     luaScript,
		State:         state,
		EventHandlers: map[string]*lua2.LFunction{},
	}
}
