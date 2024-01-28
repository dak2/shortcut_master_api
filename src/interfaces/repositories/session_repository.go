package repositories

import (
	"fmt"
)

type SessionRepository struct {
	RedisHandler RedisHandler
}

func (db *SessionRepository) Save(session string, userId string) error {
	err := db.RedisHandler.SET(session, userId)
	if err != nil {
		return fmt.Errorf("Failed to save session")
	} else {
		return nil
	}
}

func (db *SessionRepository) Delete(session string) error {
	err := db.RedisHandler.DEL(session)
	if err != nil {
		return fmt.Errorf("Failed to delete session")
	} else {
		return nil
	}
}