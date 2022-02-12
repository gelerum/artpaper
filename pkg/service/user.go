package service

import "github.com/gelerum/artpaper/pkg/model"

func (s *Service) CreateUser(username string, name string, password string) (err error) {
	password = generatePasswordHash(password)
	err = s.database.CreateUser(username, name, password)
	return
}

func (s *Service) GetUser(username string) (name string, err error) {
	name, _, err = s.database.GetUser(username)
	return
}

func (s *Service) UpdateUser(username string, newName string, newPassword string) (err error) {
	passwordHash := generatePasswordHash(newPassword)
	err = s.database.UpdateUser(username, newName, passwordHash)
	return
}

func (s *Service) DeleteUser(username string) (err error) {
	err = s.database.DeleteAllArticles(username)
	if err != nil {
		return
	}
	err = s.database.DeleteUser(username)
	return
}

func (s *Service) FindUsers(pattern string, quantity int) (users []model.User, err error) {
	users, err = s.database.FindUsers(pattern, quantity)
	return
}
