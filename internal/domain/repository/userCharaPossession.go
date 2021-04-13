package repository

import (
	"ca-mission/internal/domain/model"
	"database/sql"
)

type UserCharaPossessionRepository interface {
	GetCharacterList(user *model.User) ([]*model.UserCharaPossession, error)
	// SaveCharas(user *model.User, charas []*model.Chara) error
	SaveCharas(tx *sql.Tx, user *model.User, charas []*model.Chara) error
}
