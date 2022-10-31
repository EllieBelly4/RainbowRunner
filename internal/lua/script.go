package lua

import (
	"github.com/yuin/gopher-lua"
	"io"
	"os"
)

type LuaScript struct {
	path       string
	scriptText string
}

func (s *LuaScript) Execute(state *lua.LState) {
	s.load()

	err := state.DoString(s.scriptText)

	if err != nil {
		panic(err)
	}
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

func NewLuaScript(path string) *LuaScript {
	return &LuaScript{
		path: path,
	}
}
