package mysql

import (
	"go-short/model"
)

func AddShortUrl(sr *model.ShortUrl) error {
	err := db.Create(sr).Error
	return err
}

func UpdateShortUrl(sr *model.ShortUrl, kv map[string]interface{}) error {
	return db.Model(sr).Update(kv).Error
}

func GetShortUrl(id int64, keyword string) {

}

func DeleteShortUrl() {

}
