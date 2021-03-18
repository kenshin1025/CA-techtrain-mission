package repository

import "ca-mission/internal/model"

type GachaConfigRepository interface {
	GetAllCharas() ([]*model.Chara, error)
}
