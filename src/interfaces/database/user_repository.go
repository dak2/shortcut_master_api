package database

import (
	"errors"
	"fmt"
	entity "shortcut_master_api/src/domain"

	"gorm.io/gorm"
)

type UserRepository struct {
	SqlHandler
}

func (db *UserRepository) Store(u entity.User) {
	db.Create(&u)
}

func (db *UserRepository) Select() []entity.User {
	user := []entity.User{}
	db.FindAll(&user)
	return user
}

func (db *UserRepository) SelectByEmail(u entity.User) (entity.User, error) {
	user := entity.User{}
	res := db.FindByParams(&user, "email", u.Email)
	if err := res.Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return entity.User{}, fmt.Errorf("User not found")
		}
		return entity.User{}, fmt.Errorf("Failed to get user")
	}
	return user, nil
}

func (db *UserRepository) Delete(id string) {
	user := []entity.User{}
	db.DeleteById(&user, id)
}
