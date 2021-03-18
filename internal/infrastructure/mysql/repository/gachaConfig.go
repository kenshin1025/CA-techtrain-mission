package repository

import (
	"ca-mission/internal/domain/repository"
	"ca-mission/internal/model"
	"database/sql"
)

type GachaConfig struct {
	db *sql.DB
}

func NewGachaConfig(db *sql.DB) repository.GachaConfigRepository {
	return &GachaConfig{
		db: db,
	}
}

func (g *GachaConfig) GetAllCharas() ([]*model.Chara, error) {
	rows, err := g.db.Query("SELECT * FROM chara")
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
