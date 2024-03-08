package config

import (
	"os"
	"path"

	"github.com/replicapra/bartimaeus/util"
	"github.com/spf13/viper"
)

type Config struct {
	Hostname     string       `mapstructure:"hostname"`
	Database     Database     `mapstructure:"database"`
	Repositories []Repository `mapstructure:"repositories"`
}

type Repository struct {
	Path   string `mapstructure:"path"`
	Paused bool   `mapstructure:"paused"`
}

type Database struct {
	Path string `mapstructure:"path"`
}

// Init reads in config file and ENV variables if set.
func Init() {
	// find directory for config and create if it doesn't exists
	userConfigDir, err := os.UserConfigDir()
	util.CheckErr(err)
	configDir := path.Join(userConfigDir, "replicapra")
	err = os.MkdirAll(configDir, os.ModePerm)
	util.CheckErr(err)

	// set config file path
	viper.AddConfigPath(configDir)
	viper.SetConfigType("toml")
	viper.SetConfigName("bartimaeus")

	SetDefaults()
	Load()

}

func Load() {
	// read in environment variables that match
	viper.SetEnvPrefix("bartimaeus")
	viper.AutomaticEnv()

	// If a config file is found, read it in.
	err := viper.ReadInConfig()
	util.CheckErr(err)

	var repositories []Repository
	viper.UnmarshalKey("repositories", &repositories)
	viper.Set("repositories", repositories)

}

func Save() {
	viper.WriteConfigAs(viper.ConfigFileUsed())
}
