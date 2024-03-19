package api

import (
	"github.com/gorilla/mux"
)

func NewRouter() *mux.Router {
	router := mux.NewRouter()
	router.HandleFunc("/users", CreateUser).Methods("POST")
	router.HandleFunc("/quests", CreateQuest).Methods("POST")
	router.HandleFunc("/complete", CompleteQuest).Methods("POST")
	router.HandleFunc("/user/{id}", GetUser).Methods("GET")
	return router
}
