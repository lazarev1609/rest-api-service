package service

import (
	"rest-api-service/model"
	"rest-api-service/repository"
)

func CreateQuest(quest model.Quest) (model.Quest, error) {
	return repository.CreateQuest(quest)
}

func CompleteQuest(userID, questID int) error {
	return repository.CompleteQuest(userID, questID)
}
