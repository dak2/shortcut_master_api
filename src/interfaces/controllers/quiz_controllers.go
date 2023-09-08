package controllers

import (
	entity "shortcut_master_api/src/domain"
	repository "shortcut_master_api/src/interfaces/database"
	quizUsecase "shortcut_master_api/src/usecases/quiz"

	"github.com/labstack/echo"
)

type QuizController struct {
	Interactor quizUsecase.QuizInteractor
}

func NewQuizzesController(sqlHandler repository.SqlHandler) *QuizController {
	return &QuizController{
		Interactor: quizUsecase.QuizInteractor{
			QuizRepository: &repository.QuizRepository{
				SqlHandler: sqlHandler,
			},
		},
	}
}

// TODO: implement
func (controller *QuizController) Create(c echo.Context) {
	u := entity.Quiz{}
	c.Bind(&u)
	controller.Interactor.Create(u)
	createdUsers := controller.Interactor.GetQuizzes()
	c.JSON(201, createdUsers)
	return
}

func (controller *QuizController) GetQuizzes() []entity.Quiz {
	res := controller.Interactor.GetQuizzes()
	return res
}

// TODO: implement
func (controller *QuizController) Delete(id string) {
	controller.Interactor.Delete(id)
}
