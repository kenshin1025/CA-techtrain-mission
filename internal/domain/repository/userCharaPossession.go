package repository

import (
	"ca-mission/internal/domain/model"
	"context"
)

type UserCharaPossessionRepository interface {
	GetCharacterList(ctx context.Context, user *model.User) ([]*model.UserCharaPossession, error)
	SaveCharas(ctx context.Context, user *model.User, charas []*model.Chara) error
}
