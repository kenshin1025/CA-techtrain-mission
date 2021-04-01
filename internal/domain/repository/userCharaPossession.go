package repository

import "ca-mission/internal/domain/model"

type UserCharaPossessionRepository interface {
	GetCharacterList(user *model.User) ([]*model.UserCharaPossession, error)
	SaveCharas(user *model.User, charas []*model.Chara) error
}
