package controllers

import (
	"encoding/base64"
	"net/http"
	repository "shortcut_master_api/src/interfaces/repositories"
	loginUsecase "shortcut_master_api/src/usecases/login"
	sessionUsecase "shortcut_master_api/src/usecases/session"

	"github.com/gorilla/sessions"
	"github.com/labstack/echo/v4"
)

type LoginController struct {
	LoginInteractor   loginUsecase.LoginInteractor
	SessionInteractor sessionUsecase.SessionInteractor
}

func NewLoginController(sqlHandler repository.SqlHandler, redisHandler repository.RedisHandler) *LoginController {
	return &LoginController{
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

func (c *LoginController) Login(ctx echo.Context, sess *sessions.Session, code string) loginUsecase.GoogleUserResult {
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

	genSessErr := GenerateSession(ctx, sess, res.UserInfo)
	if genSessErr != nil {
		res.Err = genSessErr
		res.UserInfo = loginUsecase.GoogleUserInfo{}
	}

	saveSessErr := c.SessionInteractor.SaveSession(sess.Values["session"].(string), res.UserInfo.GoogleUserId)
	if saveSessErr != nil {
		res.Err = saveSessErr
		res.UserInfo = loginUsecase.GoogleUserInfo{}
	}

	return res
}

func GenerateSession(ctx echo.Context, sess *sessions.Session, userInfo loginUsecase.GoogleUserInfo) error {
	session_id := base64.StdEncoding.EncodeToString([]byte(userInfo.GoogleUserId))
	sess.Values["session"] = session_id
	sess.Options = &sessions.Options{
		Path:     "/",
		MaxAge:   86400 * 7, // 7 days
		HttpOnly: true,
		SameSite: http.SameSiteNoneMode,
		Secure:   true,
	}
	err := sess.Save(ctx.Request(), ctx.Response())
	if err != nil {
		return err
	}
	return nil
}
