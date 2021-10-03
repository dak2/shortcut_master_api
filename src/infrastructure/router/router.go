package router

import (
	"net/http"
	infrastructure "short_cut_master_api/src/infrastructure"
	controller "short_cut_master_api/src/interfaces/controllers"

	"github.com/labstack/echo/v4"
)

func hello(c echo.Context) error {
	return c.String(http.StatusOK, "Hello, World!")
}

type Token struct {
	Token string
}

func Handle(e *echo.Echo) {
	// init controllers
	userController := controller.NewUsersController(infrastructure.NewSqlHandler())
	e.GET("/", hello)

	// -- login -- //
	e.POST("/login", func(c echo.Context) error {
		var token Token
		if err := c.Bind(&token); err != nil {
			panic("token does not exists")
		}
		res := userController.Login(token.Token)
		return c.JSON(http.StatusOK, res)
	})

	// -- users -- //
	e.GET("/users", func(c echo.Context) error {
		users := userController.GetUser()
		c.Bind(&users)
		return c.JSON(http.StatusOK, users)
	})
}
