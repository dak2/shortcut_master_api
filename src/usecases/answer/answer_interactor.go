package answer

import (
	entity "shortcut_master_api/src/domain"
)

type AnswerInteractor struct {
	AnswerRepository AnswerRepository
}

func (interactor *AnswerInteractor) GetAnswers(quizType string) ([]entity.Answer, error) {
	answers, err := interactor.AnswerRepository.SelectAnswers(quizType)
	if err != nil {
		return []entity.Answer{}, err
	}
	return answers, nil
}
