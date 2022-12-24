package objects

import (
	"RainbowRunner/internal/config"
	"RainbowRunner/internal/connections"
	"RainbowRunner/internal/types/drobjecttypes"
	"RainbowRunner/pkg/byter"
	"fmt"
	"regexp"
	"strings"
)

//var currentID = uint32(0)

type EntityMessageHandler interface {
	WriteInit(b *byter.Byter)
	WriteUpdate(b *byter.Byter)
	WriteSynch(b *byter.Byter)

	ReadUpdate(reader *byter.Byter) error
}

//go:generate go run ../../scripts/generatelua -type=GCObject
type GCObject struct {
	EntityProperties RREntityProperties
	Version          uint8
	GCNativeType     string
	GCLabel          string
	GCChildren       []drobjecttypes.DRObject
	GCType           string
	Properties       []GCObjectProperty
	EntityHandler    EntityMessageHandler
	GCParent         drobjecttypes.DRObject
}

func (g *GCObject) GetChildrenByGCNativeType(s string) []drobjecttypes.DRObject {
	results := make([]drobjecttypes.DRObject, 0)

	for _, child := range g.GCChildren {
		if child.(IGCObject).GetGCObject().GCNativeType == s {
			results = append(results, child)
		}
	}

	return results
}

func (g *GCObject) GetRREntityProperties() *RREntityProperties {
	return g.RREntityProperties()
}

type IRREntityPropertiesHaver interface {
	GetRREntityProperties() *RREntityProperties
}

func (g *GCObject) GetParentEntity() drobjecttypes.DRObject {
	if g.GCParent == nil {
		return nil
	}

	entity, ok := g.GCParent.(IEntity)

	if !ok {
		if entity == nil {
			return nil
		}
	} else {
		return g.GCParent
	}

	return g.GCParent.GetParentEntity()
}

func (g *GCObject) SetParent(parent drobjecttypes.DRObject) {
	g.GCParent = parent
}

func (g *GCObject) GCObject() *GCObject {
	return g
}

func (g *GCObject) String() string {
	return fmt.Sprintf("(%d - 0x%x) %s OwnedBy: %d",
		g.EntityProperties.ID, g.EntityProperties.ID,
		g.GCType,
		g.EntityProperties.OwnerID,
	)
}

func (g *GCObject) SetOwner(conn *connections.RRConn) {
	g.RREntityProperties().SetOwner(uint16(conn.GetID()))
}

func (g *GCObject) RemoveChildAt(i int) {
	g.GCChildren = append(g.GCChildren[:i], g.GCChildren[i+1:]...)
}

func (g *GCObject) Type() drobjecttypes.DRObjectType {
	return drobjecttypes.DRObjectUnknown
}

func (g *GCObject) ID() int {
	return int(g.EntityProperties.ID)
}

func (g *GCObject) ReadUpdate(reader *byter.Byter) error {
	fmt.Printf("Unhandled readupdate for %s (%s : %s) ID: %x\n", g.GCLabel, g.GCType, g.GCNativeType, g.EntityProperties.ID)
	return nil
}

func (g *GCObject) WriteSynch(b *byter.Byter) {
	flag := 0x02
	// TODO consider checking the zone to see if it's a town, as it is 0x02 will work in town
	//b.WriteByte(0x00) // 0x00 If in town
	b.WriteByte(byte(flag)) // 0x02 If in dungeon

	if flag == 0x02 {
		// This value must be exactly 0x47E00 but I do not know why.
		b.WriteUInt32(0x47E00) // Unk - EntitySynchInfo::ReadFromStream
	}
}

func (g *GCObject) Tick() {

}

func (g *GCObject) Init() {

}

func (g *GCObject) WriteInit(b *byter.Byter) {
	fmt.Printf("GCObject init for %s (%s: %s) not implemented but ignoring\n", g.GCLabel, g.GCType, g.GCNativeType)
}

func (g *GCObject) WriteUpdate(b *byter.Byter) {
	panic("implement me")
}

func (g *GCObject) OwnerID() uint16 {
	return g.EntityProperties.OwnerID
}

func (g *GCObject) Children() []drobjecttypes.DRObject {
	return g.GCChildren
}

func (g *GCObject) RREntityProperties() *RREntityProperties {
	return &g.EntityProperties
}

type GCObjectProperty struct {
	Name  interface{}
	Value interface{}
}

func Uint32Prop(name interface{}, val uint32) GCObjectProperty {
	return GCObjectProperty{
		Name:  name,
		Value: val,
	}
}

func StringProp(name interface{}, val string) GCObjectProperty {
	return GCObjectProperty{
		Name:  name,
		Value: val,
	}
}

func NewGCObject(nativeType string) *GCObject {
	return &GCObject{
		// At version 2A or above you must use a hash I think
		//Version:    0x29, // No hash required
		Version:      0x2D,
		GCNativeType: nativeType,
		GCType:       strings.ToLower(nativeType),
	}
}

var indent = 0

func (g *GCObject) WriteFullGCObject(byter *byter.Byter) {
	byter.WriteByte(g.Version)

	useHashes := g.Version >= 0x2a

	logSerialise("========== GCObject ===========")
	logSerialise(`ID: %d
Name: %s
NativeClass: %s
GCType: %s
---`, g.EntityProperties.ID, g.GCLabel, g.GCNativeType, g.GCType)

	if useHashes {
		byter.WriteUInt32(GetTypeHash(g.GCNativeType))
	} else {
		byter.WriteCString(g.GCNativeType)
	}

	byter.WriteUInt32(uint32(g.EntityProperties.ID))
	byter.WriteCString(g.GCLabel)

	byter.WriteUInt32(uint32(len(g.GCChildren)))

	indent++
	for _, child := range g.GCChildren {
		child.WriteFullGCObject(byter)
	}
	indent--

	if useHashes {
		byter.WriteUInt32(GetTypeHash(g.GCType))
	} else {
		byter.WriteCString(g.GCType)
	}

	for _, prop := range g.Properties {
		prop.Serialise(byter, useHashes)
	}

	byter.WriteUInt32(0)
}

func logSerialise(format string, args ...interface{}) {
	regex := regexp.MustCompile("(?m)^")

	if config.Config.Logging.LogGCObjectSerialise {
		indentString := strings.Repeat("\t", indent)
		log := fmt.Sprintf(format, args...)
		log = regex.ReplaceAllString(log, indentString)
		fmt.Print(log + "\n")
	}
}

func (g *GCObject) AddChild(child drobjecttypes.DRObject) {
	if g.GCChildren == nil {
		g.GCChildren = make([]drobjecttypes.DRObject, 0, 128)
	}

	child.SetParent(g)

	g.GCChildren = append(g.GCChildren, child)
}

func (p GCObjectProperty) Serialise(b *byter.Byter, useHash bool) {
	switch name := p.Name.(type) {
	case string:
		if useHash {
			b.WriteUInt32(GetTypeHash(name))
		} else {
			b.WriteCString(name)
		}
	case int:
		b.WriteUInt32(uint32(name))
	case uint32:
		b.WriteUInt32(name)
	}

	switch p.Value.(type) {
	case string:
		b.WriteCString(p.Value.(string))
	case uint32:
		b.WriteUInt32(p.Value.(uint32))
	case uint8:
		b.WriteByte(p.Value.(uint8))
	case uint16:
		b.WriteUInt16(p.Value.(uint16))
	}
}

func GetTypeHash(name string) uint32 {
	result := uint32(5381) // eax

	a1 := len(name)

	if a1 > 0 {
		for _, v4 := range name {
			if v4 >= 0x41 && v4 <= 0x5A {
				v4 = v4 + 32
			}

			result += uint32(v4) + 32*result
		}

		if result == 0 {
			result = 1
		}
	}

	if config.Config.Logging.LogHashes {
		fmt.Printf("(%x) %s\n", result, name)
	}

	return result
}

func (g *GCObject) GetChildByGCNativeType(s string) drobjecttypes.DRObject {
	for _, child := range g.GCChildren {
		if strings.ToLower(child.(IGCObject).GetGCObject().GCNativeType) == strings.ToLower(s) {
			return child
		}
	}

	return nil
}

func (g *GCObject) GetChildByGCType(s string) drobjecttypes.DRObject {
	for _, child := range g.GCChildren {
		if strings.ToLower(child.(IGCObject).GetGCObject().GCType) == strings.ToLower(s) {
			return child
		}
	}

	return nil
}

func (g *GCObject) SetVersion(version uint8) {
	g.Version = version
}

func (g *GCObject) ReadData(b *byter.Byter) {
}

func (g *GCObject) WalkChildren(cb func(object drobjecttypes.DRObject)) {
	if len(g.Children()) == 0 {
		return
	}

	for _, object := range g.Children() {
		object.WalkChildren(cb)

		cb(object)
	}
}

func ReadData(b *byter.Byter) drobjecttypes.DRObject {
	version := b.Byte() // Version
	nativeType := b.CString()
	id := b.UInt32()

	defer func() {
		if r := recover(); r != nil {
			fmt.Printf("child was not parsed successfully, abandoning readData. parent: %s - id %x\n",
				nativeType,
				id,
			)
		}
	}()

	var gcObject drobjecttypes.DRObject

	switch nativeType {
	case "DFC3DNode":
		gcObject = NewDFC3DNode()
	case "DFC3DStaticMeshNode":
		gcObject = NewDFC3DStaticMeshNode()
	case "AdvParticleSystem":
		return nil
	default:
		gcObject = NewGCObject(nativeType)
	}

	gcObject.SetVersion(version)
	gcObject.(IRREntityPropertiesHaver).GetRREntityProperties().ID = id

	gcName := b.CString()
	gcObject.(IGCObject).GetGCObject().GCLabel = gcName

	childCount := b.UInt32()

	for i := 0; i < int(childCount); i++ {
		child := ReadData(b)

		if child == nil {
			return nil
		}

		gcObject.AddChild(child)
	}

	gcObject.ReadData(b)

	return gcObject
}

func (g *GCObject) RemoveChildrenByGCNativeType(gcNativeType string) int {
	toRemove := make([]int, 0, 0)

	for i, child := range g.Children() {
		if child.(IGCObject).GetGCObject().GCNativeType == gcNativeType {
			toRemove = append(toRemove, i)
		}
	}

	for _, i := range toRemove {
		g.RemoveChildAt(i)
	}

	return len(toRemove)
}
