package service

import (
	"time"
)

func (s *Service) UploadArticleCache(path string, title string, body string, creationDate string, username string) (err error) {
	err = s.cacheStorage.SetArticleCache(path, title, body, creationDate, username)
	if err != nil {
		return
	}
	err = s.cacheStorage.Expire(path, 5*time.Minute)
	return
}

func (s *Service) LoadArticleCache(path string) (title string, body string, creationDate string, username string, err error) {
	title, body, creationDate, username, err = s.cacheStorage.GetArticleCache(path)
	if err != nil {
		return
	}
	ttl, err := s.cacheStorage.TTL(path)
	if err != nil {
		return
	}
	ttl += time.Minute
	err = s.cacheStorage.Expire(path, ttl)
	return
}

func (s *Service) UploadUserCache(path string, name string) (err error) {
	err = s.cacheStorage.SetUserCache(path, name)
	if err != nil {
		return
	}
	err = s.cacheStorage.Expire(path, 5*time.Minute)
	return
}

func (s *Service) LoadUserCache(path string) (name string, err error) {
	name, err = s.cacheStorage.GetUserCache(path)
	if err != nil {
		return
	}
	ttl, err := s.cacheStorage.TTL(path)
	if err != nil {
		return
	}
	ttl += time.Minute
	err = s.cacheStorage.Expire(path, ttl)
	return
}
