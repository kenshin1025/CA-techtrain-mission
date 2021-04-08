package usecase

import (
	"ca-mission/internal/domain/model"
	"ca-mission/internal/domain/repository"
	"context"

	"github.com/google/uuid"
)

type UserUsecase struct {
	userRepo repository.UserRepository
}

func NewUserUsecase(userRepo repository.UserRepository) *UserUsecase {
	return &UserUsecase{
		userRepo: userRepo,
	}
}

func (u *UserUsecase) Create(ctx context.Context, m *model.User) error {
	token, err := GenerateUserToken()
	if err != nil {
		return err
	}
	m.Token = token
	_, err = u.userRepo.Create(ctx, m)
	if err != nil {
		return err
	}
	return nil
}

func (u *UserUsecase) GetByToken(ctx context.Context, token string) (*model.User, error) {
	user, err := u.userRepo.GetByToken(ctx, token)
	if err != nil {
		return nil, err
	}
	return user, nil
}

func (u *UserUsecase) Update(ctx context.Context, m *model.User) error {
	err := u.userRepo.Update(ctx, m)
	if err != nil {
		return err
	}
	return nil
}

// uuidを生成して返す関数
func GenerateUserToken() (string, error) {
	//生成したuuidが被っていないかチェックするようにした方が良いかも
	uuid, err := uuid.NewRandom()
	return uuid.String(), err
}
