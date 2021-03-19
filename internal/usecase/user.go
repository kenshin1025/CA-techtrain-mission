package usecase

import (
	"ca-mission/internal/domain/model"
	"ca-mission/internal/domain/repository"
)

type UserInterface interface {
	Create(m *model.User) error
	Get(m *model.User) error
	Update(m *model.User) error
}

type User struct {
	userRepo repository.UserRepository
}

func NewUser(userRepo repository.UserRepository) UserInterface {
	return &User{
		userRepo: userRepo,
	}
}

func (u *User) Create(m *model.User) error {
	token, err := u.userRepo.GenerateUserToken()
	if err != nil {
		return err
	}
	m.Token = token
	err = u.userRepo.Create(m)
	if err != nil {
		return err
	}
	return nil
}

func (u *User) Get(m *model.User) error {
	err := u.userRepo.Get(m)
	if err != nil {
		return err
	}
	return nil
}

func (u *User) Update(m *model.User) error {
	err := u.userRepo.Update(m)
	if err != nil {
		return err
	}
	return nil
}
