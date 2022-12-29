package extract

import (
	"RainbowRunner/cmd/rrcli/commands/globals"
	"RainbowRunner/cmd/rrcli/configurator"
	"fmt"
	"github.com/goccy/go-json"
	"github.com/spf13/cobra"
	"io/ioutil"
)

var categoriseCommand = &cobra.Command{
	Use: "categorise",
	Run: func(cmd *cobra.Command, args []string) {
		config, err := configurator.LoadFromDumpedConfigFile(globals.InputFile)

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

		if globals.OutputFile == "" {
			fmt.Printf("%s\n", jsonData)
		} else {
			err := ioutil.WriteFile(globals.OutputFile, jsonData, 0755)

			if err != nil {
				panic(err)
			}
		}
	},
}

func initCategoriseCommand() {
}
