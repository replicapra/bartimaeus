package cmd

import (
	"github.com/replicapra/bartimaeus/cmd/sync"
	"github.com/spf13/cobra"
)

// syncCmd represents the sync command
var SyncCmd = &cobra.Command{
	Use:   "sync",
	Short: "A brief description of your command",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
}

func init() {
	SyncCmd.AddCommand(sync.StartCmd)
	SyncCmd.AddCommand(sync.AddCmd)
	SyncCmd.AddCommand(sync.PauseCmd)
	SyncCmd.AddCommand(sync.RemoveCmd)
	SyncCmd.AddCommand(sync.UnpauseCmd)
	SyncCmd.AddCommand(sync.ListCmd)
}
