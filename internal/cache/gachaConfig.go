package cache

import (
	"ca-mission/internal/domain/model"
	"ca-mission/internal/domain/repository"
)

type GachaConfigGenerater struct {
	charaRepo repository.CharaRepository
}

func NewGachaConfigGenerater(charaRepo repository.CharaRepository) *GachaConfigGenerater {
	return &GachaConfigGenerater{
		charaRepo: charaRepo,
	}
}

func (g *GachaConfigGenerater) GenerateGachaConfig() (*model.GachaConfig, error) {
	charas, err := g.charaRepo.GetAllCharas()
	if err != nil {
		return nil, err
	}

	gachaConfig := &model.GachaConfig{
		Charas:            charas,
		SumAllProbability: 0,
	}

	for _, chara := range charas {
		gachaConfig.SumAllProbability += chara.Probability
	}
	gachaConfig.Charas = charas

	return gachaConfig, nil
}
