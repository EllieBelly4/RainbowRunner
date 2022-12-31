package objects

import (
	"RainbowRunner/internal/connections"
	"RainbowRunner/internal/database"
	"RainbowRunner/internal/lua"
	"RainbowRunner/internal/pathfinding"
	script2 "RainbowRunner/internal/script"
	"RainbowRunner/internal/serverconfig"
	"RainbowRunner/internal/types"
	"RainbowRunner/internal/types/drobjecttypes"
	"RainbowRunner/pkg/byter"
	"RainbowRunner/pkg/datatypes"
	"fmt"
	log "github.com/sirupsen/logrus"
	"strings"
	"sync"
)

//go:generate go run ../../scripts/generatelua -type=Zone
type Zone struct {
	sync.RWMutex

	Name     string
	entities map[uint16]drobjecttypes.DRObject
	players  map[uint16]*RRPlayer

	Scripts *ZoneLuaScripts

	BaseConfig  *database.ZoneConfig
	PathMap     *types.PathMap
	ID          uint32
	initialised bool
}

func (z *Zone) Initialised() bool {
	return z.initialised
}

func (z *Zone) Entities() []drobjecttypes.DRObject {
	z.RLock()
	defer z.RUnlock()

	l := make([]drobjecttypes.DRObject, 0)

	for _, drObject := range z.entities {
		l = append(l, drObject)
	}

	return l
}

func (z *Zone) Players() []*RRPlayer {
	z.RLock()
	defer z.RUnlock()

	l := make([]*RRPlayer, 0)

	for _, player := range z.players {
		l = append(l, player)
	}

	return l
}

func (z *Zone) RemovePlayer(id int) {
	z.Lock()
	defer z.Unlock()

	delete(z.players, uint16(id))

	toDelete := make([]uint16, 0, 1024)

	for index, entity := range z.entities {
		if entity == nil || entity.(IRREntityPropertiesHaver).GetRREntityProperties().OwnerID == uint16(id) {
			toDelete = append(toDelete, index)
		}
	}

	for _, index := range toDelete {
		z.entities[index].(IRREntityPropertiesHaver).GetRREntityProperties().Zone = nil
		delete(z.entities, index)
	}
}

func (z *Zone) SpawnEntity(owner *uint16, entity drobjecttypes.DRObject) {
	z.Lock()
	defer z.Unlock()

	z.setZone(entity)
	z.GiveID(entity)

	if owner != nil {
		entity.(IRREntityPropertiesHaver).GetRREntityProperties().SetOwner(*owner)
	}

	entity.WalkChildren(func(object drobjecttypes.DRObject) {
		z.GiveID(object)
		z.setZone(object)

		if owner != nil {
			object.(IRREntityPropertiesHaver).GetRREntityProperties().SetOwner(*owner)
		}
	})

	id := uint16(entity.(IRREntityPropertiesHaver).GetRREntityProperties().ID)

	if _, ok := z.entities[id]; ok {
		return
	}

	z.entities[id] = entity

	entity.Init()
}

func (z *Zone) AddPlayer(player *RRPlayer) {
	z.Lock()
	z.players[uint16(player.Conn.GetID())] = player
	z.Unlock()

	for _, child := range player.CurrentCharacter.Children() {
		z.SpawnEntity(types.UInt16(uint16(player.Conn.GetID())), child)
	}
}

func (z *Zone) setZone(entities ...drobjecttypes.DRObject) {
	for _, entity := range entities {
		entity.(IRREntityPropertiesHaver).GetRREntityProperties().Zone = z
		z.setZone(entity.Children()...)
	}
}

func (z *Zone) SendToAll(body *byter.Byter) {
	z.RLock()
	defer z.RUnlock()

	for _, player := range z.players {
		connections.WriteCompressedASimple(player.Conn, body)
	}
}

func (z *Zone) SpawnEntityWithPosition(
	entity drobjecttypes.DRObject,
	position datatypes.Vector3Float32,
	rotation float32,
	ownerID *uint16,
) {
	if _, ok := entity.(IWorldEntity); ok {
		worldEntity := entity.(IWorldEntity).GetWorldEntity()

		worldEntity.WorldPosition = position
		worldEntity.Heading = rotation
	}

	if unitBehavior, ok := entity.GetChildByGCNativeType("UnitBehavior").(IUnitBehavior); unitBehavior != nil && ok {
		behavior := unitBehavior.GetUnitBehavior()

		behavior.Position = position
		behavior.Rotation = rotation
	}

	z.SpawnEntity(ownerID, entity)
}

// Spawn
// Deprecated: use SpawnEntityWithPosition
func (z *Zone) Spawn(entity drobjecttypes.DRObject, position datatypes.Vector3Float32, rotation float32) {
	z.SpawnEntityWithPosition(entity, position, rotation, nil)
}

func (z *Zone) GetEntityScript(id string) script2.IEntityScript {
	script := lua.GetScript("zones." + strings.ToLower(z.Name) + "." + strings.ToLower(id))

	if script == nil {
		return nil
	}

	return script2.NewEntityScript(script, z.Scripts.State)
}

func (z *Zone) LoadWaypointFromConfig(id string) *Waypoint {
	waypointConfig, ok := z.BaseConfig.Waypoints[strings.ToLower(id)]

	if !ok {
		log.Errorf("waypoint '%s' not found in zone '%s'", id, z.Name)
		return nil
	}

	waypoint := NewWaypointFromConfig(waypointConfig.GetWaypointConfig())

	return initialiseWorldEntity[*Waypoint](z, waypoint, id)
}

func (z *Zone) LoadCheckpointEntityFromConfig(id string) *CheckpointEntity {
	checkpointEntityConfig, ok := z.BaseConfig.Checkpoints[strings.ToLower(id)]

	if !ok {
		log.Errorf("checkpointEntity '%s' not found in zone '%s'", id, z.Name)
		return nil
	}

	checkpointEntity := NewCheckpointEntityFromConfig(checkpointEntityConfig.GetCheckpointEntityConfig())

	return initialiseWorldEntity[*CheckpointEntity](z, checkpointEntity, id)
}

func (z *Zone) LoadNPCFromConfig(id string) *NPC {
	npcConfig, ok := z.BaseConfig.NPCs[strings.ToLower(id)]

	if !ok {
		log.Errorf("npc '%s' not found in zone '%s'", id, z.Name)
		return nil
	}

	npc := NewNPCFromConfig(npcConfig.GetNPCConfig())

	return initialiseWorldEntity[*NPC](z, npc, id)
}

func initialiseWorldEntity[T IWorldEntity](zone *Zone, entity IWorldEntity, id string) T {
	scriptPrefix := "entity"

	switch entity.(type) {
	case ICheckpointEntity:
		scriptPrefix = "checkpointentity"
	case IWaypoint:
		scriptPrefix = "waypoint"
	case INPC:
		scriptPrefix = "npc"
	}

	script := zone.GetEntityScript(scriptPrefix + "." + strings.ToLower(id))

	if script != nil {
		entity.GetWorldEntity().SetScript(script)
	}

	return entity.(T)
}

func (z *Zone) Init() {
	config, err := database.GetZoneConfig(strings.ToLower(z.Name))

	if err != nil {
		panic(err)
	}

	z.BaseConfig = config

	z.ReloadPathMap()
	z.initLua()

	z.initialised = true
}

func (z *Zone) initLua() {
	log.Infof("initialising zone %s", z.Name)

	z.Scripts = NewZoneLuaScripts(z)

	err := z.Scripts.Load()

	if err != nil {
		panic(err)
	}

	err = z.Scripts.Init(nil)

	if err != nil {
		log.Errorf("failed to execute zone init script %s: %s", z.Name, err.Error())
	}
}

func (z *Zone) ClearEntities() {
	z.Lock()
	defer z.Unlock()

	z.entities = make(map[uint16]drobjecttypes.DRObject)
}

func (z *Zone) ReloadPathMap() {
	z.PathMap = pathfinding.ReloadPathMap(z.Name)
}

func (z *Zone) Tick() error {
	es := z.Entities()

	for _, entity := range es {
		if entity == nil {
			continue
		}
		entity.Tick()
	}

	err := z.Scripts.Tick()

	return err
}

func (z *Zone) FindEntityByGCTypeName(name string) drobjecttypes.DRObject {
	for _, entity := range z.Entities() {
		if entity == nil {
			continue
		}

		gcType := entity.(IGCObject).GetGCObject().GCType
		if strings.ToLower(gcType) == strings.ToLower(name) {
			return entity
		}
	}

	return nil
}

func (z *Zone) FindEntityByID(id uint16) drobjecttypes.DRObject {
	z.RLock()
	defer z.RUnlock()
	for _, entity := range z.entities {
		if entity.(IRREntityPropertiesHaver).GetRREntityProperties().ID == uint32(id) {
			return entity
		}

		var foundEntity drobjecttypes.DRObject = nil

		entity.WalkChildren(func(object drobjecttypes.DRObject) {
			// TODO optimise this, no need to loop all children when found
			if object.(IRREntityPropertiesHaver).GetRREntityProperties().ID == uint32(id) {
				foundEntity = object
			}
		})

		if foundEntity != nil {
			return foundEntity
		}
	}
	return nil
}

func (z *Zone) GiveID(entity drobjecttypes.DRObject) {
	eProps := entity.(IRREntityPropertiesHaver).GetRREntityProperties()

	if eProps.ID == 0 {
		eProps.ID = uint32(NewID())
	}

	if serverconfig.Config.Logging.LogIDs {
		fmt.Printf("%d - %s(%s)\n", eProps.ID, entity.(IGCObject).GetGCObject().GCType, entity.(IGCObject).GetGCObject().GCLabel)
	}
}

func NewZone(name string, id uint32) *Zone {
	zone := &Zone{
		Name:     name,
		ID:       id,
		entities: make(map[uint16]drobjecttypes.DRObject),
		players:  make(map[uint16]*RRPlayer),
	}

	return zone
}
