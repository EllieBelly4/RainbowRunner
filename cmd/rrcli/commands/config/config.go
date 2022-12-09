package config

import "github.com/spf13/cobra"

func Init(rootCmd *cobra.Command) {
	InitDumpCommand()
	InitGetCommand()
	InitCategoriseCommand()
	InitListCategoryCommand()
	InitListCommand()

	rootCmd.AddCommand(listCommand)
	rootCmd.AddCommand(dumpCommand)
	rootCmd.AddCommand(getCommand)
	rootCmd.AddCommand(categoriseCommand)
	rootCmd.AddCommand(listCategoryCommand)
}
