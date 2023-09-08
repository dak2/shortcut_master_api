package quiz

import (
	entity "shortcut_master_api/src/domain"
)

type QuizInteractor struct {
  QuizRepository QuizRepository
}

// TODO: implement
func (interactor *QuizInteractor) Create(u entity.Quiz) {
	interactor.QuizRepository.Create(u)
}

func (interactor *QuizInteractor) GetQuizzes() []entity.Quiz {
	return interactor.QuizRepository.Select()
}

// TODO: implement
func (interactor *QuizInteractor) Delete(id string) {
	interactor.QuizRepository.Delete(id)
}
