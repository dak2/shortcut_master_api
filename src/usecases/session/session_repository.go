package session

import "github.com/labstack/echo/v4"

// MEMO : for di

type SessionRepository interface {
	Save(c echo.Context, session string, userId string) (error)
	Delete(c echo.Context) (error)
}
