package infrastructure

import (
	"github.com/gorilla/sessions"
	"github.com/labstack/echo-contrib/session"
	"github.com/labstack/echo/v4"
)

func Handle(e *echo.Echo) {
	e.Use(session.Middleware(sessions.NewCookieStore([]byte("secret"))))

	e.GET("/hello", hello)
	e.GET("/users", users)
	e.GET("/quizzes", quizzes)
	e.GET("/questions", questions)
	e.POST("/answers", answers)
	e.POST("/login", login)
	e.POST("/logout", logout)
}
