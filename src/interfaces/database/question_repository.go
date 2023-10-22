package database

import (
	"errors"
	"fmt"
	entity "shortcut_master_api/src/domain"

	"gorm.io/gorm"
)

type QuestionRepository struct {
	SqlHandler SqlHandler
}

func (db *QuestionRepository) Select() []entity.Question {
	question := []entity.Question{}
	db.SqlHandler.FindAll(&question)
	return question
}

func (db *QuestionRepository) SelectByQuiz(quizType string) ([]entity.Question, error) {
	question := []entity.Question{}
	res := db.SqlHandler.FindAllByParams(&question, "quiz_type", quizType)
	if err := res.Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return []entity.Question{}, fmt.Errorf("Record not found")
		}
		return []entity.Question{}, fmt.Errorf("Failed to get question")
	}
	return question, nil
}
