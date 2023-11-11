package answer

import entity "shortcut_master_api/src/domain"

// MEMO : for dip

type AnswerRepository interface {
	SelectAnswersByQuizType(quizType string) ([]entity.Answer, error)
}
