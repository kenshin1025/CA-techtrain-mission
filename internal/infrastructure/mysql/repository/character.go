package repository

import (
	"ca-mission/internal/domain/repository"
	"ca-mission/internal/model"
	"database/sql"
	"fmt"
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

func (c *Character) GetCharacterList(user *model.User) ([]*model.Chara, error) {
	rows, err := c.db.Query("SELECT user_chara_possession.id, chara.id, chara.name FROM user_chara_possession INNER JOIN chara ON user_chara_possession.chara_id=chara.id")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var charas []*model.Chara

	for rows.Next() {
		var userCharacterID int
		var c model.Chara
		if err := rows.Scan(&userCharacterID, &c.ID, &c.Name); err != nil {
			return nil, err
		}
		fmt.Println(userCharacterID)
		charas = append(charas, &c)
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return charas, nil
}
