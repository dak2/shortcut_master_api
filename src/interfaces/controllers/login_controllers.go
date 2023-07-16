package controllers

import (
	loginUsecase "shortcut_master_api/src/usecases/login"
)

type LoginController struct {
	Interactor loginUsecase.LoginInteractor
}

func (c *LoginController) Handle(code string) loginUsecase.GoogleUserResult {
	res := loginUsecase.GoogleUserResult{
		UserInfo: loginUsecase.GoogleUserInfo{
			Sub:           "",
			Name:          "",
			Email:         "",
			EmailVerified: false,
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
		EmailVerified: user.EmailVerified,
	}

	return res
}
