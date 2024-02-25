package repositories

import (
	"fmt"

	"github.com/labstack/echo/v4"
)

type SessionRepository struct {
	RedisHandler RedisHandler
}

func (db *SessionRepository) Save(c echo.Context, session string, userId string) error {
	err := db.RedisHandler.SET(c, session, userId)
	if err != nil {
		return fmt.Errorf("Failed to save session")
	} else {
		return nil
	}
}

func (db *SessionRepository) Delete(c echo.Context) error {
	err := db.RedisHandler.DEL(c)
	if err != nil {
		return fmt.Errorf("Failed to delete session")
	} else {
		return nil
	}
}
