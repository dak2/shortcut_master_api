package infrastructure

import (
	"net/http"
	entity "shortcut_master_api/src/domain"
	database "shortcut_master_api/src/infrastructure/database"
	redis "shortcut_master_api/src/infrastructure/redis"
	controller "shortcut_master_api/src/interfaces/controllers"

	"github.com/labstack/echo/v4"
)

type LoginRequest struct {
	Code string `json:"code"`
}

type AnswerHistoryRequest struct {
	QuizType string                              `json:"quiz_type"`
	Answers  []entity.AnswerHistoryUpdateRequest `json:"answers"`
}

func Hello(c echo.Context) error {
	return c.JSON(http.StatusOK, "Hello, World!")
}

func Quizzes(c echo.Context) error {
	quizController := getQuizzesController()
	quizzes := quizController.GetQuizzes()
	c.Bind(&quizzes)
	return c.JSON(http.StatusOK, quizzes)
}

func Questions(c echo.Context) error {
	quizType := c.QueryParam("quiz_type")
	if quizType == "" {
		return c.JSON(http.StatusBadRequest, "quiz_type is required")
	}

	questionController := getQuesionsController()
	questions := questionController.GetQuestionsByQuiz(quizType)
	if questions.Err != nil {
		return c.JSON(http.StatusInternalServerError, questions.Err)
	}

	c.Bind(&questions.Questions)
	return c.JSON(http.StatusOK, questions.Questions)
}

func Users(c echo.Context) error {
	userController := getUsersController()
	users := userController.GetUser()
	c.Bind(&users)
	return c.JSON(http.StatusOK, users)
}

func AnswerHistories(c echo.Context) error {
	quizType := c.QueryParam("quiz_type")
	if quizType == "" {
		return c.JSON(http.StatusBadRequest, "quiz_type is required")
	}

	answerHistoryController := getAnswerHistoryController()
	res := answerHistoryController.GetAnswerHistories(quizType)
	if res.Err != nil {
		return c.JSON(http.StatusInternalServerError, res.Err)
	}

	c.Bind(&res.AnswerHistories)
	return c.JSON(http.StatusOK, res.AnswerHistories)
}

func Answers(c echo.Context) error {
	req := new(AnswerHistoryRequest)
	if err := c.Bind(req); err != nil {
		return c.JSON(http.StatusBadRequest, err)
	}

	if req.QuizType == "" {
		return c.JSON(http.StatusBadRequest, "quiz_type is required")
	}

	if len(req.Answers) == 0 {
		return c.JSON(http.StatusBadRequest, "answers is required")
	}

	answerController := getAnswerController()
	answerHistories := answerController.CreateAnswerHistory(req.QuizType, req.Answers)
	if answerHistories.Err != nil {
		return c.JSON(http.StatusInternalServerError, answerHistories.Err)
	}

	c.Bind(&answerHistories.AnswerHistories)
	return c.JSON(http.StatusOK, answerHistories.AnswerHistories)
}

func Login(c echo.Context) error {
	req := new(LoginRequest)
	if err := c.Bind(req); err != nil {
		return c.JSON(http.StatusBadRequest, err)
	}

	sessionController := getSessionController()
	res := sessionController.Login(c, req.Code)
	if res.Err != nil {
		return c.JSON(http.StatusBadRequest, res.Err)
	}

	return c.JSON(http.StatusOK, res.UserInfo.Name)
}

func Logout(c echo.Context) error {
	_, err := c.Cookie("session")
	if err != nil {
		return c.JSON(http.StatusUnauthorized, "Could not get session")
	}
	sessionController := getSessionController()
	err = sessionController.Logout(c)
	if err != nil {
		return c.JSON(http.StatusBadRequest, err)
	}

	return c.JSON(http.StatusOK, "logout")
}

func getQuizzesController() *controller.QuizController {
	return controller.NewQuizzesController(database.NewSqlHandler())
}

func getQuesionsController() *controller.QuestionController {
	return controller.NewQuesionsController(database.NewSqlHandler())
}

func getAnswerController() *controller.AnswerController {
	return controller.NewAnswerController(database.NewSqlHandler())
}

func getAnswerHistoryController() *controller.AnswerHistoryController {
	return controller.NewAnswerHistoryController(database.NewSqlHandler())
}

func getUsersController() *controller.UserController {
	return controller.NewUsersController(database.NewSqlHandler())
}

func getSessionController() *controller.SessionController {
	return controller.NewSessionController(database.NewSqlHandler(), redis.NewRedisHandler())
}

