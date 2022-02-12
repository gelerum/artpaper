package main

import (
	"context"
	"log"
	"os"
	"os/signal"
	"syscall"

	"github.com/gelerum/artpaper/pkg/handler"
	"github.com/gelerum/artpaper/pkg/repository"
	"github.com/gelerum/artpaper/pkg/server"
	"github.com/gelerum/artpaper/pkg/service"
	"github.com/gelerum/artpaper/pkg/storage"
)

func main() {
	tokenStorage, err := storage.NewStorage(os.Getenv("TOKEN_STORAGE_HOST"), os.Getenv("TOKEN_STORAGE_PORT"), os.Getenv("TOKEN_STORAGE_PASSWORD"))
	if err != nil {
		return
	}
	cacheStorage, err := storage.NewStorage(os.Getenv("CACHE_HOST"), os.Getenv("CACHE_PORT"), os.Getenv("CACHE_PASSWORD"))
	if err != nil {
		return
	}
	repository, err := repository.NewRepository()
	if err != nil {
		return
	}
	service := service.NewService(repository, tokenStorage, cacheStorage)
	handler := handler.NewHandler(service)
	server := new(server.Server)
	go func() {
		err = server.Run(os.Getenv("APP_PORT"), handler.InitRoutes())
		if err != nil {
			return
		}
	}()
	log.Print("App is running on ", os.Getenv("APP_PORT"))

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGTERM, syscall.SIGINT)
	<-quit
	log.Print("App is shutting down")

	server.Shutdown(context.Background())
	repository.CloseConnection()
	tokenStorage.CloseConnection()
}
