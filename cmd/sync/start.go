package sync

import (
	"fmt"
	"os"
	"slices"
	"time"

	"github.com/charmbracelet/log"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"

	"github.com/replicapra/bartimaeus/internal/command"
	"github.com/replicapra/bartimaeus/internal/config"
)

// syncCmd represents the sync command
var StartCmd = &cobra.Command{
	Use:   "start",
	Short: "Starts syncing your repositories to github.",
	Long:  "Iterates over all repositories defined in your config file and syncs them to github.",
	Run: func(cmd *cobra.Command, args []string) {
		repositories := viper.Get("repositories").([]config.Repository)
		for _, repo := range repositories {
			syncRepo(repo)
		}
	},
}

func init() {
	StartCmd.Flags().Bool("force", false, "Force syncing of paused repositories")
	viper.BindPFlag("force", StartCmd.Flags().Lookup("force"))
}

func syncRepo(repo config.Repository) {

	if err := os.Chdir(repo.Path); err != nil {
		log.Errorf("%s doesn't exist", repo.Path)
		return
	}

	if repo.Paused && !viper.GetBool("force") {
		log.Infof("Skipping %s\n", repo.Path)
		return
	}

	commitMsg := fmt.Sprintf("Sync: %s | %s", viper.GetString("hostname"), time.Now().UTC().String())

	if status, err := command.Run("git", "add", "."); err != nil {
		log.Errorf("Error executing 'git add': %s\n", err)
		log.Info("", "status", status)
		if !slices.Contains([]int{}, status) {
			return
		}

	}

	if status, err := command.Run("git", "commit", "-m", commitMsg); err != nil {
		log.Errorf("Error executing 'git commit': %s\n", err)
		log.Info("", "status", status)
		if !slices.Contains([]int{1}, status) {
			return
		}
	}

	if status, err := command.Run("git", "pull"); err != nil {
		log.Errorf("Error executing 'git pull': %s\n", err)
		log.Info("", "status", status)
		if !slices.Contains([]int{}, status) {
			return
		}
	}

	if status, err := command.Run("git", "push"); err != nil {
		log.Errorf("Error executing 'git push': %s\n", err)
		log.Info("", "status", status)
		if !slices.Contains([]int{}, status) {
			return
		}
	}

	log.Infof("Synced %s\n", repo.Path)
}
