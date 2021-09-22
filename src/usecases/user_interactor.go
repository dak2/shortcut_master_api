package usecase

import (
	entity "short_cut_master_api/src/entities"
)

type UserInteractor struct {
  UserRepository UserRepository
}

func (interactor *UserInteractor) Add(u entity.User) {
	interactor.UserRepository.Store(u)
}

func (interactor *UserInteractor) GetInfo() []entity.User {
	return interactor.UserRepository.Select()
}

func (interactor *UserInteractor) Delete(id string) {
	interactor.UserRepository.Delete(id)
}
