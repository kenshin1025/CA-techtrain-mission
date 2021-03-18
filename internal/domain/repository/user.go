package repository

import (
	"ca-mission/internal/model"
)

type UserRepository interface {
	GenerateUserToken() (string, error)
	Create(m *model.User) error
	Get(m *model.User) error
	Update(m *model.User) error
}
