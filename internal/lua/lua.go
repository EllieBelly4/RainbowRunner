package lua

import (
	log "github.com/sirupsen/logrus"
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

		scriptID := strings.Join(append(splitPath[1:splitPathLength-1], scriptName), ".")
		log.Infof("loading lua script %s", scriptID)

		currentScriptGroup := getOrCreateScriptGroup(splitPath[1:splitPathLength-1], scripts)

		currentScriptGroup.scripts[scriptName] = NewLuaScript(path, scriptID)

		return nil
	})

	if err != nil {
		return err
	}

	return nil
}

func GetScript(path string) *LuaScript {
	splitPath := strings.Split(path, ".")

	if len(splitPath) < 2 {
		log.Errorf("getting a script requires at least 2 parts $DIR.$SCRIPTNAME got: %s", path)
		return nil
	}

	modulePath := strings.Join(splitPath[:len(splitPath)-1], ".")
	group := GetScriptGroup(modulePath)

	if group == nil {
		log.Errorf("could not find script group for %s", modulePath)
		return nil
	}

	script, ok := group.scripts[splitPath[len(splitPath)-1]]

	if len(group.scripts) == 0 || !ok {
		log.Errorf("could not find script %s in group %s", script, modulePath)
		return nil
	}

	return script
}

// GetScriptGroup
// path: '.' separated path to script file without file extension
func GetScriptGroup(path string) *LuaScriptGroup {
	sg := getScriptGroup(strings.Split(path, "."), scripts)
	return sg
}

func getOrCreateScriptGroup(i []string, sgs map[string]*LuaScriptGroup) *LuaScriptGroup {
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

		return getOrCreateScriptGroup(i[1:], sg.children)
	}

	return sg
}

func getScriptGroup(i []string, sgs map[string]*LuaScriptGroup) *LuaScriptGroup {
	if len(i) == 0 {
		return nil
	}

	sg, ok := sgs[i[0]]

	if !ok {
		return nil
	}

	if len(i) != 1 {
		if sg.children == nil {
			sg.children = make(map[string]*LuaScriptGroup)
		}

		return getOrCreateScriptGroup(i[1:], sg.children)
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
