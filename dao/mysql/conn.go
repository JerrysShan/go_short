package mysql

import (
	"fmt"

	_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
	"github.com/spf13/viper"
)

var db *gorm.DB

// Init build connection
func Init() {
	var err error
	host := viper.GetString("mysql.host")
	database := viper.GetString("mysql.database")
	user := viper.GetString("mysql.user")
	pwd := viper.GetString("mysql.pwd")
	db, err = gorm.Open("mysql", fmt.Sprintf("%s:%s@(%s)/%s?charset=utf8&parseTime=True&loc=Local", user, pwd, host, database))
	if err != nil {
		panic(err)
	}
}

// GetConn return connection
func GetConn() *gorm.DB {
	return db
}
