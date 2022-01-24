package objects

import (
	"encoding/json"
	"os"
)

var ArmourMap map[string][]string

func Init() {
	file, err := os.Open("resources/Dumps/armor_dump.json")

	if err != nil {
		panic(err)
	}

	j := json.NewDecoder(file)

	ArmourMap = make(map[string][]string)

	err = j.Decode(&ArmourMap)

	if err != nil {
		panic(err)
	}
}
