package repository

import (
	"ca-mission/internal/config"
	"ca-mission/internal/domain/model"
	"database/sql"
	"testing"

	_ "github.com/go-sql-driver/mysql"
)

const (
	name       = "test_name"
	token      = "test_token"
	updateName = "update_test_name"
)

func TestCreate(t *testing.T) {
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

	u := model.User{
		Name:  name,
		Token: token,
	}

	user := NewUser(db)

	err = user.Create(&u)
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

	u := model.User{
		Name:  name,
		Token: token,
	}

	_, err = db.Exec("INSERT INTO user(name, token) VALUES(?,?)", u.Name, u.Token)
	if err != nil {
		t.Fatal(err)
	}

	u = model.User{
		Token: token,
	}

	user := NewUser(db)

	err = user.Get(&u)
	if err != nil {
		t.Fatal(err)
	}

	if name != u.Name {
		t.Errorf("name must be %s but %s", u.Name, name)
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

	u := model.User{
		Name:  name,
		Token: token,
	}

	_, err = db.Exec("INSERT INTO user(name, token) VALUES(?,?)", u.Name, u.Token)
	if err != nil {
		t.Fatal(err)
	}

	u = model.User{
		Name:  updateName,
		Token: token,
	}

	user := NewUser(db)

	err = user.Update(&u)
	if err != nil {
		t.Fatal(err)
	}

	var actual string
	if err := db.QueryRow("select name from user where token = ?", token).Scan(&actual); err != nil {
		t.Fatal(err)
	}
	if actual != updateName {
		t.Errorf("name must be %s but %s", updateName, actual)
	}
}
