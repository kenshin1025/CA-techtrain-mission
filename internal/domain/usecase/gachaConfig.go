package usecase

import (
	"ca-mission/internal/cache"
	"ca-mission/internal/domain/repository"
)

type GachaConfigGenerater interface {
	GenerateGachaConfig() (*cache.GachaConfig, error)
}

type GachaConfig struct {
	gachaConfRepo repository.GachaConfigRepository
}

func NewGachaConfig(gachaConfRepo repository.GachaConfigRepository) GachaConfigGenerater {
	return &GachaConfig{
		gachaConfRepo: gachaConfRepo,
	}
}

func (g *GachaConfig) GenerateGachaConfig() (*cache.GachaConfig, error) {
	charas, err := g.gachaConfRepo.GetAllCharas()
	if err != nil {
		return nil, err
	}

	gachaConfig := &cache.GachaConfig{
		Charas:            charas,
		SumAllProbability: 0,
	}

	for _, chara := range charas {
		gachaConfig.SumAllProbability += chara.Probability
	}
	gachaConfig.Charas = charas

	return gachaConfig, nil
}
