package commands

import (
	"RainbowRunner/cmd/configurator/configurator"
	"encoding/json"
	"github.com/spf13/cobra"
	"io/ioutil"
)

var dumpCommand = &cobra.Command{
	Use: "dump",
	Run: func(cmd *cobra.Command, args []string) {
		config, err := configurator.LoadAllConfigurationFiles(rootDir)

		if err != nil {
			panic(err)
		}

		data, err := json.MarshalIndent(config, "", "  ")

		if err != nil {
			panic(err)
		}

		err = ioutil.WriteFile("resources/Dumps/generated/test.json", data, 0774)

		if err != nil {
			panic(err)
		}
	},
}
