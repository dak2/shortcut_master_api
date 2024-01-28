package infrastructure

import (
	"net/http"
	entity "shortcut_master_api/src/domain"
	controller "shortcut_master_api/src/interfaces/controllers"
	"shortcut_master_api/src/utils"

	"github.com/gorilla/sessions"
	"github.com/labstack/echo-contrib/session"
	"github.com/labstack/echo/v4"
)

type LoginRequest struct {
	Code string `json:"code"`
}

type AnswerHistoryRequest struct {
	QuizType string                              `json:"quiz_type"`
	Answers  []entity.AnswerHistoryUpdateRequest `json:"answers"`
}

func hello(c echo.Context) error {
	_, err := utils.GetSessionCookie(c)
	if err != nil {
		return c.JSON(http.StatusUnauthorized, err)
	}

	return c.JSON(http.StatusOK, "Hello, World!")
}

func quizzes(c echo.Context) error {
	_, err := utils.GetSessionCookie(c)
	if err != nil {
		return c.JSON(http.StatusUnauthorized, err)
	}

	quizController := getQuizzesController()
	quizzes := quizController.GetQuizzes()
	c.Bind(&quizzes)
	return c.JSON(http.StatusOK, quizzes)
}

func questions(c echo.Context) error {
	_, err := utils.GetSessionCookie(c)
	if err != nil {
		return c.JSON(http.StatusUnauthorized, err)
	}

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

func users(c echo.Context) error {
	userController := getUsersController()
	users := userController.GetUser()
	c.Bind(&users)
	return c.JSON(http.StatusOK, users)
}

func answerHistories(c echo.Context) error {
	_, err := utils.GetSessionCookie(c)
	if err != nil {
		return c.JSON(http.StatusUnauthorized, err)
	}

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

func answers(c echo.Context) error {
	_, err := utils.GetSessionCookie(c)
	if err != nil {
		return c.JSON(http.StatusUnauthorized, err)
	}

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

func login(c echo.Context) error {
	req := new(LoginRequest)
	if err := c.Bind(req); err != nil {
		return c.JSON(http.StatusBadRequest, err)
	}

	sess, err := session.Get("session", c)
	if err != nil {
		return c.JSON(http.StatusUnauthorized, "Could not get session")
	}

	loginController := getLoginController()
	res := loginController.Handle(c, sess, req.Code)
	if res.Err != nil {
		return c.JSON(http.StatusBadRequest, res.Err)
	}

	return c.JSON(http.StatusOK, res.UserInfo.Name)
}

func logout(c echo.Context) error {
	sess, err := session.Get("session", c)
	if err != nil {
		return c.JSON(http.StatusUnauthorized, "Could not get session")
	}

	sess.Options = &sessions.Options{
		MaxAge: -1,
	}
	sess.Save(c.Request(), c.Response())

	return c.JSON(http.StatusOK, "logout")
}

func getQuizzesController() *controller.QuizController {
	return controller.NewQuizzesController(NewSqlHandler())
}

func getQuesionsController() *controller.QuestionController {
	return controller.NewQuesionsController(NewSqlHandler())
}

func getAnswerController() *controller.AnswerController {
	return controller.NewAnswerController(NewSqlHandler())
}

func getAnswerHistoryController() *controller.AnswerHistoryController {
	return controller.NewAnswerHistoryController(NewSqlHandler())
}

func getUsersController() *controller.UserController {
	return controller.NewUsersController(NewSqlHandler())
}

func getLoginController() *controller.LoginController {
	return controller.NewLoginController(NewSqlHandler(), NewRedisHandler())
}
