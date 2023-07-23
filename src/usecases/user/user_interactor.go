package user

import (
	entity "shortcut_master_api/src/domain"
)

type UserInteractor struct {
  UserRepository UserRepository
}

func (interactor *UserInteractor) Create(u entity.User) {
	interactor.UserRepository.Create(u)
}

func (interactor *UserInteractor) GetInfo() []entity.User {
	return interactor.UserRepository.Select()
}

func (interactor *UserInteractor) Delete(id string) {
	interactor.UserRepository.Delete(id)
}
