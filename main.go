package main

import (
	"go-short/dao/mysql"

	"github.com/spf13/viper"
)

func init() {
	viper.AddConfigPath("./config/default.yml")
}

func main() {
	mysql.Init()
	defer func() {
		db := mysql.GetConn()
		db.Close()
	}()
}
