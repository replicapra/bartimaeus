package sync

import (
	"github.com/charmbracelet/log"
	"github.com/replicapra/bartimaeus/internal/database"
	"github.com/replicapra/bartimaeus/util"
	"github.com/spf13/cobra"
	"gorm.io/gorm"
)

// unpauseCmd represents the unpause command
var UnpauseCmd = &cobra.Command{
	Use:   "unpause",
	Short: "A brief description of your command",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Args: cobra.MinimumNArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		for _, relPath := range args {
			UnpauseAbsPath(util.GetAbsPath(relPath))
		}
	},
}

func UnpauseAbsPath(absPath string) {
	repository, error := database.GetRepositoryByAbsPath(absPath)
	if error == gorm.ErrRecordNotFound {
		log.Errorf("Repository %s not in list", absPath)
		return
	}
	util.CheckErr(error)

	repository.Paused = false

	database.Client.Save(&repository)

	log.Infof("Repository %s unpaused", absPath)
}
