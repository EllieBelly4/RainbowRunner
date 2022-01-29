package main

import (
	"RainbowRunner/internal/objects"
	"RainbowRunner/pkg/byter"
	"errors"
	"fmt"
	"io/ioutil"
	"os"
	"path"
	"strings"
)

func main() {
	outputPath := os.Args[1]
	files := os.Args[2:]

	files = []string{
		//"D:\\Work\\dungeon-runners\\666 dumps new\\TownFloor40.3dnode",
		//"D:\\Work\\dungeon-runners\\666 dumps new\\TownFloor10.3dnode",
		//"D:\\Work\\dungeon-runners\\666 dumps new\\townExit_1.3dnode",
		"D:\\Work\\dungeon-runners\\666 dumps new\\Townston_Square.3dnode",
	}

	outputPath = strings.ReplaceAll(outputPath, "\\", "/")

	for _, file := range files {
		file = strings.ReplaceAll(file, "\\", "/")
		extract(file, outputPath)
	}
}

func extract(pathString string, outputPath string) {
	file, err := os.Open(pathString)

	if err != nil {
		panic(err)
	}

	data, err := ioutil.ReadAll(file)

	if err != nil {
		panic(err)
	}

	node := objects.ReadData(byter.NewLEByter(data))

	node.WalkChildren(func(object objects.DRObject) {
		if mesh, ok := object.(*objects.DFC3DStaticMeshNode); ok {
			result := convertMeshToObj(mesh)
			fileNameWithoutExt := strings.Split(path.Base(pathString), ".")[0]

			outputDir := fmt.Sprintf("%s/%s", outputPath, fileNameWithoutExt)

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

			fileName := fmt.Sprintf("%s_%x.obj", fileNameWithoutExt, mesh.ID())
			outputFullPath := fmt.Sprintf("%s/%s", outputDir, fileName)
			err := ioutil.WriteFile(outputFullPath, []byte(result), os.ModePerm)
			if err != nil {
				panic(err)
			}
		}
	})
}

func convertMeshToObj(mesh *objects.DFC3DStaticMeshNode) string {
	objBuilder := NewOBJWriter()

	for _, vert := range mesh.Verts {
		objBuilder.WriteVert(vert)
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

	return objBuilder.String()
}
