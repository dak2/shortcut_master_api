package repositories

// MEMO : for di

type RedisHandler interface {
	Set (key string, value string) error
	Get (key string) (string, error)
	Del (key string) error
}
