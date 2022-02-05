package commands

import (
	"RainbowRunner/cmd/configurator/configurator"
	"encoding/json"
	"github.com/spf13/cobra"
	"io/ioutil"
)

var outputFile string
var rootDir string

var dumpCommand = &cobra.Command{
	Use: "dump",
	Run: func(cmd *cobra.Command, args []string) {
		config, err := configurator.LoadAllConfigurationFiles(rootDir)

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

func initDumpCommand() {
	dumpCommand.PersistentFlags().StringVarP(&rootDir, "config-root-dir", "d", "", "-d D:\\Work\\DungeonRunners\\ConfigFiles")

	err := cobra.MarkFlagRequired(dumpCommand.PersistentFlags(), "config-root-dir")

	if err != nil {
		panic(err)
	}

	dumpCommand.Flags().StringVarP(&outputFile, "output-file", "o", "dr-config-dump.json", "-o Dumps\\dump.json")
}
