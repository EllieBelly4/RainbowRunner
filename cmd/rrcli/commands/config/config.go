package config

import "github.com/spf13/cobra"

func Init(rootCmd *cobra.Command) {
	InitDumpCommand()
	InitGetCommand()
	InitCategoriseCommand()

	rootCmd.AddCommand(dumpCommand)
	rootCmd.AddCommand(getCommand)
	rootCmd.AddCommand(categoriseCommand)
}
