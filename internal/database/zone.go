package database

import (
	"RainbowRunner/internal/types/configtypes"
)

type ZoneConfig struct {
	Name string
	NPCs map[string]*NPCConfig
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
