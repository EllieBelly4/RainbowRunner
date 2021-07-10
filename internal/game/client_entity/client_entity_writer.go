package client_entity

import (
	"RainbowRunner/internal/game/messages"
	"RainbowRunner/internal/objects"
	byter "RainbowRunner/pkg/byter"
)

//body.WriteByte(0x08) // CreateInit
//body.WriteByte(0x02) // Init
//body.WriteByte(0x03) // Update
//body.WriteByte(21) // ClearEntityManager

type GCObjectLookupType byte

const (
	GCObjectLookupTypeString GCObjectLookupType = 0xFF
	GCObjectLookupTypeHash   GCObjectLookupType = 0x04
)

type ClientEntityWriter struct {
	body *byter.Byter
}

func (w *ClientEntityWriter) Start() {
	w.body.WriteByte(byte(messages.ClientEntityChannel))
}

func (w *ClientEntityWriter) Create(object objects.DRObject) {
	w.body.WriteByte(0x01) // Create

	w.body.WriteUInt16(object.RREntityProperties().ID) // Entity ID
	w.body.WriteByte(byte(GCObjectLookupTypeString))
	w.body.WriteCString(object.GetGCObject().GCType)
}

func (w *ClientEntityWriter) Init(object objects.DRObject) {
	w.body.WriteByte(0x02) // Init
	w.body.WriteUInt16(object.RREntityProperties().ID)

	object.WriteInit(w.body)
}

func (w *ClientEntityWriter) CreateComponent(component objects.DRObject, targetEntity objects.DRObject) {
	w.body.WriteByte(0x32)                                   // Create Component
	w.body.WriteUInt16(targetEntity.RREntityProperties().ID) // Parent Entity ID
	w.body.WriteUInt16(component.RREntityProperties().ID)    // Component ID
	w.body.WriteByte(byte(GCObjectLookupTypeString))
	w.body.WriteCString(component.GetGCObject().GCType) // Component Type
	w.body.WriteByte(0x01)                              // Unk

	component.WriteInit(w.body)
}

func (w *ClientEntityWriter) Update(player objects.DRObject) {
	w.body.WriteByte(0x03)                             // MsgType Update
	w.body.WriteUInt16(player.RREntityProperties().ID) // Entity ID

	player.WriteUpdate(w.body)

	// EntitySynchInfo
	// Flags
	w.body.WriteByte(0x0)
}

func NewClientEntityWriter(b *byter.Byter) *ClientEntityWriter {
	return &ClientEntityWriter{
		body: b,
	}
}
