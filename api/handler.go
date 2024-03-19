package api

import (
	"encoding/json"
	"net/http"
	"rest-api-service/model"
	"rest-api-service/service"
	"strconv"
)

func CreateUser(w http.ResponseWriter, r *http.Request) {
	var user model.User
	json.NewDecoder(r.Body).Decode(&user)
	newUser, err := service.CreateUser(user)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	json.NewEncoder(w).Encode(newUser)
}

func CreateQuest(w http.ResponseWriter, r *http.Request) {
	var quest model.Quest
	json.NewDecoder(r.Body).Decode(&quest)
	newQuest, err := service.CreateQuest(quest)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	json.NewEncoder(w).Encode(newQuest)
}

func CompleteQuest(w http.ResponseWriter, r *http.Request) {
	userID, _ := strconv.Atoi(r.FormValue("user_id"))
	questID, _ := strconv.Atoi(r.FormValue("quest_id"))
	err := service.CompleteQuest(userID, questID)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusOK)
	w.Write([]byte("Quest completed successfully"))
}

func GetUser(w http.ResponseWriter, r *http.Request) {
	userID, _ := strconv.Atoi(r.FormValue("id"))
	user, err := service.GetUser(userID)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	json.NewEncoder(w).Encode(user)
}
