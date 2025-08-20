package config

import (
	"log"

	"github.com/spf13/viper"
)

type Config struct {
	Mysql
	Redis
}
type Mysql struct {
	User     string
	Host     string
	Port     int
	Password string
	DataBase string
}
type Redis struct {
	Host     string
	Port     int
	Password string
}

var ConfigAppData Config

func InitViper() {
	viper.SetConfigFile("../../cmd/config/dev.yaml")
	viper.ReadInConfig()
	viper.Unmarshal(&ConfigAppData)
	log.Println("viper is success", ConfigAppData)
}
