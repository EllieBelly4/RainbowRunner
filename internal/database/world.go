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

		configtypes.SetPropertiesOnStruct(worldConfig, props)

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

				//props := entity.Properties

				entityConfig.Name = entity.Extends
				entityConfig.FullGCType = entity.Extends

				//dumpTypes(types)

				gosucks.VAR(types)

				config.MergeParentsSingle(entity)
				entityConfig.Init(entity)

				addEntityBehaviour(entityConfig, entity)

				if entityConfig.Desc != nil && entityConfig.Desc.Animations != "" {
					entityConfig.Animations = loadAnimations(entityConfig.Desc.Animations)
				}

				worldConfig.Entities = append(worldConfig.Entities, entityConfig)
			}
		}
	}
}

func addEntityBehaviour(entityConfig *configtypes.EntityConfig, entity *configtypes.DRClass) {
	if behaviour, ok := entity.Children["behavior"]; ok && entityConfig.Type == configtypes.EntityConfigTypeNPC {
		behavGCType := "npc.Base.Behavior"
		customBehaviourType := entityConfig.FullGCType + ".behavior"

		_, err := config.Get(customBehaviourType)

		if entity.Extends == "" && err != nil {
			behavGCType = customBehaviourType
		}

		entityConfig.Behaviour = &configtypes.BehaviourConfig{
			Type: behavGCType,
		}

		behavEntity := behaviour.Entities[0]
		configtypes.SetPropertiesOnStruct(entityConfig.Behaviour, behavEntity.Properties)

		entityConfig.Behaviour.Init(behavEntity, entityConfig.FullGCType)
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
