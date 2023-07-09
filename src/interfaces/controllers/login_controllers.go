package controllers

import (
	loginUsecase "shortcut_master_api/src/usecases/login"
)

type LoginController struct {
	Interactor loginUsecase.LoginInteractor
}

func (c *LoginController) Handle (code string) loginUsecase.GoogleUserResult {
	return c.Interactor.GetGoogleUser(code)
}
