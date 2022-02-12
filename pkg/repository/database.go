package repository

import (
	"log"
	"os"
	"strconv"

	"github.com/jackc/pgx"
)

type Repository struct {
	database *pgx.Conn
}

func NewRepository() (r *Repository, err error) {
	connection, err := NewConnection()
	r = &Repository{
		database: connection,
	}
	return
}

func NewConnection() (connection *pgx.Conn, err error) {
	port, err := strconv.Atoi(os.Getenv("POSTGRES_PORT"))
	if err != nil {
		log.Printf("Error occurred while trying to user POSTGRES_PORT environment variable: %s", err)
	}
	config := pgx.ConnConfig{
		Host:     os.Getenv("POSTGRES_HOST"),
		Port:     uint16(port),
		Database: os.Getenv("POSTGRES_DB"),
		User:     os.Getenv("POSTGRES_USER"),
		Password: os.Getenv("POSTGRES_PASSWORD"),
	}
	connection, err = pgx.Connect(config)
	if err != nil {
		log.Printf("Error occurred while trying to connect to database: %s", err)
	}
	return
}

func (r *Repository) CloseConnection() (err error) {
	err = r.database.Close()
	if err != nil {
		log.Printf("Error occurred while trying to disconnect with database: %s", err)
	}
	return
}
