package modelextractor

import (
	"RainbowRunner/internal/objects"
	"RainbowRunner/internal/types"
	"RainbowRunner/pkg/byter"
	"io/ioutil"
	"os"
)

func Extract(pathString string, objBuilder *OBJWriter) {
	file, err := os.Open(pathString)

	if err != nil {
		panic(err)
	}

	data, err := ioutil.ReadAll(file)

	if err != nil {
		panic(err)
	}

	node := objects.ReadData(byter.NewLEByter(data))

	extractFromChildren(node, objBuilder, types.Matrix324x4{
		Values: [16]float32{},
	})

	//if !singleModel {
	//	fileName := fmt.Sprintf("%s_%x.obj", fileNameWithoutExt, mesh.ID())
	//	outputFullPath := fmt.Sprintf("%s/%s", outputDir, fileName)
	//	err := ioutil.WriteFile(outputFullPath, []byte(objBuilder.String()), os.ModePerm)
	//	if err != nil {
	//		panic(err)
	//	}
	//
	//	objBuilder = NewOBJBuilder()
	//}
}

func extractFromChildren(node objects.DRObject, objBuilder *OBJWriter, matrix types.Matrix324x4) {
	//if strings.Contains(node.GetGCObject().GCLabel, "StaticObject") {
	//matrix = matrix.MultiplyMatrix324x4(node.(*objects.DFC3DNode).Matrix)
	//}

	newMatrix := node.(*objects.DFC3DNode).Matrix
	matrix.Values[0] += newMatrix.Values[0]
	matrix.Values[1] += newMatrix.Values[1]
	matrix.Values[2] += newMatrix.Values[2]

	for _, object := range node.Children() {
		if mesh, ok := object.(*objects.DFC3DStaticMeshNode); ok {
			objBuilder.WriteObject(mesh.GetGCObject().GCLabel)

			addMeshToObj(objBuilder, mesh, matrix)
		} else if mesh, ok := object.(*objects.DFC3DNode); ok {
			extractFromChildren(mesh, objBuilder, matrix)
		}
	}
}

func addMeshToObj(objBuilder *OBJWriter, mesh *objects.DFC3DStaticMeshNode, matrix types.Matrix324x4) {
	//offset := datatypes.Vector3Float32{
	//	X: (mesh.MaxBounds.X-mesh.MinBounds.X)/2.0 + mesh.MinBounds.X,
	//	Y: (mesh.MaxBounds.Y-mesh.MinBounds.Y)/2.0 + mesh.MinBounds.Y,
	//	Z: (mesh.MaxBounds.Z-mesh.MinBounds.Z)/2.0 + mesh.MinBounds.Z,
	//}
	for _, vert := range mesh.Verts {
		//offset := mesh.Center

		//objBuilder.WriteVertSwizzle(vert.Sub(offset))
		//objBuilder.WriteVert(vert.Sub(offset))
		//objBuilder.WriteVert(vert)

		vert.X += matrix.Values[0]
		vert.Y += matrix.Values[1]
		vert.Z += matrix.Values[2]

		//vert = vert.Sub(mesh.Center)

		vert = vert.DivideByFloat32(10.0)
		// TODO validate this is correct behaviour, it seems weird but otherwise the town is reversed on x axis and I don't see any obvious scaling
		vert.X *= -1

		objBuilder.WriteVertSwizzle(vert)
	}

	for _, norm := range mesh.Normals {
		objBuilder.WriteNormal(norm.MultiplyByFloat32(-1))
	}

	for _, uv := range mesh.UVs {
		objBuilder.WriteTextureCoordinates(uv)
	}

	for i := 0; i < len(mesh.Triangles); i += 3 {
		objBuilder.WriteFace(mesh.Triangles[i:i+3], len(mesh.Normals) > 0, len(mesh.UVs) > 0, true)
	}
}
