package controllers

import (
	entity "shortcut_master_api/src/domain"
	repository "shortcut_master_api/src/interfaces/repositories"
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

func (controller *QuestionController) GetQuestionsByQuiz(quizType string) QuestionResult {
	res := QuestionResult{
		Questions: []entity.Question{},
		Err:       nil,
	}

	questions, err := controller.Interactor.GetQuestionsByQuiz(quizType)

	if err != nil {
		res.Err = err
		return res
	}

	res.Questions = questions

	return res
}
