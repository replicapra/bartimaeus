package obsidian

import (
	"github.com/replicapra/bartimaeus/cmd/obsidian/open"
	"github.com/spf13/cobra"
)

// openCmd represents the open command
var OpenCmd = &cobra.Command{
	Use:   "open",
	Short: "A brief description of your command",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
}

func init() {
	OpenCmd.AddCommand(open.InboxCmd)
}
