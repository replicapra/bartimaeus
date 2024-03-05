package open

import (
	"fmt"
	"os"
	"os/exec"
	"path"
	"time"

	"github.com/replicapra/bartimaeus/internal/constants"
	"github.com/replicapra/bartimaeus/util"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

// inboxCmd represents the inbox command
var InboxCmd = &cobra.Command{
	Use:   "inbox",
	Short: "Open todays inbox file in neovim.",
	Long:  "Open todays inbox file in neovim.",
	Run: func(cmd *cobra.Command, args []string) {
		obsidianDirPath := viper.GetString("obsidian.path")
		currentInboxFilePath := path.Join(obsidianDirPath, constants.InboxSpacePath, fmt.Sprintf("%s.md", time.Now().UTC().Format("2006-01-02")))

		command := exec.Command("nvim", currentInboxFilePath)
		command.Stdin = os.Stdin
		command.Stdout = os.Stdout
		err := command.Run()
		util.CheckErr(err)
	},
}
