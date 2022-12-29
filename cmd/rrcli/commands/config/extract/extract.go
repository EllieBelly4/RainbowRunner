package extract

import "github.com/spf13/cobra"

var cmd = &cobra.Command{
	Use:   "extract",
	Short: "Extract commands",
}

func Init() *cobra.Command {
	initCategoriseCommand()
	initDumpCommand()

	cmd.AddCommand(categoriseCommand)
	cmd.AddCommand(dumpCommand)

	return cmd
}
