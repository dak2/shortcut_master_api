package middleware

import (
	"net/http"
	database "shortcut_master_api/src/infrastructure/database"
	redis "shortcut_master_api/src/infrastructure/redis"
	repository "shortcut_master_api/src/interfaces/repositories"
	"shortcut_master_api/src/utils"

	"github.com/labstack/echo/v4"
)

func VerifyUserMiddleware(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		if c.Path() == "/login" || c.Path() == "/logout" {
			return next(c)
		}

		session, err := utils.GetSessionCookie(c)
		if err != nil {
			return c.JSON(http.StatusUnauthorized, "Unauthorized")
		}

		r := redisHandler()
		gid, err := r.GET(c, session.(string))
		if err != nil || len(gid) == 0 {
			return c.JSON(http.StatusUnauthorized, "Unauthorized")
		}

		ur := userRepository()
		if !ur.ExistsUserByGoogleUserId(gid) {
			return c.JSON(http.StatusUnauthorized, "Unauthorized")
		}

		return next(c)
	}
}

func redisHandler() *redis.RedisHandler {
	return redis.NewRedisHandler()
}

func userRepository() *repository.UserRepository {
	return &repository.UserRepository{
		SqlHandler: database.NewSqlHandler(),
	}
}
