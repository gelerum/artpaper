package service

import (
	"github.com/gelerum/artpaper/pkg/repository"
	"github.com/gelerum/artpaper/pkg/storage"
)

type Service struct {
	database     *repository.Repository
	tokenStorage *storage.Storage
	cacheStorage *storage.Storage
}

func NewService(db *repository.Repository, tokenStorage *storage.Storage, cacheStorage *storage.Storage) (service *Service) {
	service = &Service{
		database:     db,
		tokenStorage: tokenStorage,
		cacheStorage: cacheStorage,
	}
	return
}
