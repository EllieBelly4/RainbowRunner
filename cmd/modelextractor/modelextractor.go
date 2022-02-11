package main

import (
	"RainbowRunner/cmd/rrcli/modelextractor"
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
	prefix := "Townston_graveyard"
	fileType := ".3dnode"
	allFiles, err := os.ReadDir(sourceDir)

	if err != nil {
		return
	}

	outputPath = strings.ReplaceAll(outputPath, "\\", "/")

	objBuilder := modelextractor.NewOBJBuilder()

	pattern := regexp.MustCompile("^" + prefix + ".*" + fileType + "$")

	for _, file := range allFiles {
		if !pattern.MatchString(file.Name()) {
			continue
		}

		fmt.Printf("Extracting from %s\n", file.Name())

		filePathFull := strings.ReplaceAll(path.Join(sourceDir, file.Name()), "\\", "/")
		modelextractor.Extract(filePathFull, objBuilder)
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
