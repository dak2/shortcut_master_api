package controllers

import (
	entity "shortcut_master_api/src/domain"
	repository "shortcut_master_api/src/interfaces/database"
	answerUsecase "shortcut_master_api/src/usecases/answer"
	answerHistoryUsecase "shortcut_master_api/src/usecases/answer_history"
)

type AnswerController struct {
	AnswerInteractor        answerUsecase.AnswerInteractor
	AnswerHistoryInteractor answerHistoryUsecase.AnswerHistoryInteractor
}

type AnswerResult struct {
	AnswerHistories []entity.AnswerHistory
	Err             error
}

func NewAnswerController(sqlHandler repository.SqlHandler) *AnswerController {
	return &AnswerController{
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

func (controller *AnswerController) GetAnswerHistories(quizType string) AnswerResult {
	res := AnswerResult{
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

func (controller *AnswerController) CreateAnswerHistory(quizType string, requestAnswerHistories []entity.AnswerHistoryUpdateRequest) AnswerResult {
	res := AnswerResult{
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

	answeredHistories, err := controller.AnswerHistoryInteractor.CreateAnswerHistories(answers, requestAnswerHistories)
	if err != nil {
		res.Err = err
		return res
	}

	res.AnswerHistories = answeredHistories
	return res
}
