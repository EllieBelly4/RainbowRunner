package commands

import (
	"RainbowRunner/cmd/rrcli/commands/config"
	"fmt"
	"github.com/spf13/cobra"
	"os"
)

var rootCmd = &cobra.Command{
	Use: "rrcli",
}

func Init() {
	config.Init(rootCmd)
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}
