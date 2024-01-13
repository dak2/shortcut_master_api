package answer

import entity "shortcut_master_api/src/domain"

// MEMO : for di

type AnswerRepository interface {
	SelectAnswers(quizType string) ([]entity.Answer, error)
}
