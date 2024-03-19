package repository

import (
	"database/sql"
	"log"
	"rest-api-service/model"
)

//var db *sql.DB

func init() {
	var err error
	db, err = sql.Open("postgres", "postgres://postgres:postgres@localhost:15432/rest-api-service?sslmode=disable")
	if err != nil {
		log.Fatal(err)
	}
}

func CreateUser(user model.User) (model.User, error) {
	_, err := db.Exec("INSERT INTO users (name, balance) VALUES ($1, $2)", user.Name, user.Balance)
	return user, err
}

func GetUserByID(userID int) (model.User, error) {
	var user model.User
	err := db.QueryRow("SELECT id, name, balance FROM users WHERE id = $1", userID).Scan(&user.ID, &user.Name, &user.Balance)
	return user, err
}
