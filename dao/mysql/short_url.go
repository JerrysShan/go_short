package mysql

import (
	"go-short/model"
)

func CreateShortUrl(sr *model.ShortUrl) bool {
	return db.NewRecord(*sr)
}

func UpdateShortUrl() {

}

func GetShortUrl() {

}

func DeleteShortUrl() {

}
