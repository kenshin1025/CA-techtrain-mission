package repository

import (
	"ca-mission/internal/domain/model"
	"ca-mission/internal/domain/repository"
	"database/sql"
)

type Character struct {
	db *sql.DB
}

func NewCharacter(db *sql.DB) repository.CharacterRepository {
	return &Character{
		db: db,
	}
}

func (c *Character) GetUserID(user *model.User) error {
	// トランザクション開始
	tx, err := c.db.Begin()
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

func (c *Character) GetCharacterList(user *model.User) ([]*model.UserCharaPossession, error) {
	rows, err := c.db.Query("SELECT user_chara_possession.id, chara.id, chara.name FROM user_chara_possession INNER JOIN chara ON user_chara_possession.chara_id=chara.id")
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
