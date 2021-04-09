package usecase

import (
	"ca-mission/internal/domain/model"
	"ca-mission/internal/domain/repository"
	"context"
)

type CharacterUsecase struct {
	userRepo repository.UserRepository
	ucpRepo  repository.UserCharaPossessionRepository
}

func NewCharacterUsecase(userRepo repository.UserRepository, ucpRepo repository.UserCharaPossessionRepository) *CharacterUsecase {
	return &CharacterUsecase{
		userRepo: userRepo,
		ucpRepo:  ucpRepo,
	}
}

func (c *CharacterUsecase) GetUsersCharaListByToken(ctx context.Context, token string) ([]*model.UserCharaPossession, error) {
	user, err := c.userRepo.GetByToken(ctx, token)
	if err != nil {
		return nil, err
	}
	userCharas, err := c.ucpRepo.GetCharacterList(ctx, user)
	if err != nil {
		return nil, err
	}
	return userCharas, nil
}
