package sync

import (
	"os"
	"path"

	"github.com/charmbracelet/log"
	"github.com/replicapra/bartimaeus/internal/database"
	"github.com/replicapra/bartimaeus/util"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	"gorm.io/gorm/clause"
)

// addCmd represents the add command
var AddCmd = &cobra.Command{
	Use:   "add",
	Short: "Adds the given repositories to the list of repositories to sync.",
	Long:  `usage: bartimaeus sync add <relative paths>`,
	Args:  cobra.MinimumNArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		for _, relPath := range args {
			AddAbsPath(util.GetAbsPath(relPath))
		}
	},
}

func init() {
	AddCmd.Flags().BoolP("paused", "p", false, "Add the repository in paused state")
	viper.BindPFlag("paused", AddCmd.Flags().Lookup("paused"))
}

func AddAbsPath(absPath string) {
	if _, err := os.Stat(path.Join(absPath, ".git")); os.IsNotExist(err) {
		log.Errorf("%s is not root of a git directory", absPath)
		return
	}

	database.Client.Clauses(clause.OnConflict{
		Columns:   []clause.Column{{Name: "path"}},
		DoUpdates: clause.AssignmentColumns([]string{"paused"}),
	}).Create(&database.Repository{
		Path:   absPath,
		Paused: viper.GetBool("paused"),
	})

	log.Infof("Repository %s added", absPath)
}
