package utils

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

func GetSessionCookie(c echo.Context) (*http.Cookie, error) {
	cookie, err := c.Cookie("session")
	if err != nil {
		if err == http.ErrNoCookie {
			return nil, echo.NewHTTPError(http.StatusUnauthorized, "Cookie doesn't exist")
		}
		return nil, echo.NewHTTPError(http.StatusBadRequest, err)
	}
	return cookie, nil
}
