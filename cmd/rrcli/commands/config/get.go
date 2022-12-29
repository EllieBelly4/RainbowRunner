package config

import (
	"RainbowRunner/cmd/rrcli/commands/globals"
	"RainbowRunner/cmd/rrcli/configurator"
	"fmt"
	"github.com/goccy/go-json"
	"github.com/spf13/cobra"
	"io/ioutil"
)

var inputFile string
var getOutputFile string

var getCommand = &cobra.Command{
	Use:  "get <fullgctype>",
	Args: cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		config, err := configurator.LoadFromDumpedConfigFile(globals.InputFile)

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

}
