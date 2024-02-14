/*
Copyright © 2024 replicapra
*/
package sync

import (
	"os"
	"path"
	"path/filepath"

	"github.com/charmbracelet/log"
	"github.com/replicapra/bartimaeus/internal/config"
	"github.com/replicapra/bartimaeus/util"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	"golang.org/x/exp/slices"
)

// addCmd represents the add command
var AddCmd = &cobra.Command{
	Use:   "add",
	Short: "A brief description of your command",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Args: cobra.MinimumNArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		for _, relPath := range args {
			addRelPath(relPath)
		}
	},
}

func init() {
	AddCmd.Flags().BoolP("paused", "p", false, "Add the repository in paused state")
	viper.BindPFlag("paused", AddCmd.Flags().Lookup("paused"))
}

func addRelPath(relPath string) {
	absPath, err := filepath.Abs(relPath)
	util.CheckErr(err)

	repositories := viper.Get("repositories").([]config.Repository)

	if slices.ContainsFunc[[]config.Repository](repositories, func(repo config.Repository) bool { return repo.Path == absPath }) {
		log.Errorf("Repository %s already in list", absPath)
		return
	}

	if _, err := os.Stat(path.Join(absPath, ".git")); os.IsNotExist(err) {
		log.Errorf("%s is not root of a git directory", absPath)
		return
	}

	repositories = append(repositories, config.Repository{
		Path:   absPath,
		Paused: viper.GetBool("paused"),
	})

	viper.Set("repositories", repositories)

	config.Save()

	log.Infof("Repository %s added", absPath)
}
