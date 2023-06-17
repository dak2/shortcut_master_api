package main

import (
	"fmt"
	"net/http"
	router "shortcut_master_api/src/infrastructure"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func hello(c echo.Context) error {
	return c.String(http.StatusOK, "Hello, World!")
}

func main() {
	fmt.Println("server start")
	e := echo.New()
	e.Use(middleware.CORSWithConfig(middleware.CORSConfig{
		// TODO : set valid origin
		AllowOrigins: []string{"*"},
		AllowMethods: []string{http.MethodGet, http.MethodHead, http.MethodPut, http.MethodPatch, http.MethodPost, http.MethodDelete},
	}))
	router.Handle(e)
	e.Logger.Fatal(e.Start(":3000"))
}
