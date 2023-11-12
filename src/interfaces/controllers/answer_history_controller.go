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
		AnswerInteractor: answerUsecase.AnswerInteractor{
			AnswerRepository: &repository.AnswerRepository{
				SqlHandler: sqlHandler,
			},
		},
		AnswerHistoryInteractor: answerHistoryUsecase.AnswerHistoryInteractor{
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

	answerHistories, err := controller.AnswerHistoryInteractor.GetAnswerHistories(quizType)

	if err != nil {
		res.Err = err
		return res
	}

	res.AnswerHistories = answerHistories

	return res
}

func (controller *AnswerHistoryController) CreateAnswerHistory(quizType string, requestAnswerHistories []entity.AnswerHistoryUpdateRequest) AnswerHistoryResult {
	res := AnswerHistoryResult{
		AnswerHistories: []entity.AnswerHistory{},
		Err:             nil,
	}

	var answers []entity.Answer
	var err error
	answers, err = controller.AnswerInteractor.GetAnswers(quizType)
	if err != nil {
		res.Err = err
		return res
	}

	if err := controller.AnswerHistoryInteractor.CreateAnswerHistories(answers, requestAnswerHistories); err != nil {
		res.Err = err
		return res
	}

	return res
}
