package utils

import (
	"net/http"

	"github.com/gorilla/sessions"
	"github.com/labstack/echo/v4"
)

var store = sessions.NewCookieStore([]byte("secret"))

func GetSessionCookie(c echo.Context) (interface{}, error) {
	session, err := store.Get(c.Request(), "session")
	if err != nil {
		if err == http.ErrNoCookie {
			return nil, echo.NewHTTPError(http.StatusUnauthorized, "Session cookie doesn't exist")
		}
		return nil, echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}

	value := session.Values["session"]
	if value == nil {
		return nil, echo.NewHTTPError(http.StatusNotFound, "Session value not found")
	}

	return value, nil
}
