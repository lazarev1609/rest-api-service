package repository

import (
	"database/sql"
	"errors"
	_ "github.com/lib/pq"
	"log"
	"rest-api-service/model"
)

var db *sql.DB

func init() {
	var err error
	db, err = sql.Open("postgres", "postgres://postgres:postgres@localhost:15432/rest-api-service?sslmode=disable")
	if err != nil {
		log.Fatal(err)
	}
}

func CreateQuest(quest model.Quest) (model.Quest, error) {
	_, err := db.Exec("INSERT INTO quests (name, cost) VALUES ($1, $2)", quest.Name, quest.Cost)
	return quest, err
}

func CompleteQuest(userID, questID int) error {
	var completed bool
	err := db.QueryRow("SELECT completed FROM user_quests WHERE user_id = $1 AND quest_id = $2", userID, questID).Scan(&completed)
	if err != nil {
		if err == sql.ErrNoRows {
			// Если запись отсутствует, задание еще не выполнено
			return errors.New("quest not completed by user")
		}
		// Возникла ошибка при выполнении запроса к базе данных
		return err
	}

	if completed {
		// Если задание уже выполнено, возвращаем ошибку
		return errors.New("quest already completed by user")
	}

	var cost int
	err = db.QueryRow("SELECT cost FROM quests WHERE id = $1", questID).Scan(&cost)
	if err != nil {
		return err
	}

	_, err = db.Exec("UPDATE users SET balance = balance + $1 WHERE id = $2", cost, userID)
	if err != nil {
		return err
	}

	// Помечаем задание как выполненное
	_, err = db.Exec("INSERT INTO user_quests (user_id, quest_id, completed) VALUES ($1, $2, true)", userID, questID)
	return err
}
