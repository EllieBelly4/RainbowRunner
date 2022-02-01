package main

import (
	"RainbowRunner/internal/objects"
	"RainbowRunner/internal/types"
	"RainbowRunner/pkg/byter"
	"errors"
	"fmt"
	"io/ioutil"
	"os"
	"path"
	"regexp"
	"strings"
)

func main() {
	//singleModel := true
	outputPath := os.Args[1]
	//files := os.Args[2:]

	//files = []string{
	//	//"D:\\Work\\dungeon-runners\\666 dumps new\\TownFloor40.3dnode",
	//	//"D:\\Work\\dungeon-runners\\666 dumps new\\TownFloor10.3dnode",
	//	"D:\\Work\\dungeon-runners\\666 dumps new\\townExit_1.3dnode",
	//	"D:\\Work\\dungeon-runners\\666 dumps new\\town_northEast_1.3dnode",
	//	"D:\\Work\\dungeon-runners\\666 dumps new\\town_lower_1.3dnode",
	//	"D:\\Work\\dungeon-runners\\666 dumps new\\town_upperMiddle_1.3dnode",
	//	"D:\\Work\\dungeon-runners\\666 dumps new\\town_east_1.3dnode",
	//	"D:\\Work\\dungeon-runners\\666 dumps new\\town_westCliff_1.3dnode",
	//	"D:\\Work\\dungeon-runners\\666 dumps new\\town_northWest_1.3dnode",
	//	//"D:\\Work\\dungeon-runners\\666 dumps new\\Townston_Square.3dnode",
	//	//"D:\\Work\\dungeon-runners\\666 dumps new\\AutumnForest_DirtDeadEnd_1.3dnode",
	//	//"D:\\Work\\dungeon-runners\\666 dumps new\\Townston_tier_1.3dnode",
	//	//"D:\\Work\\dungeon-runners\\666 dumps new\\Townston_bank.3dnode",
	//	//"D:\\Work\\dungeon-runners\\666 dumps new\\Townston_graveyard.3dnode",
	//	//"D:\\Work\\dungeon-runners\\666 dumps new\\throne.3dnode",
	//}

	sourceDir := "D:\\Work\\dungeon-runners\\666 dumps new"
	prefix := "manaPotion_"
	fileType := ".3dnode"
	allFiles, err := os.ReadDir(sourceDir)

	if err != nil {
		return
	}

	outputPath = strings.ReplaceAll(outputPath, "\\", "/")

	objBuilder := NewOBJWriter()

	pattern := regexp.MustCompile("^" + prefix + ".*" + fileType + "$")

	for _, file := range allFiles {
		if !pattern.MatchString(file.Name()) {
			continue
		}

		fmt.Printf("Extracting from %s\n", file.Name())

		filePathFull := strings.ReplaceAll(path.Join(sourceDir, file.Name()), "\\", "/")
		extract(filePathFull, outputPath, objBuilder)
	}

	//fileNameWithoutExt := strings.Split(path.Base(pathString), ".")[0]
	outputDir := fmt.Sprintf("%s", outputPath)

	if _, err := os.Stat(outputDir); err != nil {
		if errors.Is(err, os.ErrNotExist) {
			err := os.Mkdir(outputDir, os.ModeDir)

			if err != nil {
				panic(err)
			}
		} else {
			panic(err)
		}
	}

	fileName := fmt.Sprintf("%s_combined.obj", prefix)
	outputFullPath := fmt.Sprintf("%s/%s", outputDir, fileName)
	err = ioutil.WriteFile(outputFullPath, []byte(objBuilder.String()), os.ModePerm)
	if err != nil {
		panic(err)
	}
}

func extract(pathString string, outputPath string, objBuilder *OBJWriter) {
	file, err := os.Open(pathString)

	if err != nil {
		panic(err)
	}

	data, err := ioutil.ReadAll(file)

	if err != nil {
		panic(err)
	}

	node := objects.ReadData(byter.NewLEByter(data))

	extractFromChildren(node, objBuilder, types.Matrix324x4Identity)

	//if !singleModel {
	//	fileName := fmt.Sprintf("%s_%x.obj", fileNameWithoutExt, mesh.ID())
	//	outputFullPath := fmt.Sprintf("%s/%s", outputDir, fileName)
	//	err := ioutil.WriteFile(outputFullPath, []byte(objBuilder.String()), os.ModePerm)
	//	if err != nil {
	//		panic(err)
	//	}
	//
	//	objBuilder = NewOBJWriter()
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
	for _, vert := range mesh.Verts {
		//offset := mesh.Center

		//objBuilder.WriteVertSwizzle(vert.Sub(offset))
		//objBuilder.WriteVert(vert.Sub(offset))
		//objBuilder.WriteVert(vert)

		vert.X += matrix.Values[0]
		vert.Y += matrix.Values[1]
		vert.Z += matrix.Values[2]

		vert = vert.DivideByFloat32(10.0)
		// TODO validate this is correct behaviour, it seems weird but otherwise the town is reversed on x axis and I don't see any obvious scaling
		vert.X *= -1

		objBuilder.WriteVertSwizzle(vert)
	}

	for _, norm := range mesh.Normals {
		objBuilder.WriteNormal(norm)
	}

	for _, uv := range mesh.UVs {
		objBuilder.WriteTextureCoordinates(uv)
	}

	for i := 0; i < len(mesh.Triangles); i += 3 {
		objBuilder.WriteFace(mesh.Triangles[i:i+3], len(mesh.Normals) > 0, len(mesh.UVs) > 0)
	}
}
