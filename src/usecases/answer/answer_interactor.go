package answer

import (
	entity "shortcut_master_api/src/domain"
)

type AnswerInteractor struct {
	AnswerRepository AnswerRepository
}

func (interactor *AnswerInteractor) GetAnswersByQuizType(quizType string) ([]entity.Answer, error) {
	questions, err := interactor.AnswerRepository.SelectAnswersByQuizType(quizType)
	if err != nil {
		return []entity.Answer{}, err
	}
	return questions, nil
}
