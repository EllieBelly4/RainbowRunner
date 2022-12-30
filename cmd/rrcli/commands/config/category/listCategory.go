package category

import (
	"RainbowRunner/cmd/rrcli/configurator"
	"RainbowRunner/internal/gosucks"
	"RainbowRunner/internal/types/drconfigtypes"
	"fmt"
	"github.com/spf13/cobra"
	"strings"
)

var categoryPrefix string
var listCategoryDepth int

var listCategoryCommand = &cobra.Command{
	Use:  "list",
	Args: cobra.ExactArgs(0),
	Run: func(cmd *cobra.Command, args []string) {
		categoryConfig, err := configurator.LoadFromCategoryConfigFile(categoryInputFile)

		if err != nil {
			panic(err)
		}

		splitPrefix := strings.Split(categoryPrefix, ".")

		if categoryPrefix == "" {
			splitPrefix = []string{}
		}

		dumpCategories(categoryConfig, splitPrefix, 0)
	},
}

func dumpCategories(config map[string]*drconfigtypes.DRCategory, prefix []string, depth int) {
	if depth > listCategoryDepth && listCategoryDepth >= 0 {
		return
	}

	if len(prefix) == 0 {
		depth++
	}

	depthPrefix := ""

	for i := 0; i < depth-1; i++ {
		depthPrefix += "  "
	}

	for name, categoryConfig := range config {
		newPrefix := prefix

		if len(prefix) == 0 {
			fmt.Printf("%s%s\n", depthPrefix, name)
		} else {
			if prefix[0] != name {
				continue
			}

			newPrefix = prefix[1:]
		}

		dumpCategories(categoryConfig.Children, newPrefix, depth)
		gosucks.VAR(categoryConfig)
	}
}

func initListCategoryCommand() {
	listCategoryCommand.Flags().StringVarP(&categoryPrefix, "category-prefix", "p", "", "-p armor.basearmor")
	listCategoryCommand.Flags().IntVarP(&listCategoryDepth, "list-depth", "d", 0, "-d 3")
	listCategoryCommand.PersistentFlags().StringVarP(&categoryInputFile, "category-input-config-file", "c", "resources/Dumps/generated/drcategories.json", "-f resources/Dumps/generated/drcategories.json")
}
