package database

import (
	"RainbowRunner/cmd/rrcli/configurator"
	"RainbowRunner/internal/gosucks"
	"RainbowRunner/internal/types/configtypes"
	"fmt"
	log "github.com/sirupsen/logrus"
	"strings"
)

func LoadWorldConfigs() map[string]*configtypes.WorldConfig {
	log.Info("loading world configs")

	worlds := make(map[string]*configtypes.WorldConfig)

	worldsConfig, err := configurator.LoadFromDumpedConfigFile("resources/Dumps/generated/worlds.json")

	if err != nil {
		panic(err)
	}

	for worldID, worldGroup := range worldsConfig.Classes.Children {
		if len(worldGroup.Entities) > 1 {
			panic(fmt.Sprintf("world %s has more than one entity", worldID))
		}

		worldDef := worldGroup.Entities[0]

		worldConfig := configtypes.NewWorldConfig()
		props := worldDef.Properties

		setPropertiesOnStruct(worldConfig, props)

		if worldDef.Children != nil && worldDef.Children["entities"] != nil {
			loadWorldEntities(worldConfig, worldDef.Children["entities"].Entities[0])
		}

		worlds[worldID] = worldConfig
	}

	gosucks.VAR(worldsConfig)

	return worlds
}

func loadWorldEntities(worldConfig *configtypes.WorldConfig, entitiesConfig *configtypes.DRClass) {
	if entitiesConfig.Children != nil {
		for _, entityGroup := range entitiesConfig.Children {
			for _, entity := range entityGroup.Entities {
				entityConfig := configtypes.NewEntityConfig()
				props := entity.Properties

				entityConfig.Name = entity.Name

				types := config.GetInheritedTypes(entity)

				for _, t := range types {
					if t[0] == "npc" {
						entityConfig.Type = configtypes.EntityConfigTypeNPC
						break
					}

					if t[0] == "checkpointentity" {
						entityConfig.Type = configtypes.EntityConfigTypeCheckpoint
						break
					}

					if t[0] == "waypoint" {
						entityConfig.Type = configtypes.EntityConfigTypeWaypoint
						break
					}
				}

				//dumpTypes(types)

				gosucks.VAR(types)

				config.MergeParentsSingle(entity)

				setPropertiesOnStruct(entityConfig, props)

				worldConfig.Entities = append(worldConfig.Entities, entityConfig)
			}
		}
	}
}

func dumpTypes(types [][]string) {
	str := ""

	for i, t := range types {
		if i > 0 {
			str += " -> "
		}
		str += strings.Join(t, ".")
	}

	fmt.Println(str)
}
