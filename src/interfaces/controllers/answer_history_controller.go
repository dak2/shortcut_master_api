package controllers

import (
	entity "shortcut_master_api/src/domain"
	repository "shortcut_master_api/src/interfaces/database"
	answerUsecase "shortcut_master_api/src/usecases/answer"
	answerHistoryUsecase "shortcut_master_api/src/usecases/answer_history"
)

type AnswerHistoryController struct {
	AnswerInteractor        answerUsecase.AnswerInteractor
	AnswerHistoryInteractor answerHistoryUsecase.AnswerHistoryInteractor
}

type AnswerHistoryResult struct {
	AnswerHistories []entity.AnswerHistory
	Err             error
}

func NewAnswerHistoryController(sqlHandler repository.SqlHandler) *AnswerHistoryController {
	return &AnswerHistoryController{
		AnswerHistoryInteractor: answerHistoryUsecase.AnswerHistoryInteractor{
			AnswerHistoryRepository: &repository.AnswerHistoryRepository{
				SqlHandler: sqlHandler,
			},
		},
	}
}

func (controller *AnswerController) GetAnswerHistories(quizType string, historyType string) AnswerHistoryResult {
	res := AnswerHistoryResult{
		AnswerHistories: []entity.AnswerHistory{},
		Err:             nil,
	}

	answerHistories, err := controller.AnswerHistoryInteractor.GetAnswerHistories(quizType, historyType)
	if err != nil {
		res.Err = err
		return res
	}

	res.AnswerHistories = answerHistories

	return res
}
