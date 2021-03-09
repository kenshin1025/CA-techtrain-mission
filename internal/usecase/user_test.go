package usecase

import (
	"ca-mission/internal/model"
	"database/sql"
	"testing"
)

type userRepositoryMock struct {
	generateUserTokenFn func() (string, error)
	createFn            func(db *sql.DB, m *model.User) error
}

func (s *userRepositoryMock) GenerateUserToken() (string, error) {
	return s.generateUserTokenFn()
}

func (s *userRepositoryMock) Create(db *sql.DB, m *model.User) error {
	return s.createFn(db, m)
}

func TestUser_Create(t *testing.T) {
	mock := &userRepositoryMock{
		generateUserTokenFn: func() (string, error) {
			return "aaaaaaaa-aaaa-aaaa-aaaa-aaaaaaaaaaaa", nil
		},
		createFn: func(db *sql.DB, m *model.User) error {
			return nil
		},
	}
	sut := NewUser(mock, nil)

	if err := sut.Create(&model.User{
		Name: "test_name",
	}); err != nil {
		t.Fatal(err)
	}
}
