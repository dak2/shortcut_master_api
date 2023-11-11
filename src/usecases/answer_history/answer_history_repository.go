package answer_history

import (
	entity "shortcut_master_api/src/domain"
)

// MEMO : for dip

type AnswerHistoryRepository interface {
	SelectAnswerHistories(quizType string) ([]entity.AnswerHistory, error)
	InsertAnswerHistories(answers []entity.AnswerHistory) error
}
