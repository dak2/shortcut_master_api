package infrastructure

import (
	"github.com/gorilla/sessions"
	"github.com/labstack/echo-contrib/session"
	"github.com/labstack/echo/v4"
	"shortcut_master_api/src/infrastructure/middleware"
)

func Handle(e *echo.Echo) {
	e.Use(session.Middleware(sessions.NewCookieStore([]byte("secret"))))
	e.Use(middleware.VerifyUserMiddleware)

	e.GET("/hello", Hello)
	e.GET("/users", Users)
	e.GET("/quizzes", Quizzes)
	e.GET("/questions", Questions)
	e.GET("/answer_histories", AnswerHistories)
	e.POST("/answers", Answers)
	e.POST("/login", Login)
	e.POST("/logout", Logout)
}
