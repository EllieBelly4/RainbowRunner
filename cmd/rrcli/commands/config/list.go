package config

import (
	"RainbowRunner/cmd/rrcli/commands/globals"
	"RainbowRunner/cmd/rrcli/configurator"
	"RainbowRunner/internal/types/drconfigtypes"
	"fmt"
	"github.com/goccy/go-json"
	"github.com/spf13/cobra"
	"regexp"
)

var listMaxDepth int
var listRegexpFilter string

var listCommand = &cobra.Command{
	Use:  "list",
	Args: cobra.ExactArgs(0),
	Run: func(cmd *cobra.Command, args []string) {
		config, err := configurator.LoadFromDumpedConfigFile(globals.InputFile)

		if err != nil {
			panic(err)
		}

		var regex *regexp.Regexp = nil

		if listRegexpFilter != "" {
			regex = regexp.MustCompile(listRegexpFilter)
		}

		class, err := config.List(listMaxDepth, func(group *drconfigtypes.DRClassChildGroup) bool {
			if regex != nil {
				return regex.Match([]byte(group.Name))
			}

			return true
		})

		if err != nil {
			panic(err)
		}

		jsonData, err := json.MarshalIndent(class, "", " ")

		if err != nil {
			return
		}

		fmt.Printf("%s\n", jsonData)
	},
}

func initListCommand() {
	listCommand.Flags().StringVarP(&listRegexpFilter, "regexp-filter", "x", "", "-x \"^[0-9]+\"")
	listCommand.Flags().IntVarP(&listMaxDepth, "list-depth", "d", 0, "-d 3")
}
