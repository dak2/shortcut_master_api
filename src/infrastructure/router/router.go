package router

import (
	"net/http"
	"github.com/labstack/echo/v4"
	infrastructure "short_cut_master_api/src/infrastructure"
	controller "short_cut_master_api/src/interfaces/controllers"
)

func hello(c echo.Context) error {
	return c.String(http.StatusOK, "Hello, World!")
}

func Handle(e *echo.Echo) {
	e.GET("/", hello)
	userController := controller.NewUsersController(infrastructure.NewSqlHandler())

	// -- users -- //
	e.GET("/users", func(c echo.Context) error {
		users := userController.GetUser()
		c.Bind(&users)
		return c.JSON(http.StatusOK, users)
	})
	e.POST("/login", hello)
}
