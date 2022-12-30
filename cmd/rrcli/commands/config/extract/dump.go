package extract

import (
	"RainbowRunner/cmd/configparser/configparser"
	"RainbowRunner/cmd/rrcli/configurator"
	"RainbowRunner/internal/types/drconfigtypes"
	"compress/zlib"
	"github.com/goccy/go-json"
	"github.com/spf13/cobra"
	"os"
)

var outputFile string
var rootDir string
var dumpFileName string
var dumpAllowedExtensions []string

var dumpCommand = &cobra.Command{
	Use: "dump",
	Run: func(cmd *cobra.Command, args []string) {
		var config *drconfigtypes.DRConfig
		var err error

		if dumpFileName == "" {
			if rootDir == "" {
				panic("must provide config-root-dir when not specifying an input file path")
			}

			config, err = configurator.LoadAllConfigurationFiles(rootDir, dumpAllowedExtensions)
		} else {
			config, err = configparser.ParseAllFilesToDRConfig(
				[]string{dumpFileName},
				rootDir,
			)
		}

		if err != nil {
			panic(err)
		}

		//data, err := json.MarshalIndent(config, "", " ")
		data, err := json.Marshal(config)

		if err != nil {
			panic(err)
		}

		fp, err := os.OpenFile(outputFile, os.O_CREATE|os.O_WRONLY|os.O_TRUNC, 0774)

		if err != nil {
			panic(err)
		}

		defer fp.Close()

		w := zlib.NewWriter(fp)
		_, err = w.Write(data)

		defer w.Close()

		if err != nil {
			panic(err)
		}
	},
}

func initDumpCommand() {
	dumpCommand.PersistentFlags().StringVarP(&rootDir, "config-root-dir", "d", "", "-d D:\\Work\\DungeonRunners\\ConfigFiles")
	dumpCommand.Flags().StringArrayVarP(&dumpAllowedExtensions, "allowed-extensions", "e", []string{".gc", ".txt"}, "-e .gc -e .txt")

	//err := cobra.MarkFlagRequired(dumpCommand.PersistentFlags(), "config-root-dir")
	//
	//if err != nil {
	//	panic(err)
	//}

	dumpCommand.Flags().StringVarP(&outputFile, "output-file", "o", "dr-config-dump.json", "-o Dumps\\dump.json")
	dumpCommand.Flags().StringVarP(&dumpFileName, "input-file", "i", "", "-i Config_File.txt")
}
