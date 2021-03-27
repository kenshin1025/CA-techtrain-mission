package repository

import (
	"ca-mission/internal/config"
	"ca-mission/internal/domain/model"
	"database/sql"
	"strconv"
	"testing"

	_ "github.com/go-sql-driver/mysql"
)

var gachaTests = []struct {
	user   *model.User
	charas []*model.Chara
}{
	{
		user: &model.User{
			ID: 1,
		},
		charas: []*model.Chara{
			{ID: 1},
		},
	},
	{
		user: &model.User{
			ID: 1,
		},
		charas: []*model.Chara{
			{ID: 1},
			{ID: 3},
			{ID: 2},
		},
	},
	{
		user: &model.User{
			ID: 1,
		},
		charas: []*model.Chara{
			{ID: 1},
			{ID: 1},
			{ID: 2},
		},
	},
}

func TestSaveDrewCharas(t *testing.T) {
	//DBに接続する
	db, err := sql.Open("mysql", config.Config().GenerateDSN())
	if err != nil {
		t.Fatal(err)
	}
	defer func() {
		// DBのcleanを行う
		db.Exec("set foreign_key_checks = 0")
		db.Exec("truncate table user_chara_possession")
		db.Exec("set foreign_key_checks = 1")
		db.Close()
	}()

	gacha := NewGacha(db)

	for i, tt := range gachaTests {
		t.Run(strconv.Itoa(i), func(t *testing.T) {
			err := gacha.SaveDrewCharas(tt.user, tt.charas)
			if err != nil {
				t.Fatal(err)
			}
			// charaがinsertされているかチェックする
			rows, err := db.Query("select chara_id from user_chara_possession where user_id = ?", tt.user.ID)
			if err != nil {
				t.Fatal(err)
			}
			rowsLength := 0
			for rows.Next() {
				rowsLength++
			}

			if rowsLength != len(tt.charas) {
				t.Errorf(" rowsLength must be %d but %d", len(tt.charas), rowsLength)
			}
			db.Exec("set foreign_key_checks = 0")
			db.Exec("truncate table user_chara_possession")
			db.Exec("set foreign_key_checks = 1")
		})
	}
}
