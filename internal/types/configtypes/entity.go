package configtypes

import (
	"RainbowRunner/internal/types/drconfigtypes"
	"RainbowRunner/pkg/datatypes"
	"RainbowRunner/pkg/datatypes/drfloat"
)

type EntityConfigType int

const (
	EntityConfigTypeUnknown EntityConfigType = iota
	EntityConfigTypeNPC
	EntityConfigTypeCheckpoint
	EntityConfigTypeWaypoint
)

//go:generate go run ../../../scripts/generatelua -type=EntityConfig
type EntityConfig struct {
	Name             string
	HitPoints        drfloat.DRFloat
	ManaPoints       drfloat.DRFloat
	Position         datatypes.Vector3Float32
	Heading          int
	Width            int
	Zone             string
	EncounterTable   string
	Height           int
	SpawnPoint       string
	SizeX            int
	SizeY            int
	SizeZ            int
	CanBeActivated   bool
	RespawnWhenClear bool
	Blocking         bool
	TableSelector    int
	Color            uint `parse:"hex"`
	ZoneStart        bool
	Level            int
	AutoRespawn      bool
	WorldEntityTable string
	RespawnRate      int

	Animations map[int]AnimationConfig
	Type       EntityConfigType
	FullGCType string
	Desc       *EntityDesc
	Behaviour  *BehaviourConfig
	Merchant   *MerchantConfig
}

type EntityDesc struct {
	ActivationClientSyncTolerance    int
	ActivationRange                  int
	AllowMultipleActivations         bool
	Animations                       string
	AttackRange                      int
	AttackRating                     float32
	Blocking                         bool
	BroadcastActivation              bool
	CanBeClosed                      bool
	Checkpoint                       string
	CollisionRadius                  int
	CorpseLingerTime                 int
	CreatureDifficulty               string
	CreatureElement                  string
	CreatureFamily                   string
	CriticalChance                   float32
	CrushingResist                   string
	DamageMod                        float32
	DefenseRating                    float32
	Description                      string
	Difficulty                       float32
	DivineResist                     float32
	DoorsToOpenOnDeath               string
	DynamicBlocking                  bool
	EncounterTable                   string
	ExpirationDate                   string // date time string e.g. "02/14/2009 0:00:01"
	FactionID                        int
	FearResist                       int
	FireResist                       float32
	GoodbyeSound                     string
	HelloSound                       string
	IceResist                        float32
	IsAlive                          bool
	IsOneHit                         bool
	ItemCount                        int
	ItemCount2                       int
	ItemCount3                       int
	ItemGenerator                    string
	ItemGenerator2                   string
	ItemGenerator3                   string
	Label                            string
	Locked                           bool
	LockedMessage                    string
	MaxActivationRange               float32
	MaxHealth                        float32
	MaxMana                          float32
	Name                             string
	OpenCount                        int
	OpenMessage                      string
	PartialOpenMessage               string
	PiercingResist                   float32
	PoisonResist                     float32
	QuestRequired                    bool
	Selectable                       bool
	SetInteractorAsSpellEffectSource bool
	ShadowResist                     float32
	ShowResistance                   bool
	SizeMod                          float32
	SlashingResist                   float32
	Sounds                           string
	SpawnAnimation                   int
	SpawnAnimationLength             int
	SpawnRate                        float32 // potentially DRFloat
	Speed                            float32
	SpeedMod                         float32 // potentially DRFloat
	SpellEffect                      string
	StunResist                       int
	TeleportPoint                    string
	TeleportZone                     string
	TrackActivatingPlayers           bool
	TreasureCount                    int
	TreasureCount2                   int
	TreasureGenerator                string
	TreasureGenerator2               string
	TriggerDelay                     float32
	TriggerOffset                    datatypes.Vector3Float32
	TriggerRadius                    float32
	TurnRate                         int
	UnlockObject                     string
	UseGeneratedName                 bool
	Visuals                          string
	WalkSpeed                        float32
}

func (e *EntityConfig) Init(entity *drconfigtypes.DRClass) {
	SetPropertiesOnStruct(e, entity.Properties)

	if entity.Children == nil {
		return
	}

	if description, ok := entity.Children["description"]; ok {
		e.Desc = &EntityDesc{}
		SetPropertiesOnStruct(e.Desc, description.Entities[0].Properties)
	}
}

func NewEntityConfig() *EntityConfig {
	return &EntityConfig{
		CanBeActivated: true,
	}
}
