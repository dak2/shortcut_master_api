package main

import (
	"fmt"
	"net/http"
	"github.com/labstack/echo/v4"
	"short_cut_master_api/src/infrastructure"
)

func hello(c echo.Context) error {
	return c.String(http.StatusOK, "Hello, World!")
}

func main() {
	fmt.Println("server start")
	e := echo.New()
	infrastructure.HandleRouting(e)
	e.Logger.Fatal(e.Start(":3000"))
}
