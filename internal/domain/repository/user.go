package repository

import (
	"ca-mission/internal/domain/model"
)

type UserRepository interface {
	Create(m *model.User) error
	Get(m *model.User) error
	Update(m *model.User) error
}
