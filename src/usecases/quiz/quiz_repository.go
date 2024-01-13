package quiz

import entity "shortcut_master_api/src/domain"

// MEMO : for di

type QuizRepository interface {
	Create(entity.Quiz) (entity.Quiz, error) // TODO: implement
	Select() []entity.Quiz
	Delete(id string) // TODO: implement
}
