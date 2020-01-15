package redisdao

const prefix = "GO_SHORT_URL:"

func SetShortUrl(token string, url string) error {
	return client.Set(prefix+token, url, 0).Err()
}

func GetShortUrl(token string) (string, error) {
	return client.Get(prefix + token).Result()
}
