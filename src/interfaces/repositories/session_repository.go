package repositories

import (
	"fmt"
)

type SessionRepository struct {
	RedisHandler RedisHandler
}

func (db *SessionRepository) Save(session string, userId string) error {
	err := db.RedisHandler.Set(session, userId)
	if err != nil {
		return fmt.Errorf("failed to save session for user %s: %w", userId, err)
	} else {
		return nil
	}
}

func (db *SessionRepository) Delete(session string) error {
	err := db.RedisHandler.Del(session)
	if err != nil {
		return fmt.Errorf("failed to delete session %s: %w", session, err)
	} else {
		return nil
	}
}
