package lua

type LuaScriptGroup struct {
	Name     string
	scripts  map[string]*LuaScript
	children map[string]*LuaScriptGroup
}

func (g LuaScriptGroup) Get(s string) *LuaScript {
	if override, ok := g.scripts[s+"_override"]; ok {
		return override
	}

	if script, ok := g.scripts[s]; ok {
		return script
	}

	return nil
}
