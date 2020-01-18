package redisdao

import "time"

const prefix = "GO_SHORT_URL:"

func SetShortUrl(token string, url string,expire time.Duration) error {
	return client.Set(prefix+token, url, expire).Err()
}

func GetShortUrl(token string) (string, error) {
	return client.Get(prefix + token).Result()
}
