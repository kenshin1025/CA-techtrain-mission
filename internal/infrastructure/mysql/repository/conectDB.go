package repository

import (
	"ca-mission/internal/config"
	"database/sql"
)

func conectTestDB() (*sql.DB, error) {
	//DBに接続する
	db, err := sql.Open("mysql", config.Config().GenerateDSN())
	if err != nil {
		return nil, err
	}
	return db, nil
}
