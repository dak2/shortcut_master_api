package question

import entity "shortcut_master_api/src/domain"

// MEMO : for dip

type QuestionRepository interface {
	Create(entity.Question) (entity.Question, error) // TODO: implement
	Select() []entity.Question
	SelectByQuiz(id string) ([]entity.Question, error)
	Delete(id string) // TODO: implement
}
