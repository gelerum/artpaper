package service

import (
	"crypto/sha1"
	"encoding/hex"
	"errors"
	"fmt"
	"log"
	"os"
	"time"

	"github.com/golang-jwt/jwt"
)

func (s *Service) Login(username string, password string) (accessToken string, refreshToken string, err error) {
	passwordHash := generatePasswordHash(password)
	_, dbPasswordHash, err := s.database.GetUser(username)
	if err != nil {
		return
	}
	if passwordHash != dbPasswordHash {
		err = errors.New("incorrect password")
		return
	}
	accessToken, refreshToken, err = s.RefreshTokens(username)
	return
}

func (s *Service) RefreshTokens(username string) (accessToken string, refreshToken string, err error) {
	accessToken, err = generateToken(username, "access", 48)
	if err != nil {
		err = fmt.Errorf("error occurred on token generation: %s", err)
		return
	}
	refreshToken, err = generateToken(username, "refresh", 720)
	if err != nil {
		err = fmt.Errorf("error occurred on token generation: %s", err)
		return
	}
	err = s.SetRefreshToken(refreshToken, username)
	if err != nil {
		err = fmt.Errorf("error occurred on token writing into redis: %s", err)
		return
	}
	return
}

func ValidateToken(accessToken string, tokenType string) (username string, valid bool) {
	token, valid := tokenIsValid(accessToken)
	if !valid {
		valid = false
		return
	}
	claims, ok := token.Claims.(jwt.MapClaims)
	if ok && token.Valid {
		valid = claims["tokenType"] == tokenType
		username = claims["username"].(string)
		return
	}
	valid = false
	return
}

func tokenIsValid(tokenString string) (token *jwt.Token, valid bool) {
	hmacSampleSecret := []byte(os.Getenv("JWT_SIGN_KEY"))

	token, err := jwt.Parse(tokenString,
		func(token *jwt.Token) (interface{}, error) {
			if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
				return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
			}
			return hmacSampleSecret, nil
		})

	if err != nil {
		valid = false
		return
	}
	valid = true
	return
}

func generateToken(username string, tokenType string, hours int) (tokenString string, err error) {
	token := jwt.New(jwt.SigningMethodHS256)
	claims := token.Claims.(jwt.MapClaims)
	claims["exp"] = time.Now().Add(time.Duration(hours) * time.Hour).Unix()
	claims["iat"] = time.Now().Unix()
	claims["tokenType"] = tokenType
	claims["username"] = username
	tokenString, err = token.SignedString([]byte(os.Getenv("JWT_SIGN_KEY")))
	if err != nil {
		log.Println(err)
	}
	return
}

func generatePasswordHash(password string) (passwordHash string) {
	salt := os.Getenv("AUTH_SALT")
	password += salt
	hash := sha1.New()
	hash.Write([]byte(password))
	passwordHash = hex.EncodeToString(hash.Sum(nil))
	return
}
