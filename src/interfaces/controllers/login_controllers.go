package controllers

import (
	loginUsecase "shortcut_master_api/src/usecases/login"
	repository "shortcut_master_api/src/interfaces/database"
)

type LoginController struct {
	Interactor loginUsecase.LoginInteractor
}

func NewLoginController(sqlHandler repository.SqlHandler) *LoginController {
	return &LoginController {
		Interactor: loginUsecase.LoginInteractor {
			LoginRepository: &repository.UserRepository {
				SqlHandler: sqlHandler,
			},
		},
	}
}


func (c *LoginController) Handle(code string) loginUsecase.GoogleUserResult {
	res := loginUsecase.GoogleUserResult{
		UserInfo: loginUsecase.GoogleUserInfo{
			Sub:           "",
			Name:          "",
			Email:         "",
		},
		Err: nil,
	}

	user, err := c.Interactor.HandleLogin(code)
	if err != nil {
		res.Err = err
		res.UserInfo = loginUsecase.GoogleUserInfo{}
		return res
	}

	res.UserInfo = loginUsecase.GoogleUserInfo{
		Sub:   user.GoogleUserId,
		Name:  user.Name,
		Email: user.Email,
	}

	return res
}
