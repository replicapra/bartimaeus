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

	// write config file if it doesn't exists
	viper.SafeWriteConfig()
}
