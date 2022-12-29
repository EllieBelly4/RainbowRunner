package pathfinding

import (
	"RainbowRunner/internal/types"
	"github.com/goccy/go-json"
	"io"
	"os"
	"path/filepath"
	"strings"
)

var pathMapCache = make(map[string]*types.PathMap)

func ReloadPathMap(zoneName string) *types.PathMap {
	lcZoneName := strings.ToLower(zoneName)

	if _, ok := pathMapCache[lcZoneName]; ok {
		delete(pathMapCache, lcZoneName)
	}

	return LoadPathMap(zoneName)
}

func LoadPathMap(zoneName string) *types.PathMap {
	lcZoneName := strings.ToLower(zoneName)
	filePath := filepath.Join("data", "pathmaps", lcZoneName+"_pathmap.json")

	stat, err := os.Stat(filePath)

	if err != nil {
		if os.IsNotExist(err) {
			return nil
		}

		panic(err)
	}

	if stat.IsDir() {
		panic("pathmap file is a directory: " + filePath)
	}

	file, err := os.Open(filePath)

	if err != nil {
		panic(err)
	}

	var pathMap = &types.PathMap{}

	data, err := io.ReadAll(file)

	if err != nil {
		panic(err)
	}

	err = json.Unmarshal(data, pathMap)

	if err != nil {
		panic(err)
	}

	pathMap.Init()

	pathMapCache[lcZoneName] = pathMap

	return pathMap
}
