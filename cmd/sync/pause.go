package sync

import (
	"github.com/charmbracelet/log"
	"github.com/replicapra/bartimaeus/internal/database"
	"github.com/replicapra/bartimaeus/util"
	"gorm.io/gorm"

	"github.com/spf13/cobra"
)

// pauseCmd represents the pause command
var PauseCmd = &cobra.Command{
	Use:   "pause",
	Short: "A brief description of your command",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Args: cobra.MinimumNArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		for _, relPath := range args {
			PauseAbsPath(util.GetAbsPath(relPath))
		}
	},
}

func PauseAbsPath(absPath string) {
	repository, error := database.GetRepositoryByAbsPath(absPath)
	if error == gorm.ErrRecordNotFound {
		log.Errorf("Repository %s not in list", absPath)
		return
	}
	util.CheckErr(error)

	repository.Paused = true

	database.Client.Save(&repository)

	log.Infof("Repository %s paused", absPath)
}
