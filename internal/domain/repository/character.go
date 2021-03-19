package repository

import "ca-mission/internal/domain/model"

type CharacterRepository interface {
	GetUserID(user *model.User) error
	GetCharacterList(user *model.User) ([]*model.Chara, error)
}
