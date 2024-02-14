package config

import (
	"os"

	"github.com/replicapra/bartimaeus/util"
	"github.com/spf13/viper"
)

func SetDefaults() {
	hostname, err := os.Hostname()
	util.CheckErr(err)
	viper.SetDefault("hostname", hostname)
	viper.SetDefault("repositories", []Repository{{Path: "/home/user/absolute/path/to/repository", Paused: true}})

	// write config file if it doesn't exists
	viper.SafeWriteConfig()
}
