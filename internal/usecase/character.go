package usecase

import (
	"ca-mission/internal/domain/model"
	"ca-mission/internal/domain/repository"
)

type CharacterLister interface {
	GetUserCharacterList(user *model.User) ([]*model.UserCharaPossession, error)
}

type Character struct {
	charaRepo repository.CharacterRepository
}

func NewCharacter(charaRepo repository.CharacterRepository) CharacterLister {
	return &Character{
		charaRepo: charaRepo,
	}
}

func (c *Character) GetUserCharacterList(user *model.User) ([]*model.UserCharaPossession, error) {
	if err := c.charaRepo.GetUserID(user); err != nil {
		return nil, err
	}
	userCharas, err := c.charaRepo.GetCharacterList(user)
	if err != nil {
		return nil, err
	}
	return userCharas, nil
}
