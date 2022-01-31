package objects

import (
	byter "RainbowRunner/pkg/byter"
)

type AvatarMetrics struct {
	*Component
}

func (a AvatarMetrics) WriteFullGCObject(byter *byter.Byter) {
	a.GCObject.WriteFullGCObject(byter)

	// AvatarMetrics::PlayTime::readObject
	byter.WriteUInt32(0x01)
	byter.WriteUInt32(0x02)
	byter.WriteUInt32(0x03)
	byter.WriteUInt32(0x04)
	byter.WriteUInt32(0x05)

	// AvatarMetrics::ZoneToPlayTimeMap::readObject
	byter.WriteUInt32(0x00) // If > 0 it reads a string and goes to PlayTime::readObject

	//AvatarMetrics::LevelToPlayTimeMap::readObject
	byter.WriteUInt32(0x00) // If > 0 it reads a string and goes to PlayTime::readObject

	//AvatarMetrics::GoldStats::readObject
	byter.WriteUInt64(0x06)
	byter.WriteUInt64(0x07)
	byter.WriteUInt64(0x08)
	byter.WriteUInt64(0x09)
	byter.WriteUInt64(0x0a)

	//AvatarMetrics::LevelToGoldStatsMap::readObject
	byter.WriteUInt32(0x00) // If > 0 it reads a bunch more uint64

	//AvatarMetrics::SkillUseMap::readObject
	byter.WriteUInt32(0x00) // If > 0 it calls AvatarMetrics::ItemSnapshot::readObject

	//AvatarMetrics::DeathMap::readObject
	byter.WriteUInt32(0x00) // If > 0 it reads a string and some other values

	//AvatarMetrics::SkillUseMap::readObject
	byter.WriteUInt32(0x00) // If > 0 it calls AvatarMetrics::ItemSnapshot::readObject
}

func NewAvatarMetrics(id uint32, name string) *AvatarMetrics {
	component := NewComponent("avatarmetrics", "AvatarMetrics")
	component.GCLabel = name

	return &AvatarMetrics{
		Component: component,
	}
}
