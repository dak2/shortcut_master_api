package repositories

import (
	"errors"
	"fmt"
	entity "shortcut_master_api/src/domain"

	"gorm.io/gorm"
)

type UserRepository struct {
	SqlHandler SqlHandler
}

func (db *UserRepository) Create(u entity.User) (entity.User, error) {
	res := db.SqlHandler.Create(&u)
	if err := res.Error; err != nil {
		return entity.User{}, fmt.Errorf("Failed to create user")
	} else {
		return u, nil
	}
}

func (db *UserRepository) Select() []entity.User {
	user := []entity.User{}
	db.SqlHandler.FindAll(&user)
	return user
}

func (db *UserRepository) SelectByEmail(u entity.User) (entity.User, error) {
	user := entity.User{}
	res := db.SqlHandler.FindByParams(&user, "email", u.Email)
	if err := res.Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return entity.User{}, fmt.Errorf("Record not found")
		}
		return entity.User{}, fmt.Errorf("Failed to get user")
	}
	return user, nil
}

func (db *UserRepository) ExistsUserByGoogleUserId(googleUserId string) (bool) {
	user := entity.User{}
	res := db.SqlHandler.FindByParams(&user, "google_user_id", googleUserId)
	if err := res.Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return false
		}
		return false
	}
	return true
}

func (db *UserRepository) Delete(id string) {
	user := []entity.User{}
	db.SqlHandler.DeleteById(&user, id)
}
