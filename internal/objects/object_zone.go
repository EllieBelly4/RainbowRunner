package objects

import (
	"RainbowRunner/internal/helpers"
	"RainbowRunner/internal/lua"
	"RainbowRunner/pkg/byter"
	lua2 "github.com/yuin/gopher-lua"
	"sync"
)

type Zone struct {
	sync.RWMutex
	Name     string
	entities map[uint16]DRObject
	players  map[uint16]*RRPlayer
	Scripts  *lua.LuaScriptGroup
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
		if entity == nil || entity.RREntityProperties().OwnerID == id {
			toDelete = append(toDelete, index)
		}
	}

	for _, index := range toDelete {
		z.entities[index].RREntityProperties().Zone = nil
		delete(z.entities, index)
	}
}

func (z *Zone) AddEntity(entity DRObject) {
	z.setZone(entity)

	z.entities[uint16(entity.RREntityProperties().ID)] = entity
}

func (z *Zone) AddPlayer(player *RRPlayer) {
	z.players[uint16(player.Conn.GetID())] = player
}

func (z *Zone) Spawn(entity DRObject) {
	Entities.RegisterAll(nil, entity)
	z.AddEntity(entity)
}

func (z *Zone) setZone(entities ...DRObject) {
	for _, entity := range entities {
		entity.RREntityProperties().Zone = z
		z.setZone(entity.Children()...)
	}
}

func (z *Zone) SendToAll(body *byter.Byter) {
	for _, player := range z.players {
		helpers.WriteCompressedASimple(player.Conn, body)
	}
}

func (z *Zone) Init() {
	script := z.Scripts.Get("init")

	if script == nil {
		return
	}

	state := lua2.NewState()
	defer state.Close()

	//zoneConfig := database.
	AddZoneToState(state, z)

	script.Execute(state)
}
