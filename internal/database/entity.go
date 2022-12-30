package database

import (
	"RainbowRunner/internal/gosucks"
	"RainbowRunner/internal/types/configtypes"
)

func newEntityConfigFromRawConfig(entity *configtypes.DRClass) *configtypes.EntityConfig {
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
	return entityConfig
}

func addEntityBehaviour(entityConfig *configtypes.EntityConfig, entity *configtypes.DRClass) {
	if behaviour, ok := entity.Children["behavior"]; ok && entityConfig.Type == configtypes.EntityConfigTypeNPC {
		behavGCType := "npc.Base.Behavior"
		customBehaviourType := entityConfig.FullGCType + ".behavior"

		_, err := config.Get(customBehaviourType)

		if err == nil {
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
