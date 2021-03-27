package usecase

import (
	"ca-mission/internal/apierr"
	"ca-mission/internal/domain/model"
	"errors"
	"testing"
)

type userRepositoryMock struct {
	createFn func(m *model.User) error
	getFn    func(m *model.User) error
	updateFn func(m *model.User) error
}

func (s *userRepositoryMock) Create(m *model.User) error {
	return s.createFn(m)
}

func (s *userRepositoryMock) Get(m *model.User) error {
	return s.getFn(m)
}

func (s *userRepositoryMock) Update(m *model.User) error {
	return s.updateFn(m)
}

//ユーザー作成成功ケース
func TestUser_Create(t *testing.T) {
	mock := &userRepositoryMock{
		createFn: func(m *model.User) error {
			return nil
		},
	}
	sut := NewUser(mock)

	if err := sut.Create(&model.User{
		Name: "test_name",
	}); err != nil {
		t.Fatal(err)
	}
}

//ユーザー作成失敗ケース
func TestUser_Create_Failed(t *testing.T) {
	mock := &userRepositoryMock{
		createFn: func(m *model.User) error {
			return apierr.ErrInternalServerError
		},
	}
	sut := NewUser(mock)
	err := sut.Create(&model.User{
		Name: "test_name",
	})
	if !errors.Is(err, apierr.ErrInternalServerError) {
		t.Errorf("error must be %v but %v", apierr.ErrInternalServerError, err)
	}
}

//ユーザー取得成功ケース
func TestUser_Get(t *testing.T) {
	mock := &userRepositoryMock{
		getFn: func(m *model.User) error {
			return nil
		},
	}
	sut := NewUser(mock)

	if err := sut.Get(&model.User{
		Token: "aaaaaaaa-aaaa-aaaa-aaaa-aaaaaaaaaaaa",
	}); err != nil {
		t.Fatal(err)
	}
}

//ユーザーが存在しなかったケース
func TestUser_Get_NotExistToken(t *testing.T) {
	mock := &userRepositoryMock{
		getFn: func(m *model.User) error {
			return apierr.ErrUserNotExists
		},
	}
	sut := NewUser(mock)

	if err := sut.Get(&model.User{
		Token: "aaaaaaaa-aaaa-aaaa-aaaa-aaaaaaaaaaaa",
	}); !errors.Is(err, apierr.ErrUserNotExists) {
		t.Errorf("error must be %v but %v", apierr.ErrUserNotExists, err)
	}
}

//ユーザー取得失敗ケース
func TestUser_Get_Failed(t *testing.T) {
	mock := &userRepositoryMock{
		getFn: func(m *model.User) error {
			return apierr.ErrInternalServerError
		},
	}
	sut := NewUser(mock)

	if err := sut.Get(&model.User{
		Token: "aaaaaaaa-aaaa-aaaa-aaaa-aaaaaaaaaaaa",
	}); !errors.Is(err, apierr.ErrInternalServerError) {
		t.Errorf("error must be %v but %v", apierr.ErrInternalServerError, err)
	}
}
