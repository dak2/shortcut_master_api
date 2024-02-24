package middleware

import (
	"net/http"
	repository "shortcut_master_api/src/interfaces/repositories"
	redis "shortcut_master_api/src/infrastructure/redis"
	database "shortcut_master_api/src/infrastructure/database"

	"github.com/labstack/echo/v4"
)

func VerifyUserMiddleware(session string) echo.HandlerFunc {
	return func(c echo.Context) error {
		r := redisClient()
		gid, err := r.GET(session)

		if err != nil || len(gid) == 0 {
			return c.JSON(http.StatusUnauthorized, "Unauthorized")
		}

		ur := userRepository()
		if (!ur.ExistsUserByGoogleUserId(gid)) {
			return c.JSON(http.StatusUnauthorized, "Unauthorized")
		}

		return nil;
	}
}

func redisClient() *redis.RedisHandler {
	return redis.NewRedisHandler()
}

func userRepository() *repository.UserRepository {
	return &repository.UserRepository{
		SqlHandler: database.NewSqlHandler(),
	}
}
