package repository

import "ca-mission/internal/domain/model"

type GachaConfigRepository interface {
	GetAllCharas() ([]*model.Chara, error)
}
