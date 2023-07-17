package user

import entity "shortcut_master_api/src/domain"

// MEMO : for dip

type UserRepository interface {
	Create(entity.User) (entity.User, error)
	Select() []entity.User
	SelectByEmail(entity.User) (entity.User, error)
	Delete(id string)
}
