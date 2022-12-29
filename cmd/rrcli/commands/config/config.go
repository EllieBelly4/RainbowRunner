package config

import (
	"RainbowRunner/cmd/rrcli/commands/config/category"
	"RainbowRunner/cmd/rrcli/commands/globals"
	"github.com/spf13/cobra"
)

var cmd = &cobra.Command{
	Use:   "config",
	Short: "Config commands",
}

func Init(rootCmd *cobra.Command) {
	cmd.PersistentFlags().StringVarP(&globals.InputFile, "input-config-file", "f", "resources/Dumps/generated/finalconf.json", "-f config\\finalconf.json")
	cmd.PersistentFlags().StringVarP(&getOutputFile, "output-file", "o", "", "-o dumps\\dr-class-output.json")

	InitDumpCommand()
	InitGetCommand()
	InitCategoriseCommand()
	InitListCommand()

	cmd.AddCommand(listCommand)
	cmd.AddCommand(dumpCommand)
	cmd.AddCommand(getCommand)
	cmd.AddCommand(categoriseCommand)
	cmd.AddCommand(category.Init())

	rootCmd.AddCommand(cmd)
}
