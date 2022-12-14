package objects

import (
	"RainbowRunner/pkg/byter"
	"RainbowRunner/pkg/datatypes"
	"strings"
)

type DFC3DMeshNodeMaterialGroup struct {
	MaterialID    uint32
	VertexIndex   uint32
	VertexCount   uint32
	TriangleIndex uint32
	TriangleCount uint32
}

type DFC3DStaticMeshNode struct {
	*GCObject
	Materials      []DFCMeshMaterialRef
	MaterialGroups []DFC3DMeshNodeMaterialGroup
	Flags          uint32
	Verts          []datatypes.Vector3Float32 `json:"-"`
	Normals        []datatypes.Vector3Float32 `json:"-"`
	Colours        []datatypes.RGBA32         `json:"-"`
	UVs            []datatypes.Vector2Float32 `json:"-"`
	Triangles      []uint16                   `json:"-"`
	Center         datatypes.Vector3Float32
	MinBounds      datatypes.Vector3Float32
	MaxBounds      datatypes.Vector3Float32
	Angle          float32
}

type DFCMeshMaterialRef struct {
	ID   uint32
	Name string
}

func (d DFCMeshMaterialRef) SafeName() string {
	return strings.Replace(d.Name, " ", "_", -1)
}

func (d *DFC3DStaticMeshNode) ReadData(b *byter.Byter) {
	b.UInt32() // Unk
	b.UInt32() // Unk

	for i := 0; i < 7; i++ {
		b.UInt32() // Unk
	}

	b.Vector3Float32() // Unk

	b.UInt32() // Unk
	b.UInt8()  // Unk
	b.Bytes(6) // Unk - sometimes "Export" string
	b.UInt8()  // Unk

	b.UInt32() // Unk
	b.UInt32() // Unk
	b.UInt32() // Unk

	materialCount := b.UInt32()

	for i := 0; i < int(materialCount); i++ {
		d.Materials = append(d.Materials, DFCMeshMaterialRef{
			ID:   b.UInt32(),
			Name: b.CString(),
		})
	}

	vertCount := b.UInt32()
	d.Flags = b.UInt32()

	for i := 0; i < int(vertCount); i++ {
		d.Verts = append(d.Verts, b.Vector3Float32())
	}

	if (d.Flags & 0x02) > 0 {
		for i := 0; i < int(vertCount); i++ {
			d.Normals = append(d.Normals, b.Vector3Float32())
		}
	}

	if (d.Flags & 0x04) > 0 {
		for i := 0; i < int(vertCount); i++ {
			d.Colours = append(d.Colours, b.RGBA32())
		}
	}

	if (d.Flags & 0x08) > 0 {
		for i := 0; i < int(vertCount); i++ {
			d.UVs = append(d.UVs, b.Vector2Float32())
		}
	}

	triangleCount := b.UInt32()

	for i := 0; i < int(triangleCount); i++ {
		d.Triangles = append(d.Triangles, b.UInt16())
	}

	materialGroupCount := b.UInt32()

	d.MaterialGroups = make([]DFC3DMeshNodeMaterialGroup, 0, materialGroupCount)

	for i := 0; i < int(materialGroupCount); i++ {
		d.MaterialGroups = append(d.MaterialGroups, DFC3DMeshNodeMaterialGroup{
			MaterialID:    b.UInt32(),
			VertexIndex:   b.UInt32(),
			VertexCount:   b.UInt32(),
			TriangleIndex: b.UInt32(),
			TriangleCount: b.UInt32(),
		})
	}

	d.MinBounds = b.Vector3Float32() // Min Bounds
	d.MaxBounds = b.Vector3Float32() // Max Bounds
	d.Center = b.Vector3Float32()    // Center
	d.Angle = b.Float32()            // Angle

	//fmt.Printf("%v, %v, %v, %v\n", av, bv, cv, dv)
}

func NewDFC3DStaticMeshNode() *DFC3DStaticMeshNode {
	return &DFC3DStaticMeshNode{
		GCObject:  NewGCObject("DFC3DStaticMeshNode"),
		Materials: make([]DFCMeshMaterialRef, 0),
		Verts:     make([]datatypes.Vector3Float32, 0),
		Normals:   make([]datatypes.Vector3Float32, 0),
		Colours:   make([]datatypes.RGBA32, 0),
		UVs:       make([]datatypes.Vector2Float32, 0),
		Triangles: make([]uint16, 0),
	}
}
