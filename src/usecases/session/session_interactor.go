package session

import "github.com/labstack/echo/v4"

type SessionInteractor struct {
	SessionRepository SessionRepository
}

func (interactor *SessionInteractor) SaveSession(c echo.Context, session string, userId string) (error) {
	err := interactor.SessionRepository.Save(c, session, userId)
	if err != nil {
		return err
	}
	return nil
}

func (interactor *SessionInteractor) DeleteSession(c echo.Context) (error) {
	err := interactor.SessionRepository.Delete(c)
	if err != nil {
		return err
	}
	return nil
}
