package repository

import (
	"ca-mission/internal/domain/model"
	"context"
)

type UserRepository interface {
	Create(ctx context.Context, m *model.User) (int, error)
	GetByToken(ctx context.Context, token string) (*model.User, error)
	Update(ctx context.Context, m *model.User) error
}
