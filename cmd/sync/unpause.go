/*
Copyright © 2024 replicapra
*/
package sync

import (
	"path/filepath"

	"github.com/charmbracelet/log"
	"github.com/replicapra/bartimaeus/internal/config"
	"github.com/replicapra/bartimaeus/util"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	"golang.org/x/exp/slices"
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
			unpauseRelPath(relPath)
		}

	},
}

func unpauseRelPath(relPath string) {
	absPath, err := filepath.Abs(relPath)
	util.CheckErr(err)

	repositories := viper.Get("repositories").([]config.Repository)

	if !slices.ContainsFunc[[]config.Repository](repositories, func(repo config.Repository) bool { return repo.Path == absPath }) {
		log.Errorf("Repository %s not in list", absPath)
		return
	}

	index := slices.IndexFunc[[]config.Repository](repositories, func(repo config.Repository) bool { return repo.Path == absPath })
	repositories[index].Paused = false

	viper.Set("repositories", repositories)

	config.Save()

	log.Infof("Repository %s unpaused", absPath)
}