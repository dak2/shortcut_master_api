package database

import (
	entity "shortcut_master_api/src/entities"
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
func (db *UserRepository) Delete(id string) {
	user := []entity.User{}
	db.DeleteById(&user, id)
}
