package infrastructure

import (
	"encoding/json"
	"net/http"

	"net/http/httptest"
	"testing"

	"github.com/gorilla/sessions"
	"github.com/labstack/echo-contrib/session"
	"github.com/labstack/echo/v4"
	"github.com/stretchr/testify/assert"
)

func TestHelloEndpoint(t *testing.T) {
	e := echo.New()
	e.Use(session.Middleware(sessions.NewCookieStore([]byte("secret"))))
	e.GET("/hello", hello)

	t.Run("with valid cookie", func(t *testing.T) {
		req := httptest.NewRequest("GET", "/hello", nil)
		cookie := &http.Cookie{
			Name:  "session",
			Value: "some-value",
		}
		req.AddCookie(cookie)
		rec := httptest.NewRecorder()

		e.ServeHTTP(rec, req)

		var response string
		if err := json.Unmarshal(rec.Body.Bytes(), &response); err != nil {
			t.Fatalf("Failed to decode response: %s", err)
		}

		assert.Equal(t, http.StatusOK, rec.Code)
		assert.Equal(t, "Hello, World!", response)
	})

	t.Run("without cookie", func(t *testing.T) {
    req := httptest.NewRequest("GET", "/hello", nil)
    rec := httptest.NewRecorder()

    e.ServeHTTP(rec, req)

    statusCode := rec.Code
    responseBody := rec.Body.String()

    assert.Equal(t, http.StatusUnauthorized, statusCode)
    assert.Contains(t, responseBody, "Cookie doesn't exist")
	})
}
