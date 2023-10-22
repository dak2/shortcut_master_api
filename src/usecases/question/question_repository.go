package question

import entity "shortcut_master_api/src/domain"

// MEMO : for dip

type QuestionRepository interface {
	Select() []entity.Question
	SelectByQuiz(quizType string) ([]entity.Question, error)
}
