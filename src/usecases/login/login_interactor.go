package login

import (
	"context"
	"encoding/base64"
	"encoding/json"
	"fmt"
	"log"
	"os"
	entity "shortcut_master_api/src/domain"
	userUsecase "shortcut_master_api/src/usecases/user"

	"github.com/joho/godotenv"
	"golang.org/x/oauth2"
	"golang.org/x/oauth2/google"
)

type LoginInteractor struct {
	LoginRepository userUsecase.UserRepository
}

type GoogleUserInfo struct {
	GoogleUserId  string `json:"sub"`
	Name          string `json:"name"`
	Email         string `json:"email"`
	EmailVerified bool   `json:"email_verified"`
}

type GoogleUserResult struct {
	UserInfo GoogleUserInfo
	Err      error
}

func (interactor *LoginInteractor) GetUser(code string) (entity.User, error) {
	res := GetGoogleUser(code)
	if res.Err != nil {
		return entity.User{}, res.Err
	}

	u := entity.User{
		GoogleUserId: res.UserInfo.GoogleUserId,
		Name:         res.UserInfo.Name,
		Email:        res.UserInfo.Email,
	}

	user, err := interactor.GetUserByEmail(u)
	if err != nil {
		if err.Error() == "Record not found" {
			u, err := interactor.CreateUser(u)
			if err != nil {
				return entity.User{}, err
			}
			return u, nil
		}
		return entity.User{}, err
	}

	return user, nil
}

func (interactor *LoginInteractor) GetUserByEmail(u entity.User) (entity.User, error) {
	user, err := interactor.LoginRepository.SelectByEmail(u)
	if err != nil {
		return entity.User{}, err
	}
	return user, nil
}

func (interactor *LoginInteractor) CreateUser(u entity.User) (entity.User, error) {
	user, err := interactor.LoginRepository.Create(u)
	if err != nil {
		return entity.User{}, err
	}
	return user, nil
}

func InitGoogleOAuthConfig() (*oauth2.Config, error) {
	err := godotenv.Load(".env")

	if err != nil {
		log.Fatal("Error loading .env file")
		return nil, err
	}

	c := &oauth2.Config{
		ClientID:     os.Getenv("GOOGLE_CLIENT_ID"),
		ClientSecret: os.Getenv("GOOGLE_CLIENT_SECRET"),
		Endpoint:     google.Endpoint,
		RedirectURL:  os.Getenv("GOOGLE_REDIRECT_URI"),
		Scopes: []string{
			"https://www.googleapis.com/auth/userinfo.email",
			"https://www.googleapis.com/auth/userinfo.profile",
		},
	}

	return c, nil
}

func GetGoogleUser(code string) GoogleUserResult {
	dec, err := base64.StdEncoding.DecodeString(code)
	if err != nil {
		return GoogleUserResult{Err: fmt.Errorf("Failed to decode token: %w", err)}
	}

	conf, err := InitGoogleOAuthConfig()
	if err != nil {
		return GoogleUserResult{Err: fmt.Errorf("Failed to initialize config: %w", err)}
	}

	token, err := conf.Exchange(context.TODO(), string(dec))
	if err != nil {
		return GoogleUserResult{Err: fmt.Errorf("Failed to get token: %w", err)}
	}

	cl := conf.Client(context.Background(), token)
	res, err := cl.Get("https://www.googleapis.com/oauth2/v3/userinfo")
	if err != nil {
		return GoogleUserResult{Err: fmt.Errorf("Failed to get userinfo: %w", err)}
	}
	defer res.Body.Close()

	var userInfo GoogleUserInfo
	err = json.NewDecoder(res.Body).Decode(&userInfo)
	if err != nil {
		return GoogleUserResult{Err: fmt.Errorf("Failed to map userinfo: %w", err)}
	}

	return GoogleUserResult{UserInfo: userInfo, Err: nil}
}
