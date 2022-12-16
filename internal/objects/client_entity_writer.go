package objects

import (
	"RainbowRunner/internal/actions"
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
	Body      *byter.Byter
	dirty     bool
	sessionID byte
}

func (w *ClientEntityWriter) BeginStream() {
	w.Body.WriteByte(byte(messages.ClientEntityChannel))
}

func (w *ClientEntityWriter) Create(object DRObject) {
	w.dirty = true

	w.Body.WriteByte(0x01) // Create

	w.Body.WriteUInt16(uint16(object.RREntityProperties().ID)) // Entity ID
	w.Body.WriteByte(byte(GCObjectLookupTypeString))
	w.Body.WriteCString(object.GetGCObject().GCType)

	//TODO get this to work, it's a good idea I think but it causes errors
	//for _, child := range object.Children() {
	//	if child.Type() == DRObjectComponent {
	//		fmt.Printf(child.GetGCObject().GCType + "\n")
	//		w.CreateComponent(child, object)
	//	}
	//}
}

func (w *ClientEntityWriter) Init(object DRObject) {
	w.dirty = true

	w.Body.WriteByte(0x02) // Init
	w.Body.WriteUInt16(uint16(object.RREntityProperties().ID))

	object.WriteInit(w.Body)
}

func (w *ClientEntityWriter) CreateComponentAndInit(component DRObject, targetEntity DRObject) {
	w.dirty = true

	w.Body.WriteByte(0x32)                                           // Create Component
	w.Body.WriteUInt16(uint16(targetEntity.RREntityProperties().ID)) // Parent Entity ID
	w.Body.WriteUInt16(uint16(component.RREntityProperties().ID))    // Component ID
	w.Body.WriteByte(byte(GCObjectLookupTypeString))
	w.Body.WriteCString(component.GetGCObject().GCType) // Component Type
	w.Body.WriteByte(0x01)                              // Unk

	component.WriteInit(w.Body)
}

func (w *ClientEntityWriter) Update(object DRObject) {
	w.dirty = true

	w.Body.WriteByte(0x03)                                     // MsgType Update
	w.Body.WriteUInt16(uint16(object.RREntityProperties().ID)) // Entity ID

	object.WriteUpdate(w.Body)

	// EntitySynchInfo
	// Flags
	w.Body.WriteByte(0x0)
}

func (w *ClientEntityWriter) BeginComponentUpdate(object DRObject) {
	if object, ok := object.(IUnitBehavior); ok {
		w.sessionID = object.GetUnitBehavior().SessionID
	} else {
		w.sessionID = 0xFF
	}

	w.dirty = true

	w.Body.WriteByte(0x35)                                     // ComponentUpdate
	w.Body.WriteUInt16(uint16(object.RREntityProperties().ID)) // ComponentID
}

func (w *ClientEntityWriter) EndComponentUpdate(object DRObject) {
	w.dirty = true

	w.WriteSynch(object)
}

func (w *ClientEntityWriter) EndStream() {
	w.Body.WriteByte(0x06)
}

func (w *ClientEntityWriter) WriteSynch(object DRObject) {
	w.dirty = true

	object.WriteSynch(w.Body)
}

func (w *ClientEntityWriter) Clear() {
	w.dirty = false

	w.Body.Clear()
}

func (w *ClientEntityWriter) HasData() bool {
	return w.Body.HasWrittenData()
}

func (w *ClientEntityWriter) GetBody() *byter.Byter {
	return w.Body
}

func (w *ClientEntityWriter) IsDirty() bool {
	return w.dirty
}

func (w *ClientEntityWriter) CreateAction(action actions.BehaviourAction, sessionID byte) {
	w.Body.WriteByte(0x04)
	w.Body.WriteByte(byte(action))
	w.Body.WriteByte(sessionID)
}

func (w *ClientEntityWriter) CreateActionResponse(action actions.BehaviourAction, responseID byte, sessionID byte) {
	w.Body.WriteByte(0x01)
	w.Body.WriteByte(responseID)
	w.Body.WriteByte(byte(action))
	w.Body.WriteByte(sessionID)
}

func (w *ClientEntityWriter) CreateActionComplete(action actions.Action) {
	w.CreateAction(action.OpCode(), w.sessionID)
	action.Init(w.Body)
}

func NewClientEntityWriter(b *byter.Byter) *ClientEntityWriter {
	return &ClientEntityWriter{
		Body: b,
	}
}
func NewClientEntityWriterWithByter() *ClientEntityWriter {
	return NewClientEntityWriter(byter.NewLEByter(make([]byte, 0, 256)))
}
