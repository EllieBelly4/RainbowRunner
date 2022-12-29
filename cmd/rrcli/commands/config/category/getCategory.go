package category

import (
	"RainbowRunner/cmd/configparser/configparser"
	"RainbowRunner/cmd/rrcli/commands/globals"
	"RainbowRunner/cmd/rrcli/configurator"
	"RainbowRunner/internal/types/configtypes"
	"encoding/json"
	"fmt"
	"github.com/spf13/cobra"
	"io/ioutil"
	"regexp"
)

var categoryInputFile string
var categoryMinDepth int
var regexpFilter string
var getOutputFile string

var getCategoryCommand = &cobra.Command{
	Use:  "get [category]",
	Args: cobra.MinimumNArgs(0),
	Run: func(cmd *cobra.Command, args []string) {
		categoryConfig, err := configurator.LoadFromCategoryConfigFile(categoryInputFile)

		if err != nil {
			panic(err)
		}

		regex := regexp.MustCompile(regexpFilter)

		category := ""

		if len(args) > 0 {
			category = args[0]
		}

		gcTypes, err := configparser.GetGCTypesByCategory(
			category,
			categoryConfig,
			categoryMinDepth,
			regex,
		)

		if err != nil {
			panic(err)
		}

		config, err := configurator.LoadFromDumpedConfigFile(globals.InputFile)

		if err != nil {
			panic(err)
		}

		var results []*configtypes.DRClassChildGroup

		for _, gcType := range gcTypes {
			class, err := config.Get(gcType)

			if err != nil {
				panic(err)
			}

			results = append(results, class...)
		}

		jsonData, err := json.MarshalIndent(results, "", " ")

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

func initGetCategoryCommand() {
	getCategoryCommand.PersistentFlags().StringVarP(&categoryInputFile, "category-input-config-file", "c", "resources/Dumps/generated/drcategories.json", "-f resources/Dumps/generated/drcategories.json")
	getCategoryCommand.PersistentFlags().StringVarP(&getOutputFile, "output-file", "o", "", "-o tmp\\output.json")

	getCategoryCommand.Flags().IntVarP(&categoryMinDepth, "min-depth", "m", -1, "-m 5")
	getCategoryCommand.Flags().StringVarP(&regexpFilter, "regexp-filter", "x", "", "-x \"^[0-9]+\"")
}
