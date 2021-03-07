package usecase

import (
	"ca-mission/internal/model"
	"database/sql"
)

type userRepository interface {
	GenerateUserToken() (string, error)
	Create(db *sql.DB, m *model.User) error
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
