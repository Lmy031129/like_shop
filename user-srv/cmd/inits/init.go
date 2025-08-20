package inits

import "user-srv/cmd/config"

func init() {
	config.InitViper()
	InitMysql()
	InitRedis()
}
