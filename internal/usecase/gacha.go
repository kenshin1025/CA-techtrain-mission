package usecase

import (
	"ca-mission/internal/apierr"
	"ca-mission/internal/config"
	"ca-mission/internal/model"
	"database/sql"
	"math/rand"
	"time"
)

type gachaRepository interface {
	GetUserID(db *sql.DB, user *model.User) error
	SaveDrewCharas(db *sql.DB, user *model.User, charas []*model.Chara) error
}

type Gacha struct {
	gachaRepo   gachaRepository
	db          *sql.DB
	gachaConfig *config.GachaConfig
}

func NewGacha(gachaRepo gachaRepository, db *sql.DB, gachaConfig *config.GachaConfig) *Gacha {
	return &Gacha{
		gachaRepo:   gachaRepo,
		db:          db,
		gachaConfig: gachaConfig,
	}
}

func (g *Gacha) Draw(times int, token string) ([]*model.Chara, error) {
	var charas []*model.Chara

	rand.Seed(time.Now().UnixNano())
	for i := 0; i < times; i++ {
		// 1からガチャ内のキャラ全てのProbabilityの合計の値までがランダムに出る
		chara, err := oneDraw(g.db, g.gachaConfig, rand.Intn(g.gachaConfig.SumAllProbability)+1)
		if err != nil {
			return nil, err
		}
		charas = append(charas, chara)
	}

	user := &model.User{
		Token: token,
	}

	if err := g.gachaRepo.GetUserID(g.db, user); err != nil {
		return nil, err
	}

	if err := g.gachaRepo.SaveDrewCharas(g.db, user, charas); err != nil {
		return nil, err
	}

	return charas, nil
}

func oneDraw(db *sql.DB, gachaConfig *config.GachaConfig, randN int) (*model.Chara, error) {
	boundary := 0
	for _, chara := range gachaConfig.Charas {
		boundary += chara.Probability
		if randN <= boundary {
			return chara, nil
		}
	}
	return nil, apierr.ErrInternalServerError
}
