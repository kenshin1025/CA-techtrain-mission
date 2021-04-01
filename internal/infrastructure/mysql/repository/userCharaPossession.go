package repository

import (
	"ca-mission/internal/domain/model"
	"ca-mission/internal/domain/repository"
	"database/sql"
	"fmt"
)

type UserCharaPossessionRepository struct {
	db *sql.DB
}

func NewUserCharaPossessionRepository(db *sql.DB) repository.UserCharaPossessionRepository {
	return &UserCharaPossessionRepository{
		db: db,
	}
}

func (r *UserCharaPossessionRepository) GetCharacterList(user *model.User) ([]*model.UserCharaPossession, error) {
	rows, err := r.db.Query("SELECT user_chara_possession.id, chara.id, chara.name FROM user_chara_possession INNER JOIN chara ON user_chara_possession.chara_id=chara.id")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var userCharas []*model.UserCharaPossession

	for rows.Next() {
		var uc model.UserCharaPossession
		if err := rows.Scan(&uc.ID, &uc.Chara.ID, &uc.Chara.Name); err != nil {
			return nil, err
		}
		userCharas = append(userCharas, &uc)
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return userCharas, nil
}

func (r *UserCharaPossessionRepository) SaveCharas(user *model.User, charas []*model.Chara) error {
	query := "INSERT INTO user_chara_possession(user_id, chara_id) VALUES"
	for i, chara := range charas {
		if i != 0 {
			query += ","
		}
		query += fmt.Sprintf("(%d, %d)", user.ID, chara.ID)
	}
	// トランザクション開始
	tx, err := r.db.Begin()
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
