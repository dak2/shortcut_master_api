package session

// MEMO : for di

type SessionRepository interface {
	Save(session string, userId string) (error)
	Delete(session string) (error)
}
