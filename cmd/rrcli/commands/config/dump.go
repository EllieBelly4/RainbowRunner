package config

import (
	"RainbowRunner/cmd/configparser/configparser"
	"RainbowRunner/cmd/rrcli/configurator"
	"encoding/json"
	"github.com/spf13/cobra"
	"io/ioutil"
)

var outputFile string
var rootDir string
var dumpFileName string

var dumpCommand = &cobra.Command{
	Use: "dump",
	Run: func(cmd *cobra.Command, args []string) {
		var config *configparser.DRConfig
		var err error

		if dumpFileName == "" {
			if rootDir == "" {
				panic("must provide config-root-dir when not specifying an input file path")
			}

			config, err = configurator.LoadAllConfigurationFiles(rootDir)
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

		err = ioutil.WriteFile(outputFile, data, 0774)

		if err != nil {
			panic(err)
		}
	},
}

func InitDumpCommand() {
	dumpCommand.PersistentFlags().StringVarP(&rootDir, "config-root-dir", "d", "", "-d D:\\Work\\DungeonRunners\\ConfigFiles")

	//err := cobra.MarkFlagRequired(dumpCommand.PersistentFlags(), "config-root-dir")
	//
	//if err != nil {
	//	panic(err)
	//}

	dumpCommand.Flags().StringVarP(&outputFile, "output-file", "o", "dr-config-dump.json", "-o Dumps\\dump.json")
	dumpCommand.Flags().StringVarP(&dumpFileName, "input-file", "i", "", "-i Config_File.txt")
}
