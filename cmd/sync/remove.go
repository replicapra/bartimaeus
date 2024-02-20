package sync

import (
	"github.com/charmbracelet/log"
	"github.com/replicapra/bartimaeus/internal/config"
	"github.com/replicapra/bartimaeus/util"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	"golang.org/x/exp/slices"
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
	repositories := viper.Get("repositories").([]config.Repository)

	if !slices.ContainsFunc[[]config.Repository](repositories, func(repo config.Repository) bool { return repo.Path == absPath }) {
		log.Errorf("Repository %s not in list", absPath)
		return
	}

	index := slices.IndexFunc[[]config.Repository](repositories, func(repo config.Repository) bool { return repo.Path == absPath })
	repositories = append(repositories[:index], repositories[index+1:]...)

	viper.Set("repositories", repositories)

	config.Save()

	log.Infof("Repository %s removed", absPath)
}
