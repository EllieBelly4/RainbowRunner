package config

import (
	"RainbowRunner/cmd/rrcli/configurator"
	"encoding/json"
	"fmt"
	"github.com/spf13/cobra"
	"io/ioutil"
)

var categoriseinputFile string
var categoriseOutputFile string

var categoriseCommand = &cobra.Command{
	Use: "categorise",
	Run: func(cmd *cobra.Command, args []string) {
		//world.test.world_questcells.data.NCI_Teleporter_02
		config, err := configurator.LoadFromDumpedConfigFile(categoriseinputFile)

		if err != nil {
			panic(err)
		}

		categories, err := config.GenerateCategoryMap()

		if err != nil {
			panic(err)
		}

		jsonData, err := json.MarshalIndent(categories, "", " ")

		if err != nil {
			return
		}

		if categoriseOutputFile == "" {
			fmt.Printf("%s\n", jsonData)
		} else {
			err := ioutil.WriteFile(categoriseOutputFile, jsonData, 0755)

			if err != nil {
				panic(err)
			}
		}
	},
}

func InitCategoriseCommand() {
	categoriseCommand.Flags().StringVarP(&categoriseinputFile, "input-config-file", "f", "finalconf.json", "-f config\\finalconf.json")

	err := cobra.MarkFlagRequired(categoriseCommand.Flags(), "input-config-file")

	categoriseCommand.Flags().StringVarP(&categoriseOutputFile, "output-file", "o", "", "-o dumps\\dr-categories-output.json")

	if err != nil {
		panic(err)
	}
}
