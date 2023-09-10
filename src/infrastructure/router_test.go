package infrastructure

import (
	"encoding/json"
	"net/http"

	"net/http/httptest"
	entity "shortcut_master_api/src/domain"
	controller "shortcut_master_api/src/interfaces/controllers"
	"testing"

	"github.com/gorilla/sessions"
	"github.com/labstack/echo-contrib/session"
	"github.com/labstack/echo/v4"
	"github.com/stretchr/testify/assert"
)

func TestHelloEndpoint(t *testing.T) {
	e := echo.New()
	endpoint := "/hello"
	e.Use(session.Middleware(sessions.NewCookieStore([]byte("secret"))))
	e.GET(endpoint, hello)

	t.Run("with valid cookie", func(t *testing.T) {
		req := httptest.NewRequest(http.MethodGet, endpoint, nil)
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
		req := httptest.NewRequest(http.MethodGet, endpoint, nil)
		rec := httptest.NewRecorder()

		e.ServeHTTP(rec, req)

		statusCode := rec.Code
		responseBody := rec.Body.String()

		assert.Equal(t, http.StatusUnauthorized, statusCode)
		assert.Contains(t, responseBody, "Cookie doesn't exist")
	})
}

func TestQuizEndpoint(t *testing.T) {
	e := echo.New()
	endpoint := "/quizzes"

	mockSqlHandler := &SqlHandlerMock{
		MockFindAll: func(obj interface{}) {
			quizzes := obj.(*[]entity.Quiz)
			*quizzes = []entity.Quiz{
				{ID: 1, Name: "Slack", Type: "macOS"},
			}
		},
	}

	quizController := controller.NewQuizzesController(mockSqlHandler)

	makeRequest := func(cookie *http.Cookie) *httptest.ResponseRecorder {
		req := httptest.NewRequest(http.MethodGet, endpoint, nil)
		if cookie != nil {
			req.AddCookie(cookie)
		}
		rec := httptest.NewRecorder()
		e.ServeHTTP(rec, req)
		return rec
	}

	t.Run("with session cookie", func(t *testing.T) {
		e.GET(endpoint, func(c echo.Context) error {
			quizzes := quizController.GetQuizzes()
			return c.JSON(http.StatusOK, quizzes)
		})

		cookie := &http.Cookie{
			Name:  "session",
			Value: "some-value",
		}

		rec := makeRequest(cookie)

		assert.Equal(t, http.StatusOK, rec.Code)

		var quizzes []entity.Quiz
		if err := json.Unmarshal(rec.Body.Bytes(), &quizzes); err != nil {
			t.Fatalf("Failed to decode response: %s", err)
		}

		assert.Len(t, quizzes, 1)
		assert.Equal(t, "Slack", quizzes[0].Name)
	})

	t.Run("without session cookie", func(t *testing.T) {
		e.GET(endpoint, func(c echo.Context) error {
			return c.JSON(http.StatusUnauthorized, "Cookie doesn't exist")
		})

		rec := makeRequest(nil)

		assert.Equal(t, http.StatusUnauthorized, rec.Code)
		assert.Contains(t, rec.Body.String(), "Cookie doesn't exist")
	})
}
