package database

import (
	"github.com/replicapra/bartimaeus/util"
	"github.com/spf13/viper"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

var Client *gorm.DB

func Init() {
	client, err := gorm.Open(sqlite.Open(viper.GetString("database.path")), &gorm.Config{})
	util.CheckErr(err)

	client.AutoMigrate(&Repository{})

	Client = client
}
