package objects

import (
	"RainbowRunner/internal/game/messages"
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
	Body *byter.Byter
}

func (w *ClientEntityWriter) BeginStream() {
	w.Body.WriteByte(byte(messages.ClientEntityChannel))
}

func (w *ClientEntityWriter) Create(object DRObject) {
	w.Body.WriteByte(0x01) // Create

	w.Body.WriteUInt16(object.RREntityProperties().ID) // Entity ID
	w.Body.WriteByte(byte(GCObjectLookupTypeString))
	w.Body.WriteCString(object.GetGCObject().GCType)
}

func (w *ClientEntityWriter) Init(object DRObject) {
	w.Body.WriteByte(0x02) // Init
	w.Body.WriteUInt16(object.RREntityProperties().ID)

	object.WriteInit(w.Body)
}

func (w *ClientEntityWriter) CreateComponent(component DRObject, targetEntity DRObject) {
	w.Body.WriteByte(0x32)                                   // Create Component
	w.Body.WriteUInt16(targetEntity.RREntityProperties().ID) // Parent Entity ID
	w.Body.WriteUInt16(component.RREntityProperties().ID)    // Component ID
	w.Body.WriteByte(byte(GCObjectLookupTypeString))
	w.Body.WriteCString(component.GetGCObject().GCType) // Component Type
	w.Body.WriteByte(0x01)                              // Unk

	component.WriteInit(w.Body)
}

func (w *ClientEntityWriter) Update(object DRObject) {
	w.Body.WriteByte(0x03)                             // MsgType Update
	w.Body.WriteUInt16(object.RREntityProperties().ID) // Entity ID

	object.WriteUpdate(w.Body)

	// EntitySynchInfo
	// Flags
	w.Body.WriteByte(0x0)
}

func (w *ClientEntityWriter) BeginComponentUpdate(id uint16) {
	w.Body.WriteByte(0x35) // ComponentUpdate
	w.Body.WriteUInt16(id) // ComponentID
}

func (w *ClientEntityWriter) EndStream() {
	w.Body.WriteByte(0x06)
}

func (w *ClientEntityWriter) WriteSynch(object DRObject) {
	object.WriteSynch(w.Body)
}

func NewClientEntityWriter(b *byter.Byter) *ClientEntityWriter {
	return &ClientEntityWriter{
		Body: b,
	}
}
func NewClientEntityWriterWithByter() *ClientEntityWriter {
	return NewClientEntityWriter(byter.NewLEByter(make([]byte, 0, 256)))
}
