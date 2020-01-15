package handlers

import (
	"go-short/services"
	"strconv"

	"github.com/gin-gonic/gin"
)

func AddShortUrl(c *gin.Context) {
	url := c.PostForm("url")
	expire, err := strconv.ParseInt(c.PostForm("expire"), 10, 64)
	if err != nil {
		c.JSON(200, gin.H{
			"code":    -1,
			"message": "params is invalid",
		})
		return
	}
	err = services.AddShortUrl(url, expire)
	if err != nil {
		c.JSON(200, gin.H{
			"code":    -1,
			"message": err.Error(),
		})
		return
	}
	c.JSON(200, gin.H{
		"code":    0,
		"message": "OK",
	})
}

func UpdateShortUrl(c *gin.Context) {

}

func GetShortUrl(c *gin.Context) {

}

func DeleteShortUrl(c *gin.Context) {

}

func Redirect(c *gin.Context) {
	token := c.Param("token")
	if token == "" {
		c.JSON(200, gin.H{"code": -1, "message": "param is invalid"})
		return
	}
	url, err := services.GetUrlByToken(token)
	if err != nil {
		c.JSON(200, gin.H{"code": -1, "message": err.Error()})
		return
	}
	c.Redirect(302, url)
}
