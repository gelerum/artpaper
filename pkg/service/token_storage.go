package service

import (
	"time"
)

func (s *Service) SetRefreshToken(refreshToken string, username string) (err error) {
	err = s.tokenStorage.SetEx(refreshToken, username, 720*time.Hour)
	return
}

func (s *Service) GetDelRefreshToken(refreshToken string) (username string, err error) {
	username, err = s.tokenStorage.GetDel(refreshToken)
	return
}
