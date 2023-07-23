package controllers

import (
	"encoding/base64"
	repository "shortcut_master_api/src/interfaces/database"
	loginUsecase "shortcut_master_api/src/usecases/login"

	"github.com/gorilla/sessions"
	"github.com/labstack/echo/v4"
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

func (c *LoginController) Handle(ctx echo.Context, sess *sessions.Session, code string) loginUsecase.GoogleUserResult {
	res := loginUsecase.GoogleUserResult{
		UserInfo: loginUsecase.GoogleUserInfo{
			GoogleUserId:  "",
			Name:          "",
			Email:         "",
		},
		Err: nil,
	}

	user, err := c.Interactor.GetUser(code)
	if err != nil {
		res.Err = err
		res.UserInfo = loginUsecase.GoogleUserInfo{}
		return res
	}

	res.UserInfo = loginUsecase.GoogleUserInfo{
		GoogleUserId:   user.GoogleUserId,
		Name:					  user.Name,
		Email:					user.Email,
	}

	sessErr := GenerateSession(ctx, sess, res.UserInfo)
	if sessErr != nil {
		res.Err = err
		res.UserInfo = loginUsecase.GoogleUserInfo{}
	}

	return res
}

// MEMO: セッションストアではなくCookieにセッションIDを保存する
func GenerateSession(ctx echo.Context, sess *sessions.Session, userInfo loginUsecase.GoogleUserInfo) error {
	session_id := base64.StdEncoding.EncodeToString([]byte(userInfo.GoogleUserId))
	sess.Values["session"] = session_id
	sess.Options = &sessions.Options{
		Path:     "/",
		MaxAge:   86400 * 7, // 7 days
		HttpOnly: true,
	}
	err := sess.Save(ctx.Request(), ctx.Response())
	if err != nil {
		return err
	}
	return nil
}
