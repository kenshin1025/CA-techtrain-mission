package repository

import (
	"ca-mission/internal/domain/model"
	"context"
)

type CharaRepository interface {
	GetAllCharas(ctx context.Context) ([]*model.Chara, error)
}
