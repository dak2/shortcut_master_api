package question

import entity "shortcut_master_api/src/domain"

// MEMO : for dip

type QuestionRepository interface {
	Select() []entity.Question
	SelectByQuiz(id string) ([]entity.Question, error)
}
