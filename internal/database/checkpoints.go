package database

import (
	"RainbowRunner/internal/gosucks"
	"RainbowRunner/internal/types/drconfigtypes"
	"strconv"
	"strings"
)

func sortCheckpoints(rawCheckpointConfigs []*drconfigtypes.DRClassChildGroup) map[string]map[string]*CheckpointConfig {
	tmpCheckpoints := make(map[string]map[string]*CheckpointConfig)
	zoneMap := make(map[string]string)

	// Process checkpoints first
	for checkpointKey, checkpointConfig := range rawCheckpointConfigs[0].Entities[0].Children {
		entity := checkpointConfig.Entities[0]

		if entity.Extends != "base.checkpoint" {
			continue
		}

		descProperties := entity.Children["description"].Entities[0].Properties
		targetZone, ok := descProperties["Zone"]

		if !ok {
			continue
		}

		targetZone = strings.ToLower(targetZone)

		if _, ok := tmpCheckpoints[targetZone]; !ok {
			tmpCheckpoints[targetZone] = make(map[string]*CheckpointConfig)
		}

		zoneConfig := tmpCheckpoints[targetZone]

		fullGCType := "world.checkpoints." + strings.ToLower(checkpointKey)
		zoneMap[fullGCType] = targetZone

		if _, ok := zoneConfig[fullGCType]; !ok {
			zoneConfig[fullGCType] = &CheckpointConfig{}
		}

		checkConfig := zoneConfig[fullGCType]

		order, _ := strconv.ParseInt(descProperties["Order"], 10, 32)

		checkConfig.Order = int32(order)
		checkConfig.Name = checkpointKey
		checkConfig.Zone = descProperties["Zone"]
		checkConfig.FullGCType = fullGCType
		checkConfig.SpawnPoint = descProperties["SpawnPoint"]

		gosucks.VAR(checkpointKey)
	}

	// Process checkpointentities after we definitely have all of the target checkpoints available
	for checkpointKey, checkpointConfig := range rawCheckpointConfigs[0].Entities[0].Children {
		entity := checkpointConfig.Entities[0]

		if entity.Extends != "base.checkpointentity" {
			continue
		}

		properties := entity.Properties
		descProperties := entity.Children["description"].Entities[0].Properties
		targetCheckpoint := strings.ToLower(descProperties["Checkpoint"])
		targetZone, ok := zoneMap[targetCheckpoint]

		if !ok {
			continue
		}

		targetZone = strings.ToLower(targetZone)
		zoneConfig := tmpCheckpoints[targetZone]
		checkConfig := zoneConfig[targetCheckpoint]

		fullGCType := "world.checkpoints." + strings.ToLower(checkpointKey)

		blocking, _ := strconv.ParseBool(properties["Blocking"])

		checkConfig.Entity = &CheckpointEntityConfig{
			FullGCType: fullGCType,
			Label:      descProperties["Label"],
			Blocking:   blocking,
		}

		gosucks.VAR(checkpointKey)
	}

	return tmpCheckpoints
}
