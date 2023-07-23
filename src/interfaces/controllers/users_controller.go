package controllers

import (
	entity "shortcut_master_api/src/domain"
	repository "shortcut_master_api/src/interfaces/database"
	userUsecase "shortcut_master_api/src/usecases/user"
	"github.com/labstack/echo"
)

type UserController struct {
	Interactor userUsecase.UserInteractor
}

func NewUsersController(sqlHandler repository.SqlHandler) *UserController {
	return &UserController {
		Interactor: userUsecase.UserInteractor {
			UserRepository: &repository.UserRepository {
				SqlHandler: sqlHandler,
			},
		},
	}
}

func (controller *UserController) Create(c echo.Context) {
	u := entity.User{}
	c.Bind(&u)
	controller.Interactor.Create(u)
	createdUsers := controller.Interactor.GetInfo()
	c.JSON(201, createdUsers)
	return
}

func (controller *UserController) GetUser() []entity.User {
	res := controller.Interactor.GetInfo()
	return res
}

func (controller *UserController) Delete(id string) {
	controller.Interactor.Delete(id)
}
