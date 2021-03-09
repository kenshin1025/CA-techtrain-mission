package repository

import (
	"ca-mission/internal/apierr"
	"ca-mission/internal/model"
	"database/sql"

	"github.com/google/uuid"
)

type User struct{}

func NewUser() *User {
	return &User{}
}

// uuidを生成して返す関数
func (u *User) GenerateUserToken() (string, error) {
	//生成したuuidが被っていないかチェックするようにした方が良いかも
	uuid, err := uuid.NewRandom()
	return uuid.String(), err
}

// 与えられたモデルからユーザー作成する関数
func (u *User) Create(db *sql.DB, m *model.User) error {
	// トランザクション開始
	tx, err := db.Begin()
	if err != nil {
		return err
	}
	// DBに追加
	//レコードを取得する必要のない、クエリはExecメソッドを使う。
	_, execErr := tx.Exec("INSERT INTO user(name, token) VALUES(?,?)", m.Name, m.Token)
	//エラーが起きたらロールバック
	if execErr != nil {
		_ = tx.Rollback()
		return execErr
	}
	// エラーが起きなければコミット
	err = tx.Commit()
	if err != nil {
		return err
	}
	return nil
}

func (u *User) Get(db *sql.DB, m *model.User) error {
	// トランザクション開始
	tx, err := db.Begin()
	if err != nil {
		return err
	}
	// tokenを元にユーザーのnameを取得
	execErr := tx.QueryRow("SELECT name FROM user WHERE token = ?", m.Token).Scan(&m.Name)
	if execErr == sql.ErrNoRows {
		return apierr.ErrUserNotExists
	} else if execErr != nil {
		_ = tx.Rollback()
		return execErr
	}
	// エラーが起きなければコミット
	err = tx.Commit()
	if err != nil {
		return err
	}
	return nil
}
