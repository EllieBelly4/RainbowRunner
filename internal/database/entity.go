package database

import (
	"RainbowRunner/internal/gosucks"
	"RainbowRunner/internal/types/configtypes"
	"RainbowRunner/internal/types/drconfigtypes"
)

func newEntityConfigFromRawConfig(entity *drconfigtypes.DRClass) configtypes.IEntityConfig {
	var ientityConfig configtypes.IEntityConfig = nil

	types := config.GetInheritedTypes(entity)

	for _, t := range types {
		if t[0] == "npc" {
			ientityConfig = configtypes.NewNPCConfig()
			break
		}

		if t[0] == "checkpointentity" {
			ientityConfig = configtypes.NewEntityConfig(configtypes.EntityConfigTypeCheckpoint)
			break
		}

		if t[0] == "waypoint" {
			ientityConfig = configtypes.NewEntityConfig(configtypes.EntityConfigTypeWaypoint)
			break
		}
	}

	if ientityConfig == nil {
		ientityConfig = configtypes.NewEntityConfig(configtypes.EntityConfigTypeUnknown)
	}

	entityConfig := ientityConfig.GetEntityConfig()

	//props := entity.Properties

	entityConfig.Name = entity.Extends
	entityConfig.FullGCType = entity.Extends

	//dumpTypes(types)

	gosucks.VAR(types)

	config.MergeParentsSingle(entity)
	entityConfig.Init(entity)

	addEntityBehaviour(entityConfig, entity)

	if entity.Children["merchant"] != nil {
		entityConfig.Merchant = newMerchantConfigFromRawConfig(entity.Children["merchant"].Entities[0])
	}

	if entityConfig.Desc != nil && entityConfig.Desc.Animations != "" {
		entityConfig.Animations = loadAnimations(entityConfig.Desc.Animations)
	}

	return ientityConfig
}

func newMerchantConfigFromRawConfig(merchantConfig *drconfigtypes.DRClass) *configtypes.MerchantConfig {
	merchant := configtypes.NewMerchantConfig(merchantConfig.GCType)

	configtypes.SetPropertiesOnStruct(merchant, merchantConfig.Properties)

	for name, classGroup := range merchantConfig.Children {
		childEntity := classGroup.Entities[0]

		if childEntity.Extends == "MerchantInventory" {
			merchant.AddInventory(name, newMerchantInventoryConfigFromRawConfig(childEntity))
		}
	}

	return merchant
}

func newMerchantInventoryConfigFromRawConfig(entity *drconfigtypes.DRClass) *configtypes.MerchantInventoryConfig {
	inventory := configtypes.NewMerchantInventoryConfig(entity.GCType)

	configtypes.SetPropertiesOnStruct(inventory, entity.Properties)

	if desc, ok := entity.Children["description"]; ok {
		inventory.Description = configtypes.NewInventoryDescConfig()
		configtypes.SetPropertiesOnStruct(inventory.Description, desc.Entities[0].Properties)
	}

	return inventory
}

func addEntityBehaviour(entityConfig *configtypes.EntityConfig, entity *drconfigtypes.DRClass) {
	if behaviour, ok := entity.Children["behavior"]; ok && entityConfig.Type == configtypes.EntityConfigTypeNPC {
		entityConfig.Behaviour = &configtypes.BehaviourConfig{
			Type: behaviour.GCType,
		}

		behavEntity := behaviour.Entities[0]
		configtypes.SetPropertiesOnStruct(entityConfig.Behaviour, behavEntity.Properties)

		entityConfig.Behaviour.Init(behavEntity, entityConfig.FullGCType)
	}
}
