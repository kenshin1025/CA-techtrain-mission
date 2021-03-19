package cache

import "ca-mission/internal/domain/model"

type GachaConfig struct {
	SumAllProbability int
	Charas            []*model.Chara
}
