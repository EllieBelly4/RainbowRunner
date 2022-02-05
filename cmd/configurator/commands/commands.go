package commands

import (
	"fmt"
	"github.com/spf13/cobra"
	"os"
)

var rootDir string

var rootCmd = &cobra.Command{
	Use: "configurator",
}

func Init() {
	rootCmd.PersistentFlags().StringVarP(&rootDir, "config-root-dir", "d", "", "-d D:\\Work\\DungeonRunners\\ConfigFiles")

	err := cobra.MarkFlagRequired(rootCmd.PersistentFlags(), "config-root-dir")

	if err != nil {
		panic(err)
	}

	initDumpCommand()
	rootCmd.AddCommand(dumpCommand)
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}
