package config

import (
	"ca-mission/internal/domain/model"
	"database/sql"
)

type GachaConfig struct {
	SumAllProbability int
	Charas            []*model.Chara
}

func GenerateGachaConfig(db *sql.DB) (*GachaConfig, error) {
	var g GachaConfig
	rows, err := db.Query("SELECT * FROM chara")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		var c model.Chara
		if err := rows.Scan(&c.ID, &c.Name, &c.Probability); err != nil {
			return nil, err
		}
		g.SumAllProbability += c.Probability
		g.Charas = append(g.Charas, &c)
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}

	return &g, nil
}
