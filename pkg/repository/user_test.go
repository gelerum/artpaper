package repository

import (
	"testing"

	"github.com/gelerum/artpaper/pkg/model"
	"github.com/jackc/pgx"
)

func TestCreateUser(t *testing.T) {
	// before
	r, err := NewRepository()
	if err != nil {
		t.Fatal(err)
	}
	insertedUsername := "testusername"
	insertedName := "testname"
	insertedPassword := "testpassword"
	// testing
	err = r.CreateUser(insertedUsername, insertedName, insertedPassword)
	if err != nil {
		t.Error(err)
	}

	var (
		receivedName     string
		receivedPassword string
	)
	row := r.database.QueryRow("SELECT name, password FROM users WHERE username = $1;", insertedUsername)
	err = row.Scan(&receivedName, &receivedPassword)
	if err != nil {
		t.Error(err)
	}
	if insertedName != receivedName {
		t.Error("Doesn't insert properly", insertedName, "!=", receivedName)
	}
	if insertedPassword != receivedPassword {
		t.Error("Doesn't insert properly", insertedPassword, "!=", receivedPassword)
	}
	// after

	_, err = r.database.Exec("DELETE FROM users *;")
	if err != nil {
		t.Fatal(err)
	}

}

func TestGetUser(t *testing.T) {
	// before
	r, err := NewRepository()
	if err != nil {
		t.Fatal(err)
	}
	insertedUsername := "testusername"
	insertedName := "testname"
	insertedPassword := "testpassword"
	_, err = r.database.Exec("INSERT INTO users(username, name, password) VALUES ($1, $2, $3)", insertedUsername, insertedName, insertedPassword)
	if err != nil {
		t.Error(err)
	}
	// testing
	receivedName, receivedPassword, err := r.GetUser(insertedUsername)
	if err != nil {
		t.Error(err)
	}
	if insertedName != receivedName {
		t.Error("Doesn't receive properly", insertedName, "!=", receivedName)
	}
	if insertedPassword != receivedPassword {
		t.Error("Doesn't receive properly", insertedPassword, "!=", receivedPassword)
	}
	// after
	_, err = r.database.Exec("DELETE FROM users *;")
	if err != nil {
		t.Fatal(err)
	}
}

func TestUpdateUser(t *testing.T) {
	// before
	r, err := NewRepository()
	if err != nil {
		t.Fatal(err)
	}
	insertedUsername := "testusername"
	insertedName := "testname"
	insertedPassword := "testpassword"
	_, err = r.database.Exec("INSERT INTO users(username, name, password) VALUES ($1, $2, $3)", insertedUsername, insertedName, insertedPassword)
	if err != nil {
		t.Error(err)
	}
	// testing
	updatedName := "updatedname"
	updatedPassword := "updatedpassword"
	err = r.UpdateUser(insertedUsername, updatedName, updatedPassword)
	if err != nil {
		t.Error(err)
	}

	var (
		receivedName     string
		receivedPassword string
	)
	row := r.database.QueryRow("SELECT name, password FROM users WHERE username = $1;", insertedUsername)
	err = row.Scan(&receivedName, &receivedPassword)
	if err != nil {
		t.Error(err)
	}
	if updatedName != receivedName {
		t.Error("Doesn't update properly", updatedName, "!=", receivedName)
	}
	if updatedPassword != receivedPassword {
		t.Error("Doesn't update properly", updatedPassword, "!=", receivedPassword)
	}
	// after
	_, err = r.database.Exec("DELETE FROM users *;")
	if err != nil {
		t.Fatal(err)
	}
}

func TestDeleteUser(t *testing.T) {
	// before
	r, err := NewRepository()
	if err != nil {
		t.Fatal(err)
	}
	insertedUsername := "testusername"
	insertedName := "testname"
	insertedPassword := "testpassword"
	_, err = r.database.Exec("INSERT INTO users(username, name, password) VALUES ($1, $2, $3)", insertedUsername, insertedName, insertedPassword)
	if err != nil {
		t.Error(err)
	}
	// testing
	err = r.DeleteUser(insertedUsername)
	if err != nil {
		t.Error(err)
	}

	row := r.database.QueryRow("SELECT * FROM users WHERE username = $1;", insertedUsername)
	err = row.Scan()
	if err != pgx.ErrNoRows {
		t.Error(err)
	}
	// after
	_, err = r.database.Exec("DELETE FROM users *;")
	if err != nil {
		t.Fatal(err)
	}
}

func TestFindUsers(t *testing.T) {
	// before
	r, err := NewRepository()
	if err != nil {
		t.Fatal(err)
	}
	insertedFirstUsername := "testusername1"
	insertedSecondUsername := "2testusername"

	insertedName := "testname"
	insertedPassword := "testpassword"
	_, err = r.database.Exec("INSERT INTO users(username, name, password) VALUES ($1, $2, $3)", insertedFirstUsername, insertedName, insertedPassword)
	if err != nil {
		t.Error(err)
	}
	_, err = r.database.Exec("INSERT INTO users(username, name, password) VALUES ($1, $2, $3)", insertedSecondUsername, insertedName, insertedPassword)
	if err != nil {
		t.Error(err)
	}
	// testing
	users, err := r.FindUsers("test", 2)
	if err != nil {
		t.Error(err)
	}
	firstUser := model.FindUsers{
		Username: insertedFirstUsername,
		Name:     insertedName,
	}
	secondUser := model.FindUsers{
		Username: insertedSecondUsername,
		Name:     insertedName,
	}
	if users[0] != firstUser {
		t.Error("Doesn't receive properly\n", users[0], "!=\n", firstUser)
	}
	if users[1] != secondUser {
		t.Error("Doesn't receive properly\n", users[1], "!=\n", secondUser)
	}
	_, err = r.database.Exec("DELETE FROM users *;")
	if err != nil {
		t.Fatal(err)
	}
}
