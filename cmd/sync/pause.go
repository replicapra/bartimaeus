package sync

import (
	"github.com/charmbracelet/log"
	"github.com/replicapra/bartimaeus/internal/config"
	"github.com/replicapra/bartimaeus/util"
	"golang.org/x/exp/slices"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"
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
	repositories := viper.Get("repositories").([]config.Repository)

	if !slices.ContainsFunc[[]config.Repository](repositories, func(repo config.Repository) bool { return repo.Path == absPath }) {
		log.Errorf("Repository %s not in list", absPath)
		return
	}

	index := slices.IndexFunc[[]config.Repository](repositories, func(repo config.Repository) bool { return repo.Path == absPath })
	repositories[index].Paused = true

	viper.Set("repositories", repositories)

	config.Save()

	log.Infof("Repository %s paused", absPath)
}
