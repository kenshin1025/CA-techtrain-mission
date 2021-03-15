package repository

import (
	"ca-mission/internal/model"
	"database/sql"
	"fmt"
)

type Gacha struct{}

func NewGacha() *Gacha {
	return &Gacha{}
}

func (g *Gacha) GetUserID(db *sql.DB, user *model.User) error {
	// トランザクション開始
	tx, err := db.Begin()
	if err != nil {
		return err
	}
	// tokenを元にユーザーのnameを取得
	execErr := tx.QueryRow("SELECT id FROM user WHERE token = ?", user.Token).Scan(&user.ID)
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

func (g *Gacha) SaveDrewCharas(db *sql.DB, user *model.User, charas []*model.Chara) error {
	query := "INSERT INTO user_chara_possession(user_id, chara_id) VALUES"
	for i, chara := range charas {
		if i != 0 {
			query += ","
		}
		query += fmt.Sprintf("(%d, %d)", user.ID, chara.ID)
	}
	// トランザクション開始
	tx, err := db.Begin()
	if err != nil {
		return err
	}
	// DBに追加
	//レコードを取得する必要のない、クエリはExecメソッドを使う。
	_, execErr := tx.Exec(query)
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
