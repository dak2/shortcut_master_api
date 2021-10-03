package controllers

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	entity "short_cut_master_api/src/entities"
	repository "short_cut_master_api/src/interfaces/database"
	usecase "short_cut_master_api/src/usecases"

	"github.com/labstack/echo"
)

type UserController struct {
	Interactor usecase.UserInteractor
}

type JWT struct {
  Iss            string  `json:"iss"`
  Azp            string  `json:"azp"`
  Aud            string  `json:"aud"`
  Sub            string  `json:"sub"`
  Email          string  `json:"email"`
  EmailVerified  string  `json:"email_verified"`
  AtHash         string  `json:"at_hash"`
  Name           string  `json:"name"`
  Picture        string  `json:"picture"`
  GivenName      string  `json:"given_name"`
  FamilyName     string  `json:"family_name"`
  Locale         string  `json:"locale"`
  Iat            string  `json:"iat"`
  Exp            string  `json:"exp"`
  Jti            string  `json:"jti"`
  Alg            string  `json:"alg"`
  Kid            string  `json:"kid"`
  Typ            string  `json:"typ"`
}

func NewUsersController(sqlHandler repository.SqlHandler) *UserController {
	return &UserController {
		Interactor: usecase.UserInteractor {
			UserRepository: &repository.UserRepository {
				SqlHandler: sqlHandler,
			},
		},
	}
}

func (controller *UserController) Create(c echo.Context) {
	u := entity.User{}
	c.Bind(&u)
	controller.Interactor.Add(u)
	createdUsers := controller.Interactor.GetInfo()
	c.JSON(201, createdUsers)
	return
}

func (controller *UserController) GetUser() []entity.User {
	res := controller.Interactor.GetInfo()
	return res
}

func (controller *UserController) Delete(id string) {
	controller.Interactor.Delete(id)
}

func (controller *UserController) Login(token string) **JWT {
	jwt, err := fetchJWT(token)
	if err != nil {
		fmt.Println(jwt)
	}
	// TODO_1: user_id（jwtのsub）でDB検索してあればそのユーザーとセッション返す
	// TODO_2: user_id（jwtのsub）でDB検索して無ければ作成する
	// TODO_3: テーブルにexternal_user_idを保存（hash化しておく）
	return &jwt
}

func fetchJWT(token string) (*JWT, error) {
	endpoint := fmt.Sprintf("https://www.googleapis.com/oauth2/v3/tokeninfo?id_token=%s", token)
	resp, err := http.Get(endpoint)

	if err != nil {
		return nil, fmt.Errorf("get request failed because of %s", err)
	}

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("bad response status code %d", resp.StatusCode)
	}

	var jwt JWT
	if resp != nil {
		defer resp.Body.Close()
		res, err := ioutil.ReadAll(resp.Body)

		if err != nil {
			return nil, fmt.Errorf("without jwt %s", err)
		}
		json.Unmarshal(res, &jwt)
	}

	return &jwt, nil
}
