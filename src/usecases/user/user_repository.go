package user

import entity "shortcut_master_api/src/domain"

// MEMO : for dip

type UserRepository interface {
	Store(entity.User)
	Select() []entity.User
	SelectByEmail(entity.User) entity.User
	Delete(id string)
}
