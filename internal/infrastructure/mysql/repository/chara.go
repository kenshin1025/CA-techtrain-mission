package repository

import (
	"ca-mission/internal/domain/model"
	"ca-mission/internal/domain/repository"
	"database/sql"
)

type CharaRepository struct {
	db *sql.DB
}

func NewCharaRepository(db *sql.DB) repository.CharaRepository {
	return &CharaRepository{
		db: db,
	}
}

func (r *CharaRepository) GetAllCharas() ([]*model.Chara, error) {
	rows, err := r.db.Query("SELECT * FROM chara")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var charas []*model.Chara

	for rows.Next() {
		var c model.Chara
		if err := rows.Scan(&c.ID, &c.Name, &c.Probability); err != nil {
			return nil, err
		}
		charas = append(charas, &c)
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}

	return charas, nil
}
