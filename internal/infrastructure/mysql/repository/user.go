package repository

import (
	"ca-mission/internal/apierr"
	"ca-mission/internal/domain/model"
	"ca-mission/internal/domain/repository"
	"database/sql"
	"log"
)

type User struct {
	db *sql.DB
}

func NewUser(db *sql.DB) repository.UserRepository {
	return &User{
		db: db,
	}
}

// 与えられたモデルからユーザー作成する関数
func (u *User) Create(user *model.User) (int, error) {
	// DBに追加
	//レコードを取得する必要のない、クエリはExecメソッドを使う。
	result, execErr := u.db.Exec("INSERT INTO user(name, token) VALUES(?,?)", user.Name, user.Token)
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

func (u *User) GetByToken(token string) (*model.User, error) {
	user := model.User{
		Token: token,
	}
	// tokenを元にユーザーのnameを取得
	execErr := u.db.QueryRow("SELECT id, name FROM user WHERE token = ?", token).Scan(&user.ID, &user.Name)
	if execErr == sql.ErrNoRows {
		return nil, apierr.ErrUserNotExists
	} else if execErr != nil {
		return nil, execErr
	}
	return &user, nil
}

func (u *User) Update(user *model.User) error {
	// トランザクション開始
	tx, err := u.db.Begin()
	if err != nil {
		return err
	}
	// tokenを元にユーザーのnameを更新
	_, execErr := tx.Exec("UPDATE user SET name = ? WHERE token = ?", user.Name, user.Token)
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
