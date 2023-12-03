package infrastructure

import (
	"bytes"
	"encoding/json"
	"io"
	"net/http"

	"net/http/httptest"
	entity "shortcut_master_api/src/domain"
	controller "shortcut_master_api/src/interfaces/controllers"
	"testing"

	"github.com/gorilla/sessions"
	"github.com/labstack/echo-contrib/session"
	"github.com/labstack/echo/v4"
	"github.com/stretchr/testify/assert"
	"gorm.io/gorm"
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

func TestQuestionsEndpoint(t *testing.T) {
	e := echo.New()
	endpoint := "/questions"
	quizType := "slack"

	mockSqlHandler := &SqlHandlerMock{
		MockFindAllByParams: func(obj interface{}, column string, params interface{}) *gorm.DB {
			questions := obj.(*[]entity.Question)
			*questions = []entity.Question{
				{ID: 1, QuizType: quizType, Contents: "メッセージ送信の取り消し"},
			}
			return &gorm.DB{}
		},
	}

	questionsController := controller.NewQuesionsController(mockSqlHandler)

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
		t.Run("with quiz_type", func(t *testing.T) {
			e.GET(endpoint, func(c echo.Context) error {
				questions := questionsController.GetQuestionsByQuiz(quizType)
				return c.JSON(http.StatusOK, questions.Questions)
			})

			cookie := &http.Cookie{
				Name:  "session",
				Value: "some-value",
			}

			rec := makeRequest(cookie)

			assert.Equal(t, http.StatusOK, rec.Code)

			var questions []entity.Question
			if err := json.Unmarshal(rec.Body.Bytes(), &questions); err != nil {
				t.Fatalf("Failed to decode response: %s", err)
			}

			assert.Len(t, questions, 1)
			assert.Equal(t, "メッセージ送信の取り消し", questions[0].Contents)
		})
		t.Run("without quiz_type", func(t *testing.T) {
			e.GET(endpoint, func(c echo.Context) error {
				return c.JSON(http.StatusBadRequest, "quiz_type is required")
			})

			cookie := &http.Cookie{
				Name:  "session",
				Value: "some-value",
			}

			rec := makeRequest(cookie)
			assert.Equal(t, http.StatusBadRequest, rec.Code)
		})
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

func TestAnswersEndpoint(t *testing.T) {
	e := echo.New()
	endpoint := "/answers"
	quizType := "slack"
	answerHistoryRequest := []entity.AnswerHistoryUpdateRequest{
		{AnswerId: 1, Contents: "⌘+Z"},
	}

	mockSqlHandler := &SqlHandlerMock{
		MockFindAllByParams: func(obj interface{}, column string, params interface{}) *gorm.DB {
			answers := obj.(*[]entity.Answer)
			*answers = []entity.Answer{
				{ID: 1, QuestionId: 1, Contents: "⌘+Z"},
			}
			return &gorm.DB{}
		},
		MockCreate: func(obj interface{}) *gorm.DB {
			answerHistories := obj.(*[]entity.AnswerHistory)
			*answerHistories = []entity.AnswerHistory{
				{AnswerId: 1, Contents: "⌘+Z", IsCorrect: true},
			}
			return &gorm.DB{}
		},
	}

	answerController := controller.NewAnswerController(mockSqlHandler)

	makeRequest := func(cookie *http.Cookie, reqBody AnswerHistoryRequest) *httptest.ResponseRecorder {
		req := httptest.NewRequest(http.MethodPost, endpoint, nil)
		if cookie != nil {
			req.AddCookie(cookie)
		}
		req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
		reqBodyBytes, err := json.Marshal(reqBody)
		if err != nil {
			panic(err)
		}
		req.Body = io.NopCloser(bytes.NewBuffer(reqBodyBytes))
		rec := httptest.NewRecorder()
		e.ServeHTTP(rec, req)
		return rec
	}

	t.Run("with session cookie", func(t *testing.T) {
		cookie := &http.Cookie{
			Name:  "session",
			Value: "some-value",
		}

		t.Run("with quiz_type", func(t *testing.T) {
			t.Run("with answers", func(t *testing.T) {

				reqBody := AnswerHistoryRequest{
					QuizType: quizType,
					Answers: answerHistoryRequest,
				}

				e.POST(endpoint, func(c echo.Context) error {
					answerHistories := answerController.CreateAnswerHistory(reqBody.QuizType, reqBody.Answers)
					return c.JSON(http.StatusOK, answerHistories.AnswerHistories)
				})
				rec := makeRequest(cookie, reqBody)
				assert.Equal(t, http.StatusOK, rec.Code)

				var answerHistories []entity.AnswerHistory
				if err := json.Unmarshal(rec.Body.Bytes(), &answerHistories); err != nil {
					t.Fatalf("Failed to decode response: %s", err)
				}

				assert.Len(t, answerHistories, 1)
				assert.Equal(t, 1, answerHistories[0].AnswerId)
			})
			t.Run("without answers", func(t *testing.T) {
				e.POST(endpoint, func(c echo.Context) error {
					return c.JSON(http.StatusBadRequest, "answers is required")
				})

				reqBody := AnswerHistoryRequest{
					QuizType: quizType,
				}

				rec := makeRequest(cookie, reqBody)
				assert.Equal(t, http.StatusBadRequest, rec.Code)
			})
		})
		t.Run("without quiz_type", func(t *testing.T) {
			e.POST(endpoint, func(c echo.Context) error {
				return c.JSON(http.StatusBadRequest, "quiz_type is required")
			})

			reqBody := AnswerHistoryRequest{
				Answers: answerHistoryRequest,
			}

			rec := makeRequest(cookie, reqBody)
			assert.Equal(t, http.StatusBadRequest, rec.Code)
		})
	})
	t.Run("without session cookie", func(t *testing.T) {
		e.POST(endpoint, func(c echo.Context) error {
			return c.JSON(http.StatusUnauthorized, "Cookie doesn't exist")
		})

		reqBody := AnswerHistoryRequest{
			QuizType: quizType,
			Answers: answerHistoryRequest,
		}

		rec := makeRequest(nil, reqBody)

		assert.Equal(t, http.StatusUnauthorized, rec.Code)
		assert.Contains(t, rec.Body.String(), "Cookie doesn't exist")
	})
}
