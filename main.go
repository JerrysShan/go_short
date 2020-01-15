package main

import (
	"fmt"
	"go-short/dao/mysql"
	"go-short/dao/redisdao"
	"go-short/handlers"

	"github.com/gin-gonic/gin"

	"github.com/spf13/viper"
)

func init() {
	viper.SetConfigName("default.yml")
	viper.AddConfigPath("./config")
	err := viper.ReadInConfig()
	if err != nil {
		panic(fmt.Errorf("Fatal error config file :%w", err))
	}
}

func main() {
	mysql.Init()
	redisdao.Init()
	defer func() {
		db := mysql.GetConn()
		db.Close()
	}()
	app := gin.Default()
	app.GET("/api/short/url", handlers.GetShortUrl)
	app.POST("/api/short/url", handlers.AddShortUrl)
	app.PUT("/api/short/url", handlers.UpdateShortUrl)
	app.DELETE("/api/short/url", handlers.DeleteShortUrl)
	app.GET("/t/:token", handlers.Redirect)
	app.Run(":8090")
}
