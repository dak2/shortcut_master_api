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

func (interactor *AnswerHistoryInteractor) CreateAnswerHistories(answers []entity.Answer, answerHistories []entity.AnswerHistoryUpdateRequest) error {

	filterAnswerHistories := filterAnswerHistories(answers, answerHistories)
	if filterAnswerHistories == nil {
		return nil
	}

	err := interactor.AnswerHistoryRepository.InsertAnswerHistories(filterAnswerHistories)
	if err != nil {
		return err
	}
	return nil
}

func filterAnswerHistories(answers []entity.Answer, answerHistories []entity.AnswerHistoryUpdateRequest) []entity.AnswerHistory {
	var finalAnswerHistories []entity.AnswerHistory

	for _, answerHistory := range answerHistories {
		isCorrect := false
		for _, answer := range answers {
			if answerHistory.AnswerId == answer.ID && answerHistory.Contents == answer.Contents {
				isCorrect = true
				break
			}
		}

		finalAnswerHistories = append(finalAnswerHistories, entity.AnswerHistory{
			AnswerId:  answerHistory.AnswerId,
			Contents:  answerHistory.Contents,
			IsCorrect: isCorrect,
		})
	}

	return finalAnswerHistories
}
