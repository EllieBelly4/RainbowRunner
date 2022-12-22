package script

import "RainbowRunner/internal/lua"

type IEntityScript interface {
	Init()
	Tick()
}

type EntityScript struct {
	luaScript *lua.LuaScript
}

func (e EntityScript) Init() {
}

func (e EntityScript) Tick() {
}

func NewEntityScript(luaScript *lua.LuaScript) *EntityScript {
	return &EntityScript{
		luaScript: luaScript,
	}
}
