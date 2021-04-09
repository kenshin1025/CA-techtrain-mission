package usecase

import (
	"ca-mission/internal/apierr"
	"ca-mission/internal/domain/model"
	"ca-mission/internal/domain/repository"
	"context"
	"math/rand"
	"time"
)

type GachaUsecase struct {
	userRepo    repository.UserRepository
	ucpRepo     repository.UserCharaPossessionRepository
	gachaConfig *model.GachaConfig
}

func NewGachaUsecase(userRepo repository.UserRepository, ucpRepo repository.UserCharaPossessionRepository, gachaConfig *model.GachaConfig) *GachaUsecase {
	return &GachaUsecase{
		userRepo:    userRepo,
		ucpRepo:     ucpRepo,
		gachaConfig: gachaConfig,
	}
}

func (u *GachaUsecase) Draw(ctx context.Context, times int, token string) ([]*model.Chara, error) {
	var charas []*model.Chara

	rand.Seed(time.Now().UnixNano())
	for i := 0; i < times; i++ {
		// 1からガチャ内のキャラ全てのProbabilityの合計の値までがランダムに出る
		chara, err := oneDraw(u.gachaConfig, rand.Intn(u.gachaConfig.SumAllProbability)+1)
		if err != nil {
			return nil, err
		}
		charas = append(charas, chara)
	}

	user, err := u.userRepo.GetByToken(ctx, token)
	if err != nil {
		return nil, err
	}

	if err := u.ucpRepo.SaveCharas(ctx, user, charas); err != nil {
		return nil, err
	}

	return charas, nil
}

func oneDraw(gachaConfig *model.GachaConfig, randN int) (*model.Chara, error) {
	boundary := 0
	for _, chara := range gachaConfig.Charas {
		boundary += chara.Probability
		if randN <= boundary {
			return chara, nil
		}
	}
	return nil, apierr.ErrInternalServerError
}
