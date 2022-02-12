package storage

import (
	"context"
	"log"
	"time"

	"github.com/go-redis/redis/v8"
)

type Storage struct {
	Client *redis.Client
}

func NewStorage(host string, port string, password string) (storage *Storage, err error) {
	client := redis.NewClient(&redis.Options{
		Addr:     host + ":" + port,
		Password: password,
		DB:       0,
	})
	err = client.Ping(context.Background()).Err()
	if err != nil {
		log.Println("error with redis connection: ", err)
		return
	}
	storage = &Storage{
		Client: client,
	}
	return
}

func (s *Storage) CloseConnection() (err error) {
	err = s.Client.Close()
	if err != nil {
		log.Println(err)
	}
	return
}

func (s *Storage) SetEx(key string, value string, expiration time.Duration) (err error) {
	err = s.Client.SetEX(context.Background(), key, value, expiration).Err()
	if err != nil {
		log.Println(err)
	}
	return
}

func (s *Storage) Get(key string) (value string, err error) {
	value, err = s.Client.Get(context.Background(), key).Result()
	if err != nil {
		log.Println(err)
	}
	return
}

func (s *Storage) GetDel(key string) (value string, err error) {
	value, err = s.Client.GetDel(context.Background(), key).Result()
	if err != nil {
		log.Println(err)
	}
	return
}

func (s *Storage) SetArticleCache(key string, title string, body string, creationDate string, username string) (err error) {
	err = s.Client.HSet(context.Background(), key, "title", title, "body", body, "creationDate", creationDate, "username", username).Err()
	if err != nil {
		log.Println(err)
	}
	return
}

func (s *Storage) SetUserCache(key string, name string) (err error) {
	err = s.Client.HSet(context.Background(), key, "name", name).Err()
	if err != nil {
		log.Println(err)
	}
	return
}

func (s *Storage) GetArticleCache(key string) (title string, body string, creationDate string, username string, err error) {
	article, err := s.Client.HGetAll(context.Background(), key).Result()
	if err != nil {
		log.Println(err)
		return
	}
	title = article["title"]
	body = article["body"]
	creationDate = article["creationDate"]
	username = article["username"]
	return
}

func (s *Storage) GetUserCache(key string) (name string, err error) {
	user, err := s.Client.HGetAll(context.Background(), key).Result()
	if err != nil {
		log.Println(err)
		return
	}
	name = user["name"]
	return
}

func (s *Storage) Expire(key string, expiration time.Duration) (err error) {
	err = s.Client.Expire(context.Background(), key, expiration).Err()
	if err != nil {
		log.Println(err)
	}
	return
}

func (s *Storage) TTL(key string) (ttl time.Duration, err error) {
	ttl, err = s.Client.TTL(context.Background(), key).Result()
	if err != nil {
		log.Println(err)
	}
	return
}
