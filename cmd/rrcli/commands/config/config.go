package config

import (
	"RainbowRunner/cmd/rrcli/commands/config/category"
	"RainbowRunner/cmd/rrcli/commands/config/gc"
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

	InitDumpCommand()
	InitCategoriseCommand()

	cmd.AddCommand(gc.Init())
	cmd.AddCommand(category.Init())
	cmd.AddCommand(dumpCommand)
	cmd.AddCommand(categoriseCommand)

	rootCmd.AddCommand(cmd)
}
