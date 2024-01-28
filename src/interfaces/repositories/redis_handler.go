package repositories

// MEMO : for di

type RedisHandler interface {
	SET (key string, value string) error
	GET (key string) (string, error)
	DEL (key string) error
}
