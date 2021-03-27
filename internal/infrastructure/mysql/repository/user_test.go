package repository

import (
	"ca-mission/internal/config"
	"ca-mission/internal/domain/model"
	"database/sql"
	"strconv"
	"testing"

	_ "github.com/go-sql-driver/mysql"
)

var userTests = []*model.User{
	{Name: "test_name", Token: "test_token"},
}

func TestCreate(t *testing.T) {
	//DBに接続する
	db, err := conectTestDB()
	if err != nil {
		t.Fatal(err)
	}
	defer func() {
		// DBのcleanを行う
		db.Exec("set foreign_key_checks = 0")
		db.Exec("truncate table user")
		db.Exec("set foreign_key_checks = 1")
		db.Close()
	}()

	user := NewUser(db)

	for i, u := range userTests {
		t.Run(strconv.Itoa(i)+u.Name, func(t *testing.T) {
			err := user.Create(u)
			if err != nil {
				t.Fatal(err)
			}
			var actual string
			if err := db.QueryRow("select name from user where token = ?", u.Token).Scan(&actual); err != nil {
				t.Fatal(err)
			}
			if actual != u.Name {
				t.Errorf("name must be %s but %s", u.Name, actual)
			}
		})
	}
}

func TestGet(t *testing.T) {
	//DBに接続する
	db, err := sql.Open("mysql", config.Config().GenerateDSN())
	if err != nil {
		t.Fatal(err)
	}
	defer func() {
		// DBのcleanを行う
		db.Exec("set foreign_key_checks = 0")
		db.Exec("truncate table user")
		db.Exec("set foreign_key_checks = 1")
		db.Close()
	}()

	user := NewUser(db)

	for i, u := range userTests {
		t.Run(strconv.Itoa(i)+u.Name, func(t *testing.T) {
			_, err = db.Exec("INSERT INTO user(name, token) VALUES(?,?)", u.Name, u.Token)
			if err != nil {
				t.Fatal(err)
			}
			insertName := u.Name
			u.Name = ""
			err := user.Get(u)
			if err != nil {
				t.Fatal(err)
			}
			if insertName != u.Name {
				t.Errorf("name must be %s but %s", u.Name, insertName)
			}
		})
	}
}

func TestUpdate(t *testing.T) {
	//DBに接続する
	db, err := sql.Open("mysql", config.Config().GenerateDSN())
	if err != nil {
		t.Fatal(err)
	}
	defer func() {
		// DBのcleanを行う
		db.Exec("set foreign_key_checks = 0")
		db.Exec("truncate table user")
		db.Exec("set foreign_key_checks = 1")
		db.Close()
	}()

	user := NewUser(db)

	for i, u := range userTests {
		t.Run(strconv.Itoa(i)+u.Name, func(t *testing.T) {
			_, err = db.Exec("INSERT INTO user(name, token) VALUES(?,?)", "not_update_name", u.Token)
			if err != nil {
				t.Fatal(err)
			}

			err := user.Update(u)
			if err != nil {
				t.Fatal(err)
			}
			var actual string
			if err := db.QueryRow("select name from user where token = ?", u.Token).Scan(&actual); err != nil {
				t.Fatal(err)
			}
			if actual != u.Name {
				t.Errorf("name must be %s but %s", u.Name, actual)
			}
		})
	}
}
