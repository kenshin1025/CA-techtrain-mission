package usecase

import (
	"ca-mission/internal/domain/model"
	"ca-mission/internal/domain/repository"

	"github.com/google/uuid"
)

type UserInterface interface {
	Create(m *model.User) error
	Get(m *model.User) error
	Update(m *model.User) error
}

type User struct {
	userRepo repository.UserRepository
}

func NewUser(userRepo repository.UserRepository) UserInterface {
	return &User{
		userRepo: userRepo,
	}
}

func (u *User) Create(m *model.User) error {
	token, err := GenerateUserToken()
	if err != nil {
		return err
	}
	m.Token = token
	err = u.userRepo.Create(m)
	if err != nil {
		return err
	}
	return nil
}

func (u *User) Get(m *model.User) error {
	err := u.userRepo.Get(m)
	if err != nil {
		return err
	}
	return nil
}

func (u *User) Update(m *model.User) error {
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
