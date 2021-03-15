package usecase

import (
	"ca-mission/internal/model"
	"database/sql"
)

type userRepository interface {
	GenerateUserToken() (string, error)
	Create(db *sql.DB, m *model.User) error
	Get(db *sql.DB, m *model.User) error
	Update(db *sql.DB, m *model.User) error
}

type User struct {
	userRepo userRepository
	db       *sql.DB
}

func NewUser(userRepo userRepository, db *sql.DB) *User {
	return &User{
		userRepo: userRepo,
		db:       db,
	}
}

func (u *User) Create(m *model.User) error {
	token, err := u.userRepo.GenerateUserToken()
	if err != nil {
		return err
	}
	m.Token = token
	err = u.userRepo.Create(u.db, m)
	if err != nil {
		return err
	}
	return nil
}

func (u *User) Get(m *model.User) error {
	err := u.userRepo.Get(u.db, m)
	if err != nil {
		return err
	}
	return nil
}

func (u *User) Update(m *model.User) error {
	err := u.userRepo.Update(u.db, m)
	if err != nil {
		return err
	}
	return nil
}
