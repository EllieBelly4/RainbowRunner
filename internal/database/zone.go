package database

import (
	"RainbowRunner/internal/types/configtypes"
	"strings"
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
	gcRoot := []string{"world", name}

	rawConfig, err := config.Get(strings.Join(gcRoot, "."))

	if err != nil {
		return nil, err
	}

	zoneConfig := NewZoneConfig(name)

	configEntities := rawConfig[0].Entities[0].Children

	if npcConfig, ok := configEntities["npc"]; ok {
		handleNPCs(zoneConfig, npcConfig, append(gcRoot, "npc"))
	}

	if checkConfig, ok := checkpointConfigs[zoneConfig.Name]; ok {
		zoneConfig.Checkpoints = checkConfig
	}

	return zoneConfig, nil
}

func handleNPCs(zoneConfig *ZoneConfig, rawNPCConfig *configtypes.DRClassChildGroup, gcRoot []string) {
	zoneConfig.NPCs = make(map[string]*NPCConfig)

	for npcKey, npcConfig := range rawNPCConfig.Entities[0].Children {
		npcGCroot := append(gcRoot, npcKey)
		npc := NewNPCConfig(npcConfig, npcGCroot)

		npc.FullGCType = strings.Join(npcGCroot, ".")

		zoneConfig.NPCs[npcKey] = npc
	}
}

func NewZoneConfig(name string) *ZoneConfig {
	return &ZoneConfig{
		Name: name,
	}
}
