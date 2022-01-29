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

func NewOBJWriter() *OBJWriter {
	return &OBJWriter{}
}
