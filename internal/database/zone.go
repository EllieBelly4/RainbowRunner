package database

import (
	"RainbowRunner/internal/types/configtypes"
)

type CheckpointConfig struct {
	Name       string
	FullGCType string
	Order      int32
	SpawnPoint string
	Zone       string
	Entity     *CheckpointEntityConfig
}

type CheckpointEntityConfig struct {
	FullGCType string
	Label      string
	Blocking   bool
}

type ZoneConfig struct {
	Name        string
	NPCs        map[string]*NPCConfig
	Checkpoints map[string]*CheckpointConfig
}

func GetZoneConfig(name string) (*ZoneConfig, error) {
	rawConfig, err := config.Get("world." + name)

	if err != nil {
		return nil, err
	}

	zoneConfig := NewZoneConfig(name)

	configEntities := rawConfig[0].Entities[0].Children

	if npcConfig, ok := configEntities["npc"]; ok {
		handleNPCs(zoneConfig, npcConfig)
	}

	if checkConfig, ok := checkpointConfigs[zoneConfig.Name]; ok {
		zoneConfig.Checkpoints = checkConfig
	}

	return zoneConfig, nil
}

func handleNPCs(zoneConfig *ZoneConfig, rawNPCConfig *configtypes.DRClassChildGroup) {
	zoneConfig.NPCs = make(map[string]*NPCConfig)

	for npcKey, npcConfig := range rawNPCConfig.Entities[0].Children {
		npc := NewNPCConfig(npcConfig)

		npc.FullGCType = "world." + zoneConfig.Name + ".npc." + npcKey

		zoneConfig.NPCs[npcKey] = npc
	}
}

func NewZoneConfig(name string) *ZoneConfig {
	return &ZoneConfig{
		Name: name,
	}
}
