package services

import (
	"go-short/commons"
	"go-short/dao/mysql"
	"go-short/dao/redisdao"
	"go-short/model"
	"time"
)

func AddShortUrl(url string, expire int64) error {
	su := model.ShortUrl{}
	su.URL = url
	su.Expire = expire
	su.Ctime = time.Now().Unix()
	su.Utime = time.Now().Unix()
	err := mysql.AddShortUrl(&su)
	if err != nil {
		return err
	}
	su.Keyword = commons.Encode(su.ID)
	kv := make(map[string]interface{})
	kv["keyword"] = su.Keyword
	err = mysql.UpdateShortUrl(&su, kv)
	if err == nil {
		err = redisdao.SetShortUrl(su.Keyword, su.URL)
	}
	return err
}

func UpdateShortUrl() {

}

func GetShortUrl() {

}

func DeleteShortUrl() {

}

func GetUrlByToken(token string) (string, error) {
	return redisdao.GetShortUrl(token)
}
