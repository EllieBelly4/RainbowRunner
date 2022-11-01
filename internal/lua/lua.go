package lua

import (
	"io/fs"
	"path/filepath"
	"strings"
)

var scripts = make(map[string]*LuaScriptGroup)
var rootPath string

func LoadScripts(root string) error {
	rootPath = root

	return reloadScripts()
}

func reloadScripts() error {
	err := filepath.WalkDir(rootPath, func(path string, d fs.DirEntry, err error) error {
		if d.IsDir() {
			return nil
		}

		splitPath := strings.Split(path, "\\")
		splitPathLength := len(splitPath)
		fileName := splitPath[splitPathLength-1]
		scriptName := strings.Split(fileName, ".")[0]

		currentScriptGroup := getScriptGroup(splitPath[1:splitPathLength-1], scripts)

		currentScriptGroup.scripts[scriptName] = NewLuaScript(path)

		return nil
	})

	if err != nil {
		return err
	}

	return nil
}

// GetScriptGroup
// path: '.' separated path to script file without file extension
func GetScriptGroup(path string) *LuaScriptGroup {
	return getScriptGroup(strings.Split(path, "."), scripts)
}

func getScriptGroup(i []string, sgs map[string]*LuaScriptGroup) *LuaScriptGroup {
	if len(i) == 0 {
		return nil
	}

	sg, ok := sgs[i[0]]

	if !ok {
		sgs[i[0]] = NewScriptGroup(i[0])
		sg = sgs[i[0]]
	}

	if len(i) != 1 {
		if sg.children == nil {
			sg.children = make(map[string]*LuaScriptGroup)
		}

		return getScriptGroup(i[1:], sg.children)
	}

	return sg
}

func NewScriptGroup(name string) *LuaScriptGroup {
	return &LuaScriptGroup{
		Name:     name,
		scripts:  make(map[string]*LuaScript),
		children: nil,
	}
}
