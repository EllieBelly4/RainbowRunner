package commands

import (
	"fmt"
	"github.com/spf13/cobra"
	"os"
)

var rootCmd = &cobra.Command{
	Use: "configurator",
}

func Init() {
	initDumpCommand()
	initGetCommand()

	rootCmd.AddCommand(dumpCommand)
	rootCmd.AddCommand(getCommand)
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}
