package controllers

import (
	"encoding/base64"
	repository "shortcut_master_api/src/interfaces/repositories"
	loginUsecase "shortcut_master_api/src/usecases/login"
	sessionUsecase "shortcut_master_api/src/usecases/session"
	"github.com/labstack/echo/v4"
)

type SessionController struct {
	LoginInteractor   loginUsecase.LoginInteractor
	SessionInteractor sessionUsecase.SessionInteractor
}

func NewSessionController(sqlHandler repository.SqlHandler, redisHandler repository.RedisHandler) *SessionController {
	return &SessionController{
		LoginInteractor: loginUsecase.LoginInteractor{
			LoginRepository: &repository.UserRepository{
				SqlHandler: sqlHandler,
			},
		},
		SessionInteractor: sessionUsecase.SessionInteractor{
			SessionRepository: &repository.SessionRepository{
				RedisHandler: redisHandler,
			},
		},
	}
}

func (c *SessionController) Login(ctx echo.Context, code string) loginUsecase.GoogleUserResult {
	res := loginUsecase.GoogleUserResult{
		UserInfo: loginUsecase.GoogleUserInfo{
			GoogleUserId: "",
			Name:         "",
			Email:        "",
		},
		Err: nil,
	}

	user, err := c.LoginInteractor.GetUser(code)
	if err != nil {
		res.Err = err
		res.UserInfo = loginUsecase.GoogleUserInfo{}
		return res
	}

	res.UserInfo = loginUsecase.GoogleUserInfo{
		GoogleUserId: user.GoogleUserId,
		Name:         user.Name,
		Email:        user.Email,
	}

	key := base64.StdEncoding.EncodeToString([]byte(res.UserInfo.GoogleUserId))
	saveSessErr := c.SessionInteractor.SaveSession(ctx, key, res.UserInfo.GoogleUserId)
	if saveSessErr != nil {
		res.Err = saveSessErr
		res.UserInfo = loginUsecase.GoogleUserInfo{}
	}

	return res
}

func (c *SessionController) Logout(ctx echo.Context) error {
	err := c.SessionInteractor.DeleteSession(ctx)
	if err != nil {
		return err
	}
	return nil
}
