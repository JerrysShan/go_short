package model

type ShortUrl struct {
	ID      int64
	URL     string `gorm:"column:url;"`
	Keyword string `gorm:"column:keyword;"`
	Ctime   int64  `gorm:"column:ctime;"`
	Utime   int64  `gorm:"column:utime;"`
}
