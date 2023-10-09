package question

import (
	entity "shortcut_master_api/src/domain"
)

type QuestionInteractor struct {
  QuestionRepository QuestionRepository
}

// TODO: implement
func (interactor *QuestionInteractor) Create(u entity.Question) {
	interactor.QuestionRepository.Create(u)
}

func (interactor *QuestionInteractor) GetQuestions() []entity.Question {
	return interactor.QuestionRepository.Select()
}

func (interactor *QuestionInteractor) GetQuestionsByQuiz(id string) ([]entity.Question, error) {
	return interactor.QuestionRepository.SelectByQuiz(id)
}

// TODO: implement
func (interactor *QuestionInteractor) Delete(id string) {
	interactor.QuestionRepository.Delete(id)
}
