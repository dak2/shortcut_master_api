package database

import (
	"errors"
	"fmt"
	entity "shortcut_master_api/src/domain"

	"gorm.io/gorm"
)

type AnswerHistoryRepository struct {
	SqlHandler SqlHandler
}

func (db *AnswerHistoryRepository) SelectAnswerHistories(quizType string) ([]entity.Answer, error) {
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

func (db *AnswerHistoryRepository) InsertAnswerHistories(answers []entity.Answer) (error) {
	res := db.SqlHandler.Create(&answers)
	if err := res.Error; err != nil {
		return fmt.Errorf("Failed to create answer")
	}
	return nil
}
