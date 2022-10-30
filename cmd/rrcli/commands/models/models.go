package models

import "github.com/spf13/cobra"

var configFile string
var modelSourceDir string

var modelsCommand = &cobra.Command{
	Use: "models",
}

func Init(rootCmd *cobra.Command) {
	modelsCommand.PersistentFlags().StringVarP(&configFile, "input-config-file", "f", "resources/Dumps/generated/finalconf.json", "-f config\\finalconf.json")
	modelsCommand.PersistentFlags().StringVarP(&modelSourceDir, "models-source-dir", "d", "", "-d C:\\DRExtracted3DNodes")

	err := cobra.MarkFlagRequired(modelsCommand.PersistentFlags(), "models-source-dir")

	if err != nil {
		panic(err)
	}

	rootCmd.AddCommand(modelsCommand)

	InitConvertCommand()
	modelsCommand.AddCommand(convertCommand)
}
