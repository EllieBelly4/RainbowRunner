package objects

import (
	"RainbowRunner/internal/config"
	"RainbowRunner/internal/connections"
	"RainbowRunner/internal/database"
	"RainbowRunner/internal/lua"
	"RainbowRunner/internal/pathfinding"
	"RainbowRunner/internal/types"
	"RainbowRunner/pkg/byter"
	"RainbowRunner/pkg/datatypes"
	"fmt"
	log "github.com/sirupsen/logrus"
	lua2 "github.com/yuin/gopher-lua"
	"strings"
	"sync"
)

//go:generate go run ../../scripts/generatelua -type=Zone
type Zone struct {
	sync.RWMutex
	Name       string
	entities   map[uint16]DRObject
	players    map[uint16]*RRPlayer
	Scripts    *lua.LuaScriptGroup
	BaseConfig *database.ZoneConfig
	PathMap    *types.PathMap
	ID         uint32
}

func (z *Zone) Entities() []DRObject {
	z.RLock()
	defer z.RUnlock()

	l := make([]DRObject, 0)

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
		if entity == nil || entity.RREntityProperties().OwnerID == uint16(id) {
			toDelete = append(toDelete, index)
		}
	}

	for _, index := range toDelete {
		z.entities[index].RREntityProperties().Zone = nil
		delete(z.entities, index)
	}
}

func (z *Zone) AddEntity(owner *uint16, entity DRObject) {
	z.Lock()
	defer z.Unlock()

	z.setZone(entity)
	z.GiveID(entity)

	if owner != nil {
		entity.RREntityProperties().SetOwner(*owner)
	}

	entity.WalkChildren(func(object DRObject) {
		z.GiveID(object)
		z.setZone(object)

		if owner != nil {
			object.RREntityProperties().SetOwner(*owner)
		}
	})

	id := uint16(entity.RREntityProperties().ID)

	if _, ok := z.entities[id]; ok {
		return
	}

	z.entities[id] = entity
}

func (z *Zone) AddPlayer(player *RRPlayer) {
	z.Lock()
	z.players[uint16(player.Conn.GetID())] = player
	z.Unlock()

	for _, child := range player.CurrentCharacter.Children() {
		z.AddEntity(types.UInt16(uint16(player.Conn.GetID())), child)
	}
}

func (z *Zone) setZone(entities ...DRObject) {
	for _, entity := range entities {
		entity.RREntityProperties().Zone = z
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

func (z *Zone) Spawn(entity DRObject, position datatypes.Vector3Float32, rotation float32) {
	if _, ok := entity.(IWorldEntity); ok {
		worldEntity := entity.(IWorldEntity).GetWorldEntity()

		worldEntity.WorldPosition = position
		worldEntity.Rotation = rotation
	}

	if unitBehavior, ok := entity.GetChildByGCNativeType("UnitBehavior").(IUnitBehavior); unitBehavior != nil && ok {
		behavior := unitBehavior.GetUnitBehavior()

		behavior.Position = position
		behavior.Rotation = rotation
	}

	z.AddEntity(nil, entity)
}

func (z *Zone) LoadNPCFromConfig(id string) *NPC {
	npcConfig, ok := z.BaseConfig.NPCs[strings.ToLower(id)]

	if !ok {
		log.Errorf("npc '%s' not found in zone '%s'", id, z.Name)
		return nil
	}

	npc := NewNPCFromConfig(npcConfig)
	return npc
}

func (z *Zone) Init() {
	z.ReloadPathMap()

	z.Scripts = lua.GetScriptGroup("zones." + strings.ToLower(z.Name))

	log.Infof("initialising zone %s", z.Name)

	config, err := database.GetZoneConfig(strings.ToLower(z.Name))

	if err != nil {
		panic(err)
	}

	z.BaseConfig = config

	script := z.Scripts.Get("init")

	if script == nil {
		return
	}

	state := lua2.NewState()
	defer state.Close()

	RegisterLuaGlobals(state)
	//func AddZoneToState(L *lua.LState, z *Zone) {
	//	ud := L.NewUserData()
	//	ud.Value = z
	//	L.SetMetatable(ud, L.GetTypeMetatable(luaZoneTypeName))
	//	L.SetGlobal("currentZone", ud)
	//}

	state.SetGlobal("currentZone", z.ToLua(state))

	err = script.Execute(state)

	if err != nil {
		log.Errorf("failed to execute zone init script %s: %s", z.Name, err.Error())
	}
}

func (z *Zone) ClearEntities() {
	z.Lock()
	defer z.Unlock()

	z.entities = make(map[uint16]DRObject)
}

func (z *Zone) ReloadPathMap() {
	z.PathMap = pathfinding.ReloadPathMap(z.Name)
}

func (z *Zone) Tick() {
	es := z.Entities()

	for _, entity := range es {
		if entity == nil {
			continue
		}
		entity.Tick()
	}
}

func (z *Zone) FindEntityByGCTypeName(name string) DRObject {
	for _, entity := range z.Entities() {
		if entity == nil {
			continue
		}

		gcType := entity.GetGCObject().GCType
		if strings.ToLower(gcType) == strings.ToLower(name) {
			return entity
		}
	}

	return nil
}

func (z *Zone) FindEntityByID(id uint16) DRObject {
	z.RLock()
	defer z.RUnlock()
	for _, entity := range z.entities {
		if entity.RREntityProperties().ID == uint32(id) {
			return entity
		}

		var foundEntity DRObject = nil

		entity.WalkChildren(func(object DRObject) {
			// TODO optimise this, no need to loop all children when found
			if object.RREntityProperties().ID == uint32(id) {
				foundEntity = object
			}
		})

		if foundEntity != nil {
			return foundEntity
		}
	}
	return nil
}

func (z *Zone) GiveID(entity DRObject) {
	eProps := entity.RREntityProperties()

	if eProps.ID == 0 {
		eProps.ID = uint32(NewID())
	}

	if config.Config.Logging.LogIDs {
		fmt.Printf("%d - %s(%s)\n", eProps.ID, entity.GetGCObject().GCType, entity.GetGCObject().GCLabel)
	}
}

func NewZone(name string, id uint32) *Zone {
	zone := &Zone{
		Name:     name,
		ID:       id,
		entities: make(map[uint16]DRObject),
		players:  make(map[uint16]*RRPlayer),
	}

	return zone
}
