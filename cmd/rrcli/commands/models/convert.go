package models

import (
	"RainbowRunner/cmd/rrcli/configurator"
	"RainbowRunner/cmd/rrcli/modelextractor"
	"RainbowRunner/internal/gosucks"
	"errors"
	"fmt"
	"github.com/spf13/cobra"
	"os"
	"path"
	"path/filepath"
	"regexp"
	"strings"
)

var quoteRegex = regexp.MustCompile("[\"']")

var modelOutputDir string
var modelPrefix string
var modelName string
var combineModels bool

var convertCommand = &cobra.Command{
	Use: "convert",
	Run: func(cmd *cobra.Command, args []string) {
		config, err := configurator.LoadFromDumpedConfigFile(configFile)

		if err != nil {
			panic(err)
		}

		fileType := ".3dnode"
		allFiles, err := os.ReadDir(modelSourceDir)

		if err != nil {
			return
		}

		modelOutputDir = strings.ReplaceAll(modelOutputDir, "\\", "/")
		outputName := "output"

		objBuilder := modelextractor.NewOBJBuilder()
		mtlBuilder := modelextractor.NewMTLBuilder()

		pattern := regexp.MustCompile(fileType + "$")

		if modelPrefix != "" {
			outputName = modelPrefix
			pattern = regexp.MustCompile("^" + modelPrefix + ".*" + fileType + "$")
		} else if modelName != "" {
			outputName = modelName
			pattern = regexp.MustCompile("^" + modelName + fileType + "$")
		}

		outputDir := filepath.Join(fmt.Sprintf("%s", modelOutputDir), outputName)
		mustCreateDir(outputDir)

		for _, file := range allFiles {
			if !pattern.MatchString(file.Name()) {
				continue
			}

			fmt.Printf("Extracting from %s\n", file.Name())

			filePathFull := strings.ReplaceAll(path.Join(modelSourceDir, file.Name()), "\\", "/")
			modelextractor.Extract(filePathFull, objBuilder, mtlBuilder)

			fileNameWithoutExt := strings.Split(path.Base(file.Name()), ".")[0]

			if !combineModels {
				createObjectFile(outputDir, fileNameWithoutExt, objBuilder, mtlBuilder)
				objBuilder = modelextractor.NewOBJBuilder()
				mtlBuilder = modelextractor.NewMTLBuilder()
			}
		}

		if combineModels {
			createObjectFile(outputDir, outputName+"_combined", objBuilder, mtlBuilder)
		}

		gosucks.VAR(config)
	},
}

func mustCreateDir(outputDir string) {
	if combineModels {
		outputDir += "_combined"
	}

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
}

func createObjectFile(outputDir string, fileName string, objBuilder *modelextractor.OBJBuilder, mtlBuilder *modelextractor.MTLBuilder) {
	objFileName := fmt.Sprintf("%s.obj", fileName)
	mtlFileName := fmt.Sprintf("%s.mtl", fileName)

	objBuilder.WriteIncludeMTL(mtlFileName)

	fullOutputOBJFilePath := fmt.Sprintf("%s/%s", outputDir, objFileName)
	err := os.WriteFile(fullOutputOBJFilePath, []byte(objBuilder.String()), os.ModePerm)
	if err != nil {
		panic(err)
	}

	fullOutputMTLFilePath := fmt.Sprintf("%s/%s", outputDir, mtlFileName)
	err = os.WriteFile(fullOutputMTLFilePath, []byte(mtlBuilder.String()), os.ModePerm)

	for _, textureFilename := range mtlBuilder.TextureFilenames() {
		textureFilename = quoteRegex.ReplaceAllString(textureFilename, "")

		textureFullPath := filepath.Join(modelSourceDir, textureFilename)

		if _, err := os.Stat(textureFullPath); err != nil {
			if errors.Is(err, os.ErrNotExist) {
				fmt.Printf("could not find texture file %s\n", textureFullPath)
				continue
			} else {
				panic(fmt.Sprintf("could not stat texture file %s: %s\n", textureFullPath, err.Error()))
			}
		}

		textureData, err := os.ReadFile(textureFullPath)

		if err != nil {
			panic(err)
		}

		err = os.WriteFile(fmt.Sprintf("%s/%s", outputDir, textureFilename), textureData, 0755)

		if err != nil {
			panic(err)
		}
	}
}

func InitConvertCommand() {
	convertCommand.PersistentFlags().StringVarP(&modelOutputDir, "models-output-dir", "o", "", "-o D:\\ConvertedDRModels")
	convertCommand.PersistentFlags().StringVarP(&modelPrefix, "model-prefix", "p", "", "-p Townston_")
	convertCommand.PersistentFlags().StringVarP(&modelName, "model-name", "n", "", "-n ZonePortal")
	convertCommand.PersistentFlags().BoolVarP(&combineModels, "combine-models", "m", false, "-m")

	err := cobra.MarkFlagRequired(convertCommand.PersistentFlags(), "models-output-dir")

	if err != nil {
		panic(err)
	}
}
