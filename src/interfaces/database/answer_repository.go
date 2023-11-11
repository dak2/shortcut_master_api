package database

import (
	"errors"
	"fmt"
	entity "shortcut_master_api/src/domain"

	"gorm.io/gorm"
)

type AnswerRepository struct {
	SqlHandler SqlHandler
}

func (db *AnswerRepository) SelectAnswersByQuizType(quizType string) ([]entity.Answer, error) {
	answers := []entity.Answer{}
	res := db.SqlHandler.FindAllByParams(&answers, "quiz_type", quizType)
	if err := res.Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return []entity.Answer{}, fmt.Errorf("Record not found")
		}
		return []entity.Answer{}, fmt.Errorf("Failed to get question")
	}
	return answers, nil
}
