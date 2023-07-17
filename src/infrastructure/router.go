package infrastructure

import (
	"net/http"
	controller "shortcut_master_api/src/interfaces/controllers"

	"github.com/gorilla/sessions"
	"github.com/labstack/echo-contrib/session"
	"github.com/labstack/echo/v4"
)

type LoginRequest struct {
	Code string `json:"code"`
}

func hello(c echo.Context) error {
	return c.String(http.StatusOK, "Hello, World!")
}

func Handle(e *echo.Echo) {
	// for session
	e.Use(session.Middleware(sessions.NewCookieStore([]byte("secret"))))

	e.GET("/", hello)

	// -- users -- //
	e.GET("/users", func(c echo.Context) error {
		userController := controller.NewUsersController(NewSqlHandler())
		users := userController.GetUser()
		c.Bind(&users)
		return c.JSON(http.StatusOK, users)
	})

	// -- login -- //
	e.POST("/login", func(c echo.Context) error {
		req := new(LoginRequest)
		if err := c.Bind(req); err != nil {
			return err
		}

		sess, err := session.Get("session", c)
		if err != nil {
			return err
		}

		loginController := controller.NewLoginController(NewSqlHandler())
		res := loginController.Handle(c, sess, req.Code)
		if res.Err != nil {
			return res.Err
		}

		return c.JSON(http.StatusOK, res.UserInfo.Email)
	})
}
