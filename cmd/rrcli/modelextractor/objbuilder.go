package modelextractor

import (
	"RainbowRunner/pkg/datatypes"
	"fmt"
	"strings"
)

type OBJWriter struct {
	body           strings.Builder
	vertsThisModel int
	baseFaceIndex  int
}

func (w *OBJWriter) WriteVert(vert datatypes.Vector3Float32) {
	w.body.WriteString(fmt.Sprintf("v %f %f %f\n", vert.X, vert.Y, vert.Z))

	w.vertsThisModel++
}

func (w *OBJWriter) WriteVertSwizzle(vert datatypes.Vector3Float32) {
	w.body.WriteString(fmt.Sprintf("v %f %f %f\n", vert.X, vert.Z, vert.Y))

	w.vertsThisModel++
}

func (w *OBJWriter) String() string {
	return w.body.String()
}

func (w *OBJWriter) WriteFace(tri []uint16, withNormals, withUVs bool, flipTriangles bool) {
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

func (w *OBJWriter) WriteNormal(norm datatypes.Vector3Float32) {
	w.body.WriteString(fmt.Sprintf("vn %f %f %f\n", norm.X, norm.Y, norm.Z))
}

func (w *OBJWriter) WriteTextureCoordinates(texcoord datatypes.Vector2Float32) {
	w.body.WriteString(fmt.Sprintf("vt %f %f\n", texcoord.X, texcoord.Y))
}

func (w *OBJWriter) WriteObject(label string) {
	w.body.WriteString(fmt.Sprintf("o [%s]\n", label))
	w.baseFaceIndex += w.vertsThisModel
	w.vertsThisModel = 0
}

func NewOBJBuilder() *OBJWriter {
	return &OBJWriter{}
}
