package repository

import (
	"ca-mission/internal/model"
)

type GachaRepository interface {
	GetUserID(user *model.User) error
	SaveDrewCharas(user *model.User, charas []*model.Chara) error
}
