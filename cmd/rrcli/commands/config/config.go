package config

import (
	"RainbowRunner/cmd/rrcli/commands/config/category"
	"RainbowRunner/cmd/rrcli/commands/config/extract"
	"RainbowRunner/cmd/rrcli/commands/globals"
	"github.com/spf13/cobra"
)

var cmd = &cobra.Command{
	Use:   "config",
	Short: "Config commands",
}

func Init(rootCmd *cobra.Command) {
	cmd.PersistentFlags().StringVarP(&globals.InputFile, "input-config-file", "f", "resources/Dumps/generated/finalconf.json", "-f config\\finalconf.json")
	cmd.PersistentFlags().StringVarP(&globals.OutputFile, "output-file", "o", "", "-o dumps\\dr-class-output.json")

	initGetCommand()
	initListCommand()

	cmd.AddCommand(category.Init())
	cmd.AddCommand(extract.Init())

	cmd.AddCommand(getCommand)
	cmd.AddCommand(listCommand)

	rootCmd.AddCommand(cmd)
}
