package infrastructure

import (
	"net/http"
	controller "shortcut_master_api/src/interfaces/controllers"
	"shortcut_master_api/src/utils"
	"strconv"

	"github.com/gorilla/sessions"
	"github.com/labstack/echo-contrib/session"
	"github.com/labstack/echo/v4"
)

func hello(c echo.Context) error {
	_, err := utils.GetSessionCookie(c)
	if err != nil {
		return err
	}

	return c.JSON(http.StatusOK, "Hello, World!")
}

func quizzes(c echo.Context) error {
	_, err := utils.GetSessionCookie(c)
	if err != nil {
		return err
	}

	quizController := getQuizzesController()
	quizzes := quizController.GetQuizzes()
	c.Bind(&quizzes)
	return c.JSON(http.StatusOK, quizzes)
}

func questions(c echo.Context) error {
	_, err := utils.GetSessionCookie(c)
	if err != nil {
		return err
	}

	quizIdStr := c.QueryParam("quiz_id")
	quizId, err := strconv.Atoi(quizIdStr)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err)
	}

	questionController := getQuesionsController()
	questions := questionController.GetQuestionsByQuiz(quizId)
	if questions.Err != nil {
		return c.JSON(http.StatusInternalServerError, questions.Err)
	}

	c.Bind(&questions.Questions)
	return c.JSON(http.StatusOK, questions)
}

func users(c echo.Context) error {
	userController := getUsersController()
	users := userController.GetUser()
	c.Bind(&users)
	return c.JSON(http.StatusOK, users)
}

func login(c echo.Context) error {
	req := new(LoginRequest)
	if err := c.Bind(req); err != nil {
		return c.JSON(http.StatusBadRequest, err)
	}

	sess, err := session.Get("session", c)
	if err != nil {
		return c.JSON(http.StatusBadRequest, "Could not get session")
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
		return c.JSON(http.StatusBadRequest, "Could not get session")
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

func getUsersController() *controller.UserController {
	return controller.NewUsersController(NewSqlHandler())
}

func getLoginController() *controller.LoginController {
	return controller.NewLoginController(NewSqlHandler())
}
