package controllers

import (
	"encoding/base64"
)

type GoogleUserInfo struct {
	Sub           string `json:"sub"`   // ユーザーの一意のID
	Name          string `json:"name"`  // ユーザーの表示名
	Email         string `json:"email"` // ユーザーのメールアドレス
	Picture       string `json:"picture"`
	EmailVerified bool   `json:"email_verified"`
}

func LoginController(credential string) string {
	// GetGoogleUserInfo(credential)
	return "success"
}

func GetGoogleUserInfo(credential string) (string, error) {
	// TODO: implement
	dec, err := base64.StdEncoding.DecodeString(credential)
	if err != nil {
		return "", err
	}
	token := string(dec)
	return token, nil
}
