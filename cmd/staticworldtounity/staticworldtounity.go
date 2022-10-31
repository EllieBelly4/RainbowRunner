package main

import (
	"RainbowRunner/cmd/rrcli/configurator"
	"RainbowRunner/cmd/rrcli/modelextractor"
	"RainbowRunner/internal/gosucks"
	"RainbowRunner/internal/types/configtypes"
	"RainbowRunner/pkg/datatypes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
	"path"
	"path/filepath"
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
	//destPath := "E:\\Unity\\RainbowRunnerGameBackend\\Assets\\Resources\\ModelDumps\\Townston"
	destPath := "D:\\Work\\dungeon-runners\\models"

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

	extractModelsForObjects(allEntities, rootPath, destPath, false)
}

func extractModelsForObjects(entities []struct {
	string
	*configtypes.DRClass
}, rootPath string, destPath string, split bool) {
	useCollisionModel := false

	entityFileName := "static_town"

	objBuilder := modelextractor.NewOBJBuilder()
	mtlBuilder := modelextractor.NewMTLBuilder()

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

		if split {
			modelextractor.Split(filePath, filepath.Join(destPath, strings.ReplaceAll(entityGCType, ".", "-")))
		} else {
			dumpedPos := entity.CustomProperties["DumpedPosition"].(StaticObjectDumpLocation)
			objBuilder.SetModelOffset(datatypes.Vector3Float32{
				X: dumpedPos.X,
				Y: dumpedPos.Y,
				Z: dumpedPos.Z,
			}.DivideByFloat32(10))
			modelextractor.Extract(filePath, objBuilder, mtlBuilder)
		}

		gosucks.VAR(entityGCType, entityFileName, mtlBuilder, objBuilder)
	}

	if !split {
		err := ioutil.WriteFile(path.Join(destPath, entityFileName+".obj"), []byte(objBuilder.String()), 0755)

		if err != nil {
			panic(err)
		}
	}
}

func getDRClassesForStaticObjects(dump []StaticObjectDump, config *configtypes.DRConfig) []struct {
	string
	*configtypes.DRClass
} {
	var allEntities []struct {
		string
		*configtypes.DRClass
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

		if entity.CustomProperties == nil {
			entity.CustomProperties = make(map[string]interface{})
		}

		entity.CustomProperties["DumpedPosition"] = dumpObject.Position

		allEntities = append(allEntities, struct {
			string
			*configtypes.DRClass
		}{dumpObject.GCTypeName, entity})
	}

	return allEntities
}
