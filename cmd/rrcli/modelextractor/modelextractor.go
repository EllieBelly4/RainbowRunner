package modelextractor

import (
	"RainbowRunner/cmd/configparser/configparser"
	"RainbowRunner/cmd/rrcli/configurator"
	"RainbowRunner/internal/gosucks"
	"RainbowRunner/internal/objects"
	"RainbowRunner/internal/types"
	"RainbowRunner/pkg/byter"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
)

var basePath string
var config *configparser.DRConfig

func LoadConfig(configPath string) {
	fmt.Println("loading dumped config")

	var err error
	config, err = configurator.LoadFromDumpedConfigFile(configPath)

	if err != nil {
		panic(err)
	}

	fmt.Println("config load complete")
}

func SetConfig(drConfig *configparser.DRConfig) {
	config = drConfig
}

func Extract(pathString string, objBuilder *OBJBuilder, mtlBuilder *MTLBuilder) {
	setBasePath(pathString)

	file, err := os.Open(pathString)

	if err != nil {
		panic(err)
	}

	data, err := ioutil.ReadAll(file)

	if err != nil {
		panic(err)
	}

	node := objects.ReadData(byter.NewLEByter(data))

	extractFromChildren(node, objBuilder, mtlBuilder, types.Matrix324x4{
		Values: [16]float32{},
	})
}

func setBasePath(pathString string) {
	basePath = filepath.Dir(pathString)
}

func Split(pathString string, destPath string) {
	setBasePath(pathString)

	file, err := os.Open(pathString)

	if err != nil {
		panic(err)
	}

	data, err := ioutil.ReadAll(file)

	if err != nil {
		panic(err)
	}

	err = os.MkdirAll(destPath, 0755)

	if err != nil {
		panic(err)
	}

	node := objects.ReadData(byter.NewLEByter(data))

	splitObjects(node, func(writer *OBJBuilder, mtlBuilder *MTLBuilder, meshNode *objects.DFC3DStaticMeshNode) {
		err := ioutil.WriteFile(
			filepath.Join(
				destPath,
				fmt.Sprintf("%d_%s.obj",
					meshNode.GetGCObject().ID(),
					meshNode.GCLabel),
			),
			[]byte(writer.String()), 0755)

		if err != nil {
			panic(err)
		}
	})

	bytes, err := json.Marshal(node)

	if err != nil {
		panic(err)
	}

	err = ioutil.WriteFile(filepath.Join(destPath, "3dnode.json"), bytes, 0755)

	if err != nil {
		panic(err)
	}
}

func splitObjects(node objects.DRObject, f func(writer *OBJBuilder, mtlBuilder *MTLBuilder, meshNode *objects.DFC3DStaticMeshNode)) {
	for _, child := range node.Children() {
		objBuilder := NewOBJBuilder()
		mtlBuilder := NewMTLBuilder()

		fmt.Printf("%d %s\n", child.GetGCObject().ID(), child.GetGCObject().GCLabel)

		if d3node, ok := child.(*objects.DFC3DNode); ok {
			splitObjects(d3node, f)
		} else if meshNode, ok := child.(*objects.DFC3DStaticMeshNode); ok {
			addMeshToObj(objBuilder, meshNode, types.Matrix324x4{
				Values: [16]float32{},
			})

			addMaterials(meshNode, objBuilder, mtlBuilder)

			f(objBuilder, mtlBuilder, meshNode)
		}
	}
}

var depth = -1

func extractFromChildren(node objects.DRObject, objBuilder *OBJBuilder, mtlBuilder *MTLBuilder, matrix types.Matrix324x4) {
	//if strings.Contains(node.GetGCObject().GCLabel, "StaticObject") {
	//matrix = matrix.MultiplyMatrix324x4(node.(*objects.DFC3DNode).Matrix)
	//}
	depth++

	if d3Node, ok := node.(*objects.DFC3DNode); ok {
		newMatrix := d3Node.Matrix
		matrix.Values[0] += newMatrix.Values[0]
		matrix.Values[1] += newMatrix.Values[1]
		matrix.Values[2] += newMatrix.Values[2]
	}

	//pad := strings.Repeat(" ", depth)
	//
	//fmt.Printf("%smatrix position %f, %f, %f\n", pad, matrix.Values[0], matrix.Values[1], matrix.Values[2])

	for _, object := range node.Children() {
		if mesh, ok := object.(*objects.DFC3DStaticMeshNode); ok {
			subMatrix := types.Matrix324x4{Values: [16]float32{
				matrix.Values[0],
				matrix.Values[1],
				matrix.Values[2],
			}}

			addMeshToObj(objBuilder, mesh, subMatrix)
			addMaterials(mesh, objBuilder, mtlBuilder)
		} else if mesh, ok := object.(*objects.DFC3DNode); ok {
			extractFromChildren(mesh, objBuilder, mtlBuilder, matrix)
		}
	}

	depth--
}

func addMaterials(mesh *objects.DFC3DStaticMeshNode, objBuilder *OBJBuilder, mtlBuilder *MTLBuilder) {
	for _, materialRef := range mesh.Materials {
		material := createMaterial(materialRef, objBuilder, mtlBuilder)
		//addMaterialToObj(material)
		gosucks.VAR(material)
	}
}

func createMaterial(ref objects.DFCMeshMaterialRef, objBuilder *OBJBuilder, mtlBuilder *MTLBuilder) string {
	matFilePath := filepath.Join(basePath, ref.Name+".mat")

	drConfig, err := configparser.ParseAllFilesToDRConfig([]string{matFilePath}, basePath)

	gosucks.VAR(drConfig)

	if err != nil {
		panic(err)
	}

	//The options for the map_Kd statement are listed below.  These options
	//are described in detail in "Options for texture map statements" on page
	//5-18.
	//
	// 	-blendu on | off
	// 	-blendv on | off
	// 	-cc on | off
	// 	-clamp on | off
	// 	-mm base gain
	// 	-o u v w
	// 	-s u v w
	// 	-t u v w
	// 	-texres value
	if !mtlBuilder.HasMaterial(ref.Name) {
		mtlBuilder.WriteNewMaterial(ref.Name)
		for childName, childGroup := range drConfig.Classes.Children["material"].Entities[0].Children {
			if childName != "texture" {
				panic(fmt.Sprintf("unknown child %s", childName))
			}

			for _, texture := range childGroup.Entities {
				textureFileName := texture.Properties["Filename"] + ".dds"
				mtlBuilder.WriteNewTexture(ref.Name, MTLTexture{
					Type:     MTLTextureTypeDiffuse,
					Filename: textureFileName,
				})
			}
		}
		//map_Kd -s 1 1 1 -o 0 0 0 -mm 0 1 chrome.mpc
	}

	return ""
}

func addMeshToObj(objBuilder *OBJBuilder, mesh *objects.DFC3DStaticMeshNode, matrix types.Matrix324x4) {
	//for _, materialRef := range mesh.Materials {
	//	gosucks.VAR(materialRef)
	//}

	fmt.Println(mesh.GCLabel)
	fmt.Println(len(mesh.Materials))

	objBuilder.WriteUseMaterial(mesh.Materials[0])

	objBuilder.WriteObject(mesh.GetGCObject().GCLabel)

	for _, vert := range mesh.Verts {
		vert.X += matrix.Values[0]
		vert.Y += matrix.Values[1]
		vert.Z += matrix.Values[2]

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
