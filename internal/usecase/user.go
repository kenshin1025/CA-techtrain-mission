package usecase

import (
	"ca-mission/internal/domain/model"
	"ca-mission/internal/domain/repository"

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

func (u *UserUsecase) Create(m *model.User) error {
	token, err := GenerateUserToken()
	if err != nil {
		return err
	}
	m.Token = token
	_, err = u.userRepo.Create(m)
	if err != nil {
		return err
	}
	return nil
}

func (u *UserUsecase) GetByToken(token string) (*model.User, error) {
	user, err := u.userRepo.GetByToken(token)
	if err != nil {
		return nil, err
	}
	return user, nil
}

func (u *UserUsecase) Update(m *model.User) error {
	err := u.userRepo.Update(m)
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
