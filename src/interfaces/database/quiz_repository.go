package database

import (
	"fmt"
	entity "shortcut_master_api/src/domain"
)

type QuizRepository struct {
	SqlHandler SqlHandler
}

// TODO: implement
func (db *QuizRepository) Create(u entity.Quiz) (entity.Quiz, error) {
	res := db.SqlHandler.Create(&u)
	if err := res.Error; err != nil {
		return entity.Quiz{}, fmt.Errorf("Failed to create quiz: %w", err)
	} else {
		return u, nil
	}
}

func (db *QuizRepository) Select() []entity.Quiz {
	quiz := []entity.Quiz{}
	db.SqlHandler.FindAll(&quiz)
	return quiz
}

// TODO: implement
func (db *QuizRepository) Delete(id string) {
	quiz := []entity.Quiz{}
	db.SqlHandler.DeleteById(&quiz, id)
}
