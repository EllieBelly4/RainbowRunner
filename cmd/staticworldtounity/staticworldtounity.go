package main

import (
	"RainbowRunner/cmd/configparser/configparser"
	"RainbowRunner/cmd/rrcli/configurator"
	"RainbowRunner/cmd/rrcli/modelextractor"
	"RainbowRunner/internal/database"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
	"path"
	"strings"
)

type StaticObjectDumpLocation struct {
	Rotation float32 `json:"rotation,omitempty"`
	X        float32 `json:"x,omitempty"`
	Y        float32 `json:"y,omitempty"`
	Z        float32 `json:"z,omitempty"`
}

type StaticObjectDump struct {
	GCTypeHash int                      `json:"gctypehash"`
	GCTypeName string                   `json:"gctypename"`
	Position   StaticObjectDumpLocation `json:"position"`
}

func main() {
	filePath := "resources/Dumps/Objects/townston_staticobjects.json"
	rootPath := "D:\\Work\\dungeon-runners\\666 dumps new"
	destPath := "E:\\Unity\\RainbowRunnerGameBackend\\Assets\\Resources\\ModelDumps\\Townston"

	file, err := os.Open(filePath)

	if err != nil {
		panic(err)
	}

	data, err := ioutil.ReadAll(file)

	if err != nil {
		return
	}

	var dump []StaticObjectDump

	err = json.Unmarshal(data, &dump)

	if err != nil {
		panic(err)
	}

	config, err := configurator.LoadFromDumpedConfigFile("C:\\Users\\Sophie\\go\\src\\RainbowRunner\\resources\\Dumps\\generated\\finalconf.json")

	if err != nil {
		panic(err)
	}

	allEntities := getDRClassesForStaticObjects(dump, config)

	data, err = json.MarshalIndent(allEntities, "", "  ")

	if err != nil {
		panic(err)
	}

	err = ioutil.WriteFile("resources/Dumps/generated/staticworlddtounity.json", data, 0755)

	if err != nil {
		panic(err)
	}

	extractModelsForObjects(allEntities, rootPath, destPath)
}

func extractModelsForObjects(entities []struct {
	string
	*database.DRClass
}, rootPath string, destPath string) {
	useCollisionModel := false

	childName := "visual"
	propName := "Visual"

	if useCollisionModel {
		childName = "description"
		propName = "CollisionObject"
	}

	for _, entityPair := range entities {
		entityGCType := entityPair.string
		entity := entityPair.DRClass
		if entity.Name == "blacksmith_boiler" ||
			entity.Name == "waterfall" ||
			entity.Name == "fireflies_1" ||
			entity.Name == "elmforest_lightray_4" ||
			entity.Name == "elmforest_lightray_1" ||
			entity.Name == "fireflies_2" {
			fmt.Printf("skipping %s because we can't parse the 3dnode yet\n", entity.Name)
			continue
		}

		if entity.Children == nil {
			continue
		}

		if _, ok := entity.Children[childName]; !ok {
			continue
		}

		desc := entity.Children[childName].Entities[0]

		if desc.Properties == nil {
			continue
		}

		descProps := desc.Properties

		if _, ok := descProps[propName]; !ok {
			fmt.Println(fmt.Sprintf("%s does not have a visual object", entity.Name))
			continue
		}

		collisionObjectName := descProps.StringVal(propName)

		filePath := path.Join(rootPath, collisionObjectName+".3dnode")

		fmt.Println(filePath)

		if stat, err := os.Stat(filePath); err != nil || stat.IsDir() {
			fmt.Printf("could not find .3dnode file for %s, skipping\n", entity.Name)
			continue
		}

		objBuilder := modelextractor.NewOBJBuilder()

		modelextractor.Extract(filePath, objBuilder)

		entityFileName := strings.ReplaceAll(entityGCType, ".", "-")

		err := ioutil.WriteFile(path.Join(destPath, entityFileName+".obj"), []byte(objBuilder.String()), 0755)

		if err != nil {
			panic(err)
		}
	}
}

func getDRClassesForStaticObjects(dump []StaticObjectDump, config *configparser.DRConfig) []struct {
	string
	*database.DRClass
} {
	var allEntities []struct {
		string
		*database.DRClass
	}

	for _, dumpObject := range dump {
		fmt.Printf("processing %s\n", dumpObject.GCTypeName)

		gcObject, err := config.Get(dumpObject.GCTypeName)

		if err != nil {
			panic(err)
		}

		if len(gcObject) > 1 || len(gcObject[0].Entities) > 1 {
			panic("too many results")
		}

		entity := gcObject[0].Entities[0]

		entity.Name = gcObject[0].Name

		allEntities = append(allEntities, struct {
			string
			*database.DRClass
		}{dumpObject.GCTypeName, entity})
	}

	return allEntities
}
