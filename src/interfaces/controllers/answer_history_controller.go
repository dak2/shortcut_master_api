package controllers

import (
	entity "shortcut_master_api/src/domain"
	repository "shortcut_master_api/src/interfaces/repositories"
	usecase "shortcut_master_api/src/usecases/answer_history"
)

type AnswerHistoryController struct {
	Interactor usecase.AnswerHistoryInteractor
}

type AnswerHistoryResult struct {
	AnswerHistories []entity.AnswerHistory
	Err             error
}

func NewAnswerHistoryController(sqlHandler repository.SqlHandler) *AnswerHistoryController {
	return &AnswerHistoryController{
		Interactor: usecase.AnswerHistoryInteractor{
			AnswerHistoryRepository: &repository.AnswerHistoryRepository{
				SqlHandler: sqlHandler,
			},
		},
	}
}

func (controller *AnswerHistoryController) GetAnswerHistories(quizType string) AnswerHistoryResult {
	res := AnswerHistoryResult{
		AnswerHistories: []entity.AnswerHistory{},
		Err:             nil,
	}

	answerHistories, err := controller.Interactor.GetAnswerHistories(quizType)
	if err != nil {
		res.Err = err
		return res
	}

	res.AnswerHistories = answerHistories
	return res
}
