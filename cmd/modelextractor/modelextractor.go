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
			fmt.Println(result)
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
			fmt.Println(fileName)
			//ioutil.WriteFile()
		}
	})

	fmt.Printf("%+v\n", node)
}

func convertMeshToObj(mesh *objects.DFC3DStaticMeshNode) string {
	objBuilder := NewOBJWriter()

	for _, vert := range mesh.Verts {
		objBuilder.WriteVert(vert)
	}

	return objBuilder.String()
}
