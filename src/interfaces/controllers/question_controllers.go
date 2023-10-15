package controllers

import (
	entity "shortcut_master_api/src/domain"
	repository "shortcut_master_api/src/interfaces/database"
	quesionUsecase "shortcut_master_api/src/usecases/question"
)

type QuestionController struct {
	Interactor quesionUsecase.QuestionInteractor
}

type QuestionResult struct {
	Questions []entity.Question
	Err       error
}

func NewQuesionsController(sqlHandler repository.SqlHandler) *QuestionController {
	return &QuestionController{
		Interactor: quesionUsecase.QuestionInteractor{
			QuestionRepository: &repository.QuestionRepository{
				SqlHandler: sqlHandler,
			},
		},
	}
}

func (controller *QuestionController) GetQuestionsByQuiz(id int) QuestionResult {
	res := QuestionResult{
		Questions: []entity.Question{},
		Err:       nil,
	}

	questions, err := controller.Interactor.GetQuestionsByQuiz(id)

	if err != nil {
		res.Err = err
		return res
	}

	res.Questions = questions

	return res
}
