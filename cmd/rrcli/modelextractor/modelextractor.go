package modelextractor

import (
	"RainbowRunner/cmd/configparser/configparser"
	"RainbowRunner/internal/gosucks"
	"RainbowRunner/internal/objects"
	"RainbowRunner/internal/types"
	"RainbowRunner/internal/types/configtypes"
	"RainbowRunner/internal/types/drobjecttypes"
	"RainbowRunner/pkg/byter"
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"os"
	"path/filepath"
	"sort"
	"strconv"
)

var basePath string

func Extract(pathString string, objBuilder *OBJBuilder, mtlBuilder *MTLBuilder) {
	setBasePath(pathString)

	file, err := os.Open(pathString)

	if err != nil {
		panic(err)
	}

	data, err := io.ReadAll(file)

	if err != nil {
		panic(err)
	}

	node := objects.ReadData(byter.NewLEByter(data))

	if node == nil {
		return
	}

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

func splitObjects(node drobjecttypes.DRObject, f func(writer *OBJBuilder, mtlBuilder *MTLBuilder, meshNode *objects.DFC3DStaticMeshNode)) {
	for _, child := range node.Children() {
		objBuilder := NewOBJBuilder()
		mtlBuilder := NewMTLBuilder()

		gcObject := child.(objects.IGCObject).GetGCObject()

		fmt.Printf("%d %s\n", gcObject.ID(), gcObject.GCLabel)

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

func extractFromChildren(node drobjecttypes.DRObject, objBuilder *OBJBuilder, mtlBuilder *MTLBuilder, matrix types.Matrix324x4) {
	depth++

	if d3Node, ok := node.(*objects.DFC3DNode); ok {
		newMatrix := d3Node.Matrix
		matrix.Values[0] += newMatrix.Values[0]
		matrix.Values[1] += newMatrix.Values[1]
		matrix.Values[2] += newMatrix.Values[2]
	}

	for _, object := range node.Children() {
		if mesh, ok := object.(*objects.DFC3DStaticMeshNode); ok {
			subMatrix := types.Matrix324x4{Values: [16]float32{
				matrix.Values[0],
				matrix.Values[1],
				matrix.Values[2],
			}}

			addMaterials(mesh, objBuilder, mtlBuilder)
			addMeshToObj(objBuilder, mesh, subMatrix)
		} else if mesh, ok := object.(*objects.DFC3DNode); ok {
			extractFromChildren(mesh, objBuilder, mtlBuilder, matrix)
		}
	}

	depth--
}

func addMaterials(mesh *objects.DFC3DStaticMeshNode, objBuilder *OBJBuilder, mtlBuilder *MTLBuilder) {
	for _, materialRef := range mesh.Materials {
		err := createMaterial(materialRef, objBuilder, mtlBuilder)

		if err != nil {
			fmt.Printf("could not add material %s: %s\n", materialRef.SafeName(), err.Error())
		}
	}
}

func createMaterial(ref objects.DFCMeshMaterialRef, objBuilder *OBJBuilder, mtlBuilder *MTLBuilder) error {
	matFilePath := filepath.Join(basePath, ref.SafeName()+".mat")

	drConfig, err := configparser.ParseAllFilesToDRConfig([]string{matFilePath}, basePath)

	if err != nil {
		return err
	}

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
	if !mtlBuilder.HasMaterial(ref.SafeName()) {
		mtlBuilder.WriteNewMaterial(ref.SafeName())
		for childName, childGroup := range drConfig.Classes.Children["material"].Entities[0].Children {
			if childName == "texture" {
				for _, texture := range childGroup.Entities {
					textureFileName := texture.Properties["Filename"] + ".dds"
					mtlBuilder.WriteNewTexture(ref.SafeName(), MTLTexture{
						Type:     MTLTextureTypeDiffuse,
						Filename: textureFileName,
					})
				}
				continue
			} else if childName == "materialcolor" {
				for _, colourGroup := range childGroup.Entities {
					for colourType, colour := range colourGroup.Children {
						colourProperties := colour.Entities[0].Properties

						var colourTypeEnum MTLColourType

						switch colourType {
						case "specular":
							colourTypeEnum = MTLColourTypeSpecular
						case "diffuse":
							colourTypeEnum = MTLColourTypeDiffuse
						case "emissive":
							colourTypeEnum = MTLColourTypeAmbient
						default:
							panic(fmt.Sprintf("unhandled colour type %s", colourType))
						}

						r := parseColour(colourProperties, "Red")
						g := parseColour(colourProperties, "Green")
						b := parseColour(colourProperties, "Blue")
						a := parseColour(colourProperties, "Alpha")

						if r != nil && g != nil && b != nil {
							mtlBuilder.WriteNewColour(ref.SafeName(), MTLColour{
								Type: colourTypeEnum,
								R:    *r,
								G:    *g,
								B:    *b,
							})
						}

						if a != nil {
							mtlBuilder.WriteNewAlpha(ref.SafeName(), *a)
						}
					}
				}

				continue
			}

			panic(fmt.Sprintf("unknown child %s", childName))

		}
		//map_Kd -s 1 1 1 -o 0 0 0 -mm 0 1 chrome.mpc
	}

	return nil
}

func parseColour(colourProperties configtypes.DRClassProperties, property string) *float32 {
	if rString, ok := colourProperties[property]; ok {
		val, err := strconv.ParseInt(rString, 10, 32)

		if err != nil {
			panic(err)
		}

		result := float32(val) / 255

		return &result
	}

	return nil
}

func addMeshToObj(objBuilder *OBJBuilder, mesh *objects.DFC3DStaticMeshNode, matrix types.Matrix324x4) {
	//for _, materialRef := range mesh.Materials {
	//	gosucks.VAR(materialRef)
	//}

	fmt.Println("adding mesh " + mesh.GCLabel)

	matGroups := mesh.MaterialGroups

	sort.Slice(matGroups, func(i, j int) bool {
		return matGroups[i].TriangleIndex < matGroups[j].TriangleIndex
	})

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

	materialGroundIndex := 0
	currentMaterialGroup := matGroups[materialGroundIndex]

	foundMat := false

	for _, material := range mesh.Materials {
		if material.ID == currentMaterialGroup.MaterialID {
			objBuilder.WriteUseMaterial(material)
			foundMat = true
			break
		}
	}

	if !foundMat {
		for _, material := range mesh.Materials {
			if material.ID == 0 {
				objBuilder.WriteUseMaterial(material)
				break
			}
		}
	}

	for i := 0; i < len(mesh.Triangles); i += 3 {
		objBuilder.WriteFace(mesh.Triangles[i:i+3], len(mesh.Normals) > 0, len(mesh.UVs) > 0, true)

		if i == int(currentMaterialGroup.TriangleIndex+currentMaterialGroup.TriangleCount) {
			materialGroundIndex++

			if materialGroundIndex < len(matGroups) {
				currentMaterialGroup = matGroups[materialGroundIndex]

				foundMat := false

				for _, material := range mesh.Materials {
					if material.ID == currentMaterialGroup.MaterialID {
						objBuilder.WriteUseMaterial(material)
						foundMat = true
						break
					}
				}

				if !foundMat {
					fmt.Printf("could not find material with ID %d, ignoring", currentMaterialGroup.MaterialID)
				}
			}
		}
	}
}
