package answer_history

import (
	"fmt"
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

func (interactor *AnswerHistoryInteractor) CreateAnswerHistories(answers []entity.Answer, answerHistories []entity.AnswerHistoryUpdateRequest) ([]entity.AnswerHistory, error) {

	filterAnswerHistories := filterAnswerHistories(answers, answerHistories)
	if len(filterAnswerHistories) == 0 {
		return []entity.AnswerHistory{}, fmt.Errorf("Failed to create answer")
	}

	answeredHistories, err := interactor.AnswerHistoryRepository.InsertAnswerHistories(filterAnswerHistories)
	if err != nil {
		return []entity.AnswerHistory{}, err
	}
	return answeredHistories, nil
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
			UserId:    answerHistory.UserId,
			AnswerId:  answerHistory.AnswerId,
			Contents:  answerHistory.Contents,
			IsCorrect: isCorrect,
		})
	}

	return finalAnswerHistories
}
