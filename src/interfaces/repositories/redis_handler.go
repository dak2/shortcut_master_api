package repositories

import "github.com/labstack/echo/v4"

// MEMO : for di

type RedisHandler interface {
	SET (c echo.Context, key string, value string) error
	GET (c echo.Context, key string) (string, error)
	DEL (c echo.Context) error
}
