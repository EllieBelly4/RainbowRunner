package lua

import (
	lua2 "github.com/yuin/gopher-lua"
	"io"
	"os"
)

//go:generate go run ../../scripts/generatelua -type LuaScript
type LuaScript struct {
	path       string
	scriptText string
	id         string
}

func (s *LuaScript) Execute(state *lua2.LState) error {
	s.load()

	err := state.DoString(s.scriptText)

	if err != nil {
		return err
	}

	return nil
}

func (s *LuaScript) load() {
	//TODO add optional caching
	fh, err := os.Open(s.path)

	if err != nil {
		panic(err)
	}

	data, err := io.ReadAll(fh)

	if err != nil {
		panic(err)
	}

	s.scriptText = string(data)
}

func (s *LuaScript) RegisterModule(state *lua2.LState) {
	s.load()

	mod, err := state.LoadString(s.scriptText)

	if err != nil {
		return
	}

	preload := state.GetField(state.GetField(state.Get(lua2.EnvironIndex), "package"), "preload")
	state.SetField(preload, s.id, mod)
}

func NewLuaScript(path string, id string) *LuaScript {
	return &LuaScript{
		path: path,
		id:   id,
	}
}
