package mysql

import (
	"github.com/jinzhu/gorm"
)

var db *gorm.DB

// Init build connection
func Init() {
	var err error
	db, err = gorm.Open("")
	if err != nil {
		panic(err)
	}
}

// GetConn return connection
func GetConn() *gorm.DB {
	return db
}
