package repository

import "ca-mission/internal/domain/model"

type CharaRepository interface {
	GetAllCharas() ([]*model.Chara, error)
}
