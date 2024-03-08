package sync

import (
	"github.com/charmbracelet/log"
	"github.com/replicapra/bartimaeus/internal/database"
	"github.com/replicapra/bartimaeus/util"
	"github.com/spf13/cobra"
)

// removeCmd represents the remove command
var RemoveCmd = &cobra.Command{
	Use:   "remove",
	Short: "A brief description of your command",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Args: cobra.MinimumNArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		for _, relPath := range args {
			RemoveAbsPath(util.GetAbsPath(relPath))
		}
	},
}

func RemoveAbsPath(absPath string) {
	database.Client.Delete(&database.Repository{}, "path = ?", absPath)

	log.Infof("Repository %s removed", absPath)
}
