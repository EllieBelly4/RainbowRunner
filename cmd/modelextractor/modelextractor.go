package main

import (
	"RainbowRunner/cmd/rrcli/modelextractor"
	"errors"
	"fmt"
	"io/ioutil"
	"os"
	"path"
	"path/filepath"
	"regexp"
	"strings"
)

var quoteRegex = regexp.MustCompile("[\"']")

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

	//configDumpPath := "resources/Dumps/generated/finalconf.json"
	//modelextractor.LoadConfig(configDumpPath)

	sourceDir := "D:\\Work\\dungeon-runners\\666 dumps new"
	//prefix := "Townston_graveyard"
	prefix := "Townston_"
	fileType := ".3dnode"
	allFiles, err := os.ReadDir(sourceDir)

	if err != nil {
		return
	}

	outputPath = strings.ReplaceAll(outputPath, "\\", "/")

	objBuilder := modelextractor.NewOBJBuilder()
	mtlBuilder := modelextractor.NewMTLBuilder()

	pattern := regexp.MustCompile("^" + prefix + ".*" + fileType + "$")

	for _, file := range allFiles {
		if !pattern.MatchString(file.Name()) {
			continue
		}

		fmt.Printf("Extracting from %s\n", file.Name())

		filePathFull := strings.ReplaceAll(path.Join(sourceDir, file.Name()), "\\", "/")
		modelextractor.Extract(filePathFull, objBuilder, mtlBuilder)
	}

	//fileNameWithoutExt := strings.Split(path.Base(pathString), ".")[0]
	outputDir := filepath.Join(fmt.Sprintf("%s", outputPath), prefix+"_combined")

	if _, err := os.Stat(outputDir); err != nil {
		if errors.Is(err, os.ErrNotExist) {
			err := os.MkdirAll(outputDir, os.ModeDir)

			if err != nil {
				panic(err)
			}
		} else {
			panic(err)
		}
	}

	objFileName := fmt.Sprintf("%s_combined.obj", prefix)
	mtlFileName := fmt.Sprintf("%s.mtl", prefix)

	objBuilder.WriteIncludeMTL(mtlFileName)

	fullOuputOBJFilePath := fmt.Sprintf("%s/%s", outputDir, objFileName)
	err = ioutil.WriteFile(fullOuputOBJFilePath, []byte(objBuilder.String()), os.ModePerm)
	if err != nil {
		panic(err)
	}

	fullOuputMTLFilePath := fmt.Sprintf("%s/%s", outputDir, mtlFileName)
	err = ioutil.WriteFile(fullOuputMTLFilePath, []byte(mtlBuilder.String()), os.ModePerm)

	for _, textureFilename := range mtlBuilder.TextureFilenames() {
		textureFilename = quoteRegex.ReplaceAllString(textureFilename, "")

		textureFullPath := filepath.Join(sourceDir, textureFilename)

		textureData, err := ioutil.ReadFile(textureFullPath)

		if err != nil {
			panic(err)
		}

		err = ioutil.WriteFile(fmt.Sprintf("%s/%s", outputDir, textureFilename), textureData, 0755)

		if err != nil {
			panic(err)
		}
	}
}
