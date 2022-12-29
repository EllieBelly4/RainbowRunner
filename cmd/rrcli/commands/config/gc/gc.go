package gc

import "github.com/spf13/cobra"

var cmd = &cobra.Command{
	Use:   "gc",
	Short: "GC commands",
}

func Init() *cobra.Command {
	initGetCommand()
	initListCommand()

	cmd.AddCommand(getCommand)
	cmd.AddCommand(listCommand)

	return cmd
}
