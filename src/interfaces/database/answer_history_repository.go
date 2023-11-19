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

func (db *AnswerHistoryRepository) SelectAnswerHistories(quizType string) ([]entity.AnswerHistory, error) {
	answerHistories := []entity.AnswerHistory{}
	res := db.SqlHandler.FindAllByParams(&answerHistories, "quiz_type", quizType)
	if err := res.Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return []entity.AnswerHistory{}, fmt.Errorf("Record not found")
		}
		return []entity.AnswerHistory{}, fmt.Errorf("Failed to get question")
	}
	return answerHistories, nil
}

func (db *AnswerHistoryRepository) InsertAnswerHistories(answerHistories []entity.AnswerHistory) ([]entity.AnswerHistory, error) {
	res := db.SqlHandler.Create(&answerHistories)
	if err := res.Error; err != nil {
		return []entity.AnswerHistory{}, fmt.Errorf("Failed to create answer")
	}
	return answerHistories, nil
}
