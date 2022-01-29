package main

import (
	"RainbowRunner/pkg/datatypes"
	"fmt"
	"strings"
)

type OBJWriter struct {
	body strings.Builder
}

func (w *OBJWriter) WriteVert(vert datatypes.Vector3Float32) {
	w.body.WriteString(fmt.Sprintf("v %f %f %f\n", vert.X, vert.Y, vert.Z))
}

func (w *OBJWriter) String() string {
	return w.body.String()
}

func (w *OBJWriter) WriteFace(tri []uint16, withNormals, withUVs bool) {
	fullStr := "f "

	for i := 0; i < 3; i++ {
		str := ""

		if i > 0 {
			str += " "
		}

		index := fmt.Sprintf("%d", tri[i]+1)

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

func NewOBJWriter() *OBJWriter {
	return &OBJWriter{}
}
