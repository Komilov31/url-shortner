package service

import (
	"fmt"

	"github.com/Komilov31/url-shortener/internal/model"
	"github.com/go-redis/redis/v8"
)

func (s *Service) CreateShortUrl(url model.Url) (*model.Url, error) {
	url.Url = validateUrlScheme(url.Url)

	short_url, err := s.cache.Get(url.Url)
	if err != nil && err != redis.Nil {
		return nil, fmt.Errorf("could not get value from redis: %w", err)
	}

	if err == redis.Nil {
		maxAttempts := 5
		for i := 0; i < maxAttempts; i++ {
			url.ShortUrl = generateShortLink()
			urlInfo, err := s.storage.CreateShortUrl(url)
			if err == nil {
				return urlInfo, nil
			}
		}
		return nil, err
	}

	var urlInfo model.Url
	urlInfo.Url = url.Url
	urlInfo.ShortUrl = short_url

	return &urlInfo, nil
}
