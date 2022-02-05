package commands

import (
	"RainbowRunner/cmd/configparser/configparser"
	"RainbowRunner/cmd/configurator/configurator"
	"RainbowRunner/internal/database"
	"encoding/json"
	"fmt"
	"github.com/spf13/cobra"
	"io/ioutil"
	"regexp"
)

var categoryInputFile string
var minDepth int
var regexpFilter string

var getCategoryCommand = &cobra.Command{
	Use:  "category <category>",
	Args: cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		categoryConfig, err := configurator.LoadFromCategoryConfigFile(categoryInputFile)

		if err != nil {
			panic(err)
		}

		regex := regexp.MustCompile(regexpFilter)

		gcTypes, err := configparser.GetGCTypesByCategory(
			args[0],
			categoryConfig,
			minDepth,
			regex,
		)

		if err != nil {
			panic(err)
		}

		config, err := configurator.LoadFromDumpedConfigFile(inputFile)

		if err != nil {
			panic(err)
		}

		var results []*database.DRClassChildGroup

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
	getCategoryCommand.PersistentFlags().StringVarP(&categoryInputFile, "category-input-config-file", "c", "drcategories.json", "-f config\\drcategories.json")

	err := cobra.MarkFlagRequired(getCategoryCommand.PersistentFlags(), "category-input-config-file")

	if err != nil {
		panic(err)
	}

	getCategoryCommand.Flags().IntVarP(&minDepth, "min-depth", "m", -1, "-m 5")
	getCategoryCommand.Flags().StringVarP(&regexpFilter, "regexp-filter", "x", "", "-x \"^[0-9]+\"")
}
