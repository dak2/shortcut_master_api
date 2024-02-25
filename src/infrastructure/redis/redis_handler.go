package redis

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/go-redis/redis/v8"
	"github.com/gorilla/sessions"
	"github.com/joho/godotenv"
	"github.com/labstack/echo/v4"
	"github.com/rbcervilla/redisstore/v8"
)

type RedisHandler struct {
	store *redisstore.RedisStore
}

func NewRedisHandler() *RedisHandler {
	store := redisStore()
	return &RedisHandler{store: store}
}

func redisStore() *redisstore.RedisStore {
	err := godotenv.Load(".env")

	if err != nil {
		log.Fatal("Error loading .env file")
	}

	client := redis.NewClient(&redis.Options{
		Addr:     os.Getenv("REDIS_ADDR"),
		Password: "",
		DB:       0,
	})

	store, err := redisstore.NewRedisStore(context.Background(), client)
	if err != nil {
		log.Fatal("Error creating redis store")
	}

	store.KeyPrefix("session-val-")
	store.Options(sessions.Options{
		Path:     "/",
		MaxAge:   86400 * 7, // 7 days
		HttpOnly: true,
		SameSite: http.SameSiteNoneMode,
		Secure:   true,
	})

	return store
}

func (handler *RedisHandler) SET(c echo.Context, key string, value string) error {
	req := c.Request()
	res := c.Response()
	store := handler.store

	session, err := store.New(req, "session")
	if err != nil {
		return fmt.Errorf("failed to create session")
	}

	session.Values[key] = value
	err = session.Save(req, res)
	if err != nil {
		return fmt.Errorf("failed to save session")
	}

	return nil
}

func (handler *RedisHandler) GET(c echo.Context, key string) (string, error) {
	req := c.Request()
	store := handler.store

	session, err := store.Get(req, key)
	if err != nil {
		log.Fatal("failed getting session: ", err)
	}

	return session.Values[key].(string), nil
}

func (handler *RedisHandler) DEL(c echo.Context) error {
	req := c.Request()
	res := c.Response()
	store := handler.store

	session, err := store.New(req, "session")
	if err != nil {
		return fmt.Errorf("failed to create session")
	}

	session.Options.MaxAge = -1
	err = session.Save(req, res) // delete session from redis store
	if err != nil {
		return fmt.Errorf("failed to delete session")
	}

	return nil
}
