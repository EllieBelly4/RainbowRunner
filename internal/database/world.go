package database

import (
	"RainbowRunner/cmd/rrcli/configurator"
	"RainbowRunner/internal/gosucks"
	"RainbowRunner/internal/types/configtypes"
	"fmt"
	log "github.com/sirupsen/logrus"
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

				config.MergeParentsSingle(entity)

				setPropertiesOnStruct(entityConfig, props)

				worldConfig.Entities = append(worldConfig.Entities, entityConfig)
			}
		}
	}
}
