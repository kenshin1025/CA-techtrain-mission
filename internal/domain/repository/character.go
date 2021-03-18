package repository

import "ca-mission/internal/model"

type CharacterRepository interface {
	GetCharacterList(user *model.User) ([]*model.Chara, error)
}
