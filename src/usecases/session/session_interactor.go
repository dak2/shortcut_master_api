package session

type SessionInteractor struct {
	SessionRepository SessionRepository
}

func (interactor *SessionInteractor) SaveSession(session string, userId string) (error) {
	err := interactor.SessionRepository.Save(session, userId)
	if err != nil {
		return err
	}
	return nil
}

func (interactor *SessionInteractor) DeleteSession(session string) (error) {
	err := interactor.SessionRepository.Delete(session)
	if err != nil {
		return err
	}
	return nil
}
