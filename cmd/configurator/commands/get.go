package commands

import (
	"RainbowRunner/cmd/configurator/configurator"
	"encoding/json"
	"fmt"
	"github.com/spf13/cobra"
	"io/ioutil"
)

var inputFile string
var getOutputFile string

var getCommand = &cobra.Command{
	Use:  "get <fullgctype>",
	Args: cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		//world.test.world_questcells.data.NCI_Teleporter_02
		config, err := configurator.LoadFromDumpedConfigFile(inputFile)

		if err != nil {
			panic(err)
		}

		class, err := config.Get(args[0])

		if err != nil {
			panic(err)
		}

		jsonData, err := json.MarshalIndent(class, "", " ")

		if err != nil {
			return
		}

		if getOutputFile == "" {
			fmt.Printf("%s\n", jsonData)
		} else {
			err := ioutil.WriteFile(getOutputFile, jsonData, 0755)

			if err != nil {
				panic(err)
			}
		}
	},
}

func initGetCommand() {
	getCommand.Flags().StringVarP(&inputFile, "input-config-file", "f", "finalconf.json", "-f config\\finalconf.json")

	err := cobra.MarkFlagRequired(getCommand.Flags(), "input-config-file")

	getCommand.Flags().StringVarP(&getOutputFile, "output-file", "o", "", "-o dumps\\dr-class-output.json")

	if err != nil {
		panic(err)
	}
}
