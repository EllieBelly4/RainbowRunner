package modelextractor

import (
	"RainbowRunner/internal/objects"
	"RainbowRunner/pkg/datatypes"
	"fmt"
	"strings"
)

type OBJBuilder struct {
	header         strings.Builder
	body           strings.Builder
	vertsThisModel int
	baseFaceIndex  int
	positionOffset datatypes.Vector3Float32
}

func (w *OBJBuilder) WriteVert(vert datatypes.Vector3Float32) {
	vert = vert.Add(w.positionOffset)

	w.body.WriteString(fmt.Sprintf("v %f %f %f\n", vert.X, vert.Y, vert.Z))

	w.vertsThisModel++
}

func (w *OBJBuilder) WriteVertSwizzle(vert datatypes.Vector3Float32) {
	vert = vert.Add(w.positionOffset)

	w.body.WriteString(fmt.Sprintf("v %f %f %f\n", vert.X, vert.Z, vert.Y))

	w.vertsThisModel++
}

func (w *OBJBuilder) String() string {
	return w.header.String() + w.body.String()
}

func (w *OBJBuilder) WriteFace(tri []uint16, withNormals, withUVs bool, flipTriangles bool) {
	fullStr := "f "

	if flipTriangles {
		two := tri[2]
		one := tri[1]
		tri[2] = one
		tri[1] = two
	}

	for i := 0; i < 3; i++ {
		str := ""

		if i > 0 {
			str += " "
		}

		index := fmt.Sprintf("%d", int(tri[i])+1+w.baseFaceIndex)

		str += index

		if withUVs {
			str += "/" + index
		}

		if withNormals {
			str += "/" + index
		}

		fullStr += str
	}

	w.body.WriteString(fullStr + "\n")
}

func (w *OBJBuilder) WriteNormal(norm datatypes.Vector3Float32) {
	w.body.WriteString(fmt.Sprintf("vn %f %f %f\n", norm.X, norm.Y, norm.Z))
}

func (w *OBJBuilder) WriteTextureCoordinates(texcoord datatypes.Vector2Float32) {
	w.body.WriteString(fmt.Sprintf("vt %f %f\n", texcoord.X, texcoord.Y))
}

func (w *OBJBuilder) WriteObject(label string) {
	w.body.WriteString(fmt.Sprintf("o %s\n", label))
	w.baseFaceIndex += w.vertsThisModel
	w.vertsThisModel = 0
}

func (w *OBJBuilder) WriteIncludeMTL(name string) {
	w.header.WriteString("mtllib ")
	w.header.WriteString(name)
	w.header.WriteString("\n")
}

func (w *OBJBuilder) WriteUseMaterial(ref objects.DFCMeshMaterialRef) {
	w.body.WriteString("usemtl ")
	w.body.WriteString(ref.SafeName())
	w.body.WriteRune('\n')
}

func (w *OBJBuilder) SetModelOffset(offset datatypes.Vector3Float32) {
	w.positionOffset = offset
}

func NewOBJBuilder() *OBJBuilder {
	return &OBJBuilder{}
}
