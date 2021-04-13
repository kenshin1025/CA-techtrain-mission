package repository

import (
	"ca-mission/internal/apierr"
	"ca-mission/internal/domain/model"
	"ca-mission/internal/domain/repository"
	"context"
	"database/sql"
	"log"
)

type UserRepository struct {
	db *sql.DB
}

func NewUserRepository(db *sql.DB) repository.UserRepository {
	return &UserRepository{
		db: db,
	}
}

// 与えられたモデルからユーザー作成する関数
func (r *UserRepository) Create(ctx context.Context, user *model.User) (int, error) {
	// DBに追加
	//レコードを取得する必要のない、クエリはExecメソッドを使う。
	result, execErr := r.db.ExecContext(ctx, "INSERT INTO user(name, token) VALUES(?,?)", user.Name, user.Token)
	//エラーが起きたらロールバック
	if execErr != nil {
		return 0, execErr
	}

	id, err := result.LastInsertId()
	if err != nil {
		return 0, execErr
	}

	return int(id), nil
}

func (r *UserRepository) GetByToken(ctx context.Context, token string) (*model.User, error) {
	user := model.User{
		Token: token,
	}
	// tokenを元にユーザーのnameを取得
	execErr := r.db.QueryRowContext(ctx, "SELECT id, name FROM user WHERE token = ?", token).Scan(&user.ID, &user.Name)
	if execErr == sql.ErrNoRows {
		return nil, apierr.ErrUserNotExists
	} else if execErr != nil {
		return nil, execErr
	}
	return &user, nil
}

func (r *UserRepository) Update(ctx context.Context, user *model.User) error {
	// トランザクション開始
	tx, err := r.db.Begin()
	if err != nil {
		return err
	}
	// tokenを元にユーザーのnameを更新
	_, execErr := tx.ExecContext(ctx, "UPDATE user SET name = ? WHERE token = ?", user.Name, user.Token)
	if execErr != nil {
		_ = tx.Rollback()
		return execErr
	}
	// エラーが起きなければコミット
	err = tx.Commit()
	if err != nil {
		log.Fatal(err)
	}
	return nil
}
