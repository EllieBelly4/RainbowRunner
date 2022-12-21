package lua

import lua2 "github.com/yuin/gopher-lua"

func RegisterModules(l *lua2.LState) {
	l.SetGlobal("module", l.NewTable())

	walkScripts(func(script *LuaScript) {
		script.RegisterModule(l)
	})
}

func walkScripts(f func(group *LuaScript)) {
	for _, group := range scripts {
		walkScriptsRecursive(group, f)
	}
}

func walkScriptsRecursive(group *LuaScriptGroup, f func(group *LuaScript)) {
	for _, script := range group.scripts {
		f(script)
	}

	for _, child := range group.children {
		walkScriptsRecursive(child, f)
	}
}
