package database

import (
	"ca-mission/internal/config"
	"database/sql"
	"log"

	_ "github.com/go-sql-driver/mysql"
)

//https://qiita.com/fiemon/items/eb38c8d681ed1ae05925
//この記事のほぼパクリ、sql.openだけ自分の作ったDBにつながるように変えた

type MyDB struct {
	db *sql.DB
}

var Db MyDB

// Connection
func (m *MyDB) Connection() error {
	var err error
	m.db, err = sql.Open("mysql", config.Config().GenerateDSN())
	if err != nil {
		return err
	}
	return nil
}

// Close
func (m *MyDB) Close() {
	if m.db != nil {
		m.db.Close()
	}
}

//　Transaction
func (m *MyDB) Transaction(txFunc func(*sql.Tx) error) error {
	tx, err := m.db.Begin()
	if err != nil {
		return err
	}

	defer func() {
		if p := recover(); p != nil {
			log.Println("recover")
			tx.Rollback()
			panic(p)
		} else if err != nil {
			log.Println("rollback")
			tx.Rollback()
		} else {
			log.Println("commit")
			err = tx.Commit()
		}
	}()
	err = txFunc(tx)
	return err
}
