package usecase

import (
	"ca-mission/internal/apierr"
	"ca-mission/internal/cache"
	"ca-mission/internal/domain/model"
	"ca-mission/internal/domain/repository"
	"math/rand"
	"time"
)

type Drawer interface {
	Draw(times int, token string) ([]*model.Chara, error)
}

type Gacha struct {
	gachaRepo   repository.GachaRepository
	gachaConfig *cache.GachaConfig
}

func NewGacha(gachaRepo repository.GachaRepository, gachaConfig *cache.GachaConfig) Drawer {
	return &Gacha{
		gachaRepo:   gachaRepo,
		gachaConfig: gachaConfig,
	}
}

func (g *Gacha) Draw(times int, token string) ([]*model.Chara, error) {
	var charas []*model.Chara

	rand.Seed(time.Now().UnixNano())
	for i := 0; i < times; i++ {
		// 1からガチャ内のキャラ全てのProbabilityの合計の値までがランダムに出る
		chara, err := oneDraw(g.gachaConfig, rand.Intn(g.gachaConfig.SumAllProbability)+1)
		if err != nil {
			return nil, err
		}
		charas = append(charas, chara)
	}

	user := &model.User{
		Token: token,
	}

	if err := g.gachaRepo.GetUserID(user); err != nil {
		return nil, err
	}

	if err := g.gachaRepo.SaveDrewCharas(user, charas); err != nil {
		return nil, err
	}

	return charas, nil
}

func oneDraw(gachaConfig *cache.GachaConfig, randN int) (*model.Chara, error) {
	boundary := 0
	for _, chara := range gachaConfig.Charas {
		boundary += chara.Probability
		if randN <= boundary {
			return chara, nil
		}
	}
	return nil, apierr.ErrInternalServerError
}
