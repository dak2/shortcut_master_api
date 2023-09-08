package infrastructure

import (
	"net/http"
	controller "shortcut_master_api/src/interfaces/controllers"
	"shortcut_master_api/src/utils"

	"github.com/gorilla/sessions"
	"github.com/labstack/echo-contrib/session"
	"github.com/labstack/echo/v4"
)

type LoginRequest struct {
	Code string `json:"code"`
}

func hello(c echo.Context) error {
	_, err := utils.GetSessionCookie(c)
	if err != nil {
		return err
	}

	return c.JSON(http.StatusOK, "Hello, World!")
}

func Handle(e *echo.Echo) {
	// for session
	e.Use(session.Middleware(sessions.NewCookieStore([]byte("secret"))))

	e.GET("/hello", hello)

	// -- users -- //
	e.GET("/users", func(c echo.Context) error {
		userController := controller.NewUsersController(NewSqlHandler())
		users := userController.GetUser()
		c.Bind(&users)
		return c.JSON(http.StatusOK, users)
	})

	// -- quizzes -- //
	e.GET("/quizzes", func(c echo.Context) error {
		_, err := utils.GetSessionCookie(c)
		if err != nil {
			return err
		}

		quizController := controller.NewQuizzesController(NewSqlHandler())
		quizzes := quizController.GetQuizzes()
		c.Bind(&quizzes)
		return c.JSON(http.StatusOK, quizzes)
	})

	// -- login -- //
	e.POST("/login", func(c echo.Context) error {
		req := new(LoginRequest)
		if err := c.Bind(req); err != nil {
			return c.JSON(http.StatusBadRequest, err)
		}

		sess, err := session.Get("session", c)
		if err != nil {
			return c.JSON(http.StatusBadRequest, "Could not get session")
		}

		loginController := controller.NewLoginController(NewSqlHandler())
		res := loginController.Handle(c, sess, req.Code)
		if res.Err != nil {
			return c.JSON(http.StatusBadRequest, res.Err)
		}

		return c.JSON(http.StatusOK, res.UserInfo.Name)
	})

	// -- logout -- //
	e.POST("/logout", func(c echo.Context) error {
		sess, err := session.Get("session", c)
		if err != nil {
			return c.JSON(http.StatusBadRequest, "Could not get session")
		}

		sess.Options = &sessions.Options{
			MaxAge: -1,
		}
		sess.Save(c.Request(), c.Response())

		return c.JSON(http.StatusOK, "logout")
	})
}
