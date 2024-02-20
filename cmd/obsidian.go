package cmd

import (
	"github.com/charmbracelet/log"
	"github.com/replicapra/bartimaeus/cmd/obsidian"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

// obsidianCmd represents the obsidian command
var ObsidianCmd = &cobra.Command{
	Use:   "obsidian",
	Short: "A brief description of your command",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	PersistentPreRun: func(cmd *cobra.Command, args []string) {
		obsidianDirPath := viper.GetString("obsidian.path")
		if obsidianDirPath == "" {
			log.Fatal("Obsidian path not set")
		}
	},
}

func init() {
	ObsidianCmd.AddCommand(obsidian.AddCmd)
}
