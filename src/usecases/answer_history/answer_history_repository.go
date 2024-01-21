package answer_history

import (
	entity "shortcut_master_api/src/domain"
)

// MEMO : for di

type AnswerHistoryRepository interface {
	SelectAnswerHistories(quizType string) ([]entity.AnswerHistory, error)
	InsertAnswerHistories(answers []entity.AnswerHistory) ([]entity.AnswerHistory, error)
}
