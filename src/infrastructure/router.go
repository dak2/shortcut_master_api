package infrastructure

import (
	"context"
	"encoding/base64"
	"fmt"
	"log"
	"os"
	"io/ioutil"
	"net/http"
	controller "shortcut_master_api/src/interfaces/controllers"
	"github.com/joho/godotenv"
	"github.com/labstack/echo/v4"
	"golang.org/x/oauth2"
	"golang.org/x/oauth2/google"
)

type LoginRequest struct {
	Code string `json:"code"`
}

func hello(c echo.Context) error {
	return c.String(http.StatusOK, "Hello, World!")
}

func Handle(e *echo.Echo) {
	err := godotenv.Load(".env")

	if err != nil {
		log.Fatal("Error loading .env file")
	}
	userController := controller.NewUsersController(NewSqlHandler())
	e.GET("/", hello)

	// -- users -- //
	e.GET("/users", func(c echo.Context) error {
		users := userController.GetUser()
		c.Bind(&users)
		return c.JSON(http.StatusOK, users)
	})

	// -- login -- //
	e.POST("/login", func(c echo.Context) error {
		req := new(LoginRequest)
		if err := c.Bind(req); err != nil {
			return err
		}
		loginController := controller.LoginController(req.Code)
		c.Bind(&loginController)
		dec, err := base64.StdEncoding.DecodeString(req.Code)
		if err != nil {
			return err
		}
		fmt.Println(string(dec))

		config := &oauth2.Config{
			ClientID:     os.Getenv("GOOGLE_CLIENT_ID"),
			ClientSecret: os.Getenv("GOOGLE_CLIENT_SECRET"),
			Endpoint:     google.Endpoint,
			RedirectURL:  os.Getenv("GOOGLE_REDIRECT_URI"),
			Scopes: []string{
				"https://www.googleapis.com/auth/userinfo.email",
				"https://www.googleapis.com/auth/userinfo.profile",
			},
		}

		token, err := config.Exchange(context.TODO(), string(dec))
		if err != nil {
			fmt.Printf("トークンの取得エラー: %v\n", err)
			return err
		}

		// ユーザー情報を取得
		client := config.Client(context.Background(), token)
		resp, err := client.Get("https://www.googleapis.com/oauth2/v3/userinfo")
		if err != nil {
			fmt.Printf("ユーザー情報の取得エラー: %v\n", err)
			return err
		}
		defer resp.Body.Close()

		body, err := ioutil.ReadAll(resp.Body)
		if err != nil {
			fmt.Printf("ボディの読み取りエラー: %v\n", err)
			return err
		}

		fmt.Printf("ユーザー情報: %s\n", body)

		return c.JSON(http.StatusOK, "loginController")
	})
}
