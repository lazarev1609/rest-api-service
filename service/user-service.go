package service

import (
	"rest-api-service/model"
	"rest-api-service/repository"
)

func CreateUser(user model.User) (model.User, error) {
	return repository.CreateUser(user)
}

func GetUser(userID int) (model.User, error) {
	return repository.GetUserByID(userID)
}
