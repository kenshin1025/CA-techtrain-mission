package repository

import (
	"ca-mission/internal/domain/model"
)

type UserRepository interface {
	Create(m *model.User) (int, error)
	GetByToken(token string) (*model.User, error)
	Update(m *model.User) error
}
