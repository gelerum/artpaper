package repository

import (
	"log"

	"github.com/gelerum/artpaper/pkg/model"
)

func (r *Repository) CreateUser(username string, name string, password string) (err error) {
	_, err = r.database.Exec("INSERT INTO users(username, name, password) VALUES ($1, $2, $3)", username, name, password)
	if err != nil {
		log.Printf("Error occurred on user creation from database: %s", err)
	}
	return
}

func (r *Repository) GetUser(username string) (name string, password string, err error) {
	row := r.database.QueryRow("SELECT name, password FROM users WHERE username = $1;", username)
	err = row.Scan(&name, &password)
	if err != nil {
		log.Printf("Error occurred on user selection from database: %s", err)
	}
	return
}

func (r *Repository) UpdateUser(username string, newName string, newPassword string) (err error) {
	_, err = r.database.Exec("UPDATE users SET name = $2, password = $3 WHERE username = $1;", username, newName, newPassword)
	if err != nil {
		log.Printf("Error occurred on user updation from database: %s", err)
	}
	return
}

func (r *Repository) DeleteUser(username string) (err error) {
	_, err = r.database.Exec("DELETE FROM users WHERE username = $1;", username)
	if err != nil {
		log.Printf("Error occurred on user deletion from database: %s", err)
	}
	return
}

func (r *Repository) FindUsers(pattern string, quantity int) (users []model.FindUsers, err error) {
	pattern += "%"
	rows, err := r.database.Query("SELECT username, name FROM users WHERE username LIKE $1 OR name LIKE $1 LIMIT $2;", pattern, quantity)
	if err != nil {
		log.Printf("Error occurred on users selection from database: %s", err)
		return
	}
	for rows.Next() {
		var (
			username string
			name     string
		)
		err = rows.Scan(&username, &name)
		if err != nil {
			log.Println(err)
			return
		}
		user := model.FindUsers{
			Name:     name,
			Username: username,
		}
		users = append(users, user)
	}
	return
}
