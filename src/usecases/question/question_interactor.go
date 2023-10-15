package question

import (
	entity "shortcut_master_api/src/domain"
)

type QuestionInteractor struct {
	QuestionRepository QuestionRepository
}

func (interactor *QuestionInteractor) GetQuestions() []entity.Question {
	return interactor.QuestionRepository.Select()
}

func (interactor *QuestionInteractor) GetQuestionsByQuiz(id int) ([]entity.Question, error) {
	questions, err := interactor.QuestionRepository.SelectByQuiz(id)
	if err != nil {
		return []entity.Question{}, err
	}
	return questions, nil
}
