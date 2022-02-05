package main

import (
	"RainbowRunner/cmd/rrcli/configurator"
	"RainbowRunner/internal/database"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
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

	var allEntities []*database.DRClass

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

		allEntities = append(allEntities, entity)
	}

	data, err = json.MarshalIndent(allEntities, "", "  ")

	if err != nil {
		panic(err)
	}

	err = ioutil.WriteFile("resources/Dumps/generated/staticworlddtounity.json", data, 0755)

	if err != nil {
		panic(err)
	}
}
