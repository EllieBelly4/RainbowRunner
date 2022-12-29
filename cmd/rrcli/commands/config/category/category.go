package category

import "github.com/spf13/cobra"

var cmd = &cobra.Command{
	Use:   "category",
	Short: "Category commands",
}

func Init() *cobra.Command {
	initListCategoryCommand()
	initGetCategoryCommand()

	cmd.AddCommand(listCategoryCommand)
	cmd.AddCommand(getCategoryCommand)

	return cmd
}
