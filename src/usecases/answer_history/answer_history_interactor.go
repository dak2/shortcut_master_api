package answer_history

import (
	entity "shortcut_master_api/src/domain"
)

type AnswerHistoryInteractor struct {
	AnswerHistoryRepository AnswerHistoryRepository
}

func (interactor *AnswerHistoryInteractor) GetAnswerHistories(quizType string) ([]entity.AnswerHistory, error) {
	answerHistories, err := interactor.AnswerHistoryRepository.SelectAnswerHistories(quizType)
	if err != nil {
		return []entity.AnswerHistory{}, err
	}
	return answerHistories, nil
}

func (interactor *AnswerHistoryInteractor) CreateAnswerHistories(answers []entity.AnswerHistory) (error) {
	err := interactor.AnswerHistoryRepository.InsertAnswerHistories(answers)
	if err != nil {
		return err
	}
	return nil
}
