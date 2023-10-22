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

func (interactor *QuestionInteractor) GetQuestionsByQuiz(quizType string) ([]entity.Question, error) {
	questions, err := interactor.QuestionRepository.SelectByQuiz(quizType)
	if err != nil {
		return []entity.Question{}, err
	}
	return questions, nil
}
