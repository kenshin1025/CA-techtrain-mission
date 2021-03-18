package e2e

import (
	"bytes"
	"ca-mission/internal/config"
	"ca-mission/internal/domain/usecase"
	"ca-mission/internal/handler"
	"ca-mission/internal/infrastructure/mysql/repository"
	"database/sql"
	"encoding/json"
	"log"
	"net/http"
	"net/http/httptest"
	"testing"

	_ "github.com/go-sql-driver/mysql"
)

// testで作成するuserのdata
const (
	name  = "test_name"
	token = "aaaaaaaa-aaaa-aaaa-aaaa-aaaaaaaaaaaaaaaa"
)

func Test_E2E_CreateUser(t *testing.T) {
	//DBに接続する
	db, err := sql.Open("mysql", config.Config().GenerateDSN())
	if err != nil {
		log.Fatal(err)
	}
	defer func() {
		// DBのcleanを行う
		db.Exec("set foreign_key_checks = 0")
		db.Exec("truncate table user")
		db.Exec("set foreign_key_checks = 1")
		db.Close()
	}()

	userUsecase := usecase.NewUser(repository.NewUser(db))

	var body bytes.Buffer
	if err := json.NewEncoder(&body).Encode(&handler.ReqCreateUserJSON{
		Name: name,
	}); err != nil {
		t.Fatal(err)
	}

	// requestをシュミレートする
	req := httptest.NewRequest(http.MethodPost, "/", &body)
	rec := httptest.NewRecorder()
	http.HandlerFunc(handler.CreateUser(userUsecase)).ServeHTTP(rec, req)

	// responseのStatus Codeをチェックする
	if rec.Code != http.StatusCreated {
		t.Errorf("status code must be 201 but: %d", rec.Code)
		t.Fatalf("body: %s", rec.Body.String())
	}

	var result handler.ResCreateUserJSON
	if err := json.NewDecoder(rec.Body).Decode(&result); err != nil {
		t.Fatal(err)
	}

	// responseで返ってきたIDでuserが作られているかどうかをチェックする
	var actual string
	if err := db.QueryRow("select name from user where token = ?", result.Token).Scan(&actual); err != nil {
		t.Fatal(err)
	}
	if actual != name {
		t.Errorf("name must be %s but %s", name, actual)
	}
}

func Test_E2E_GetUser(t *testing.T) {
	//DBに接続する
	db, err := sql.Open("mysql", config.Config().GenerateDSN())
	if err != nil {
		log.Fatal(err)
	}
	defer func() {
		// DBのcleanを行う
		db.Exec("set foreign_key_checks = 0")
		db.Exec("truncate table user")
		db.Exec("set foreign_key_checks = 1")
		db.Close()
	}()

	if _, err := db.Exec("insert into user(name, token) values (?, ?)", name, token); err != nil {
		t.Fatal(err)
	}

	userUsecase := usecase.NewUser(repository.NewUser(db))

	// requestをシュミレートする
	req := httptest.NewRequest(http.MethodGet, "/", nil)
	req.Header.Set("x-token", token)
	rec := httptest.NewRecorder()
	http.HandlerFunc(handler.GetUser(userUsecase)).ServeHTTP(rec, req)

	// responseのStatus Codeをチェックする
	if rec.Code != http.StatusOK {
		t.Errorf("status code must be 200 but: %d", rec.Code)
		t.Fatalf("body: %s", rec.Body.String())
	}

	var result handler.ResGetUserJSON
	if err := json.NewDecoder(rec.Body).Decode(&result); err != nil {
		t.Fatal(err)
	}

	// responseで返ってきたnameがinsertしたtokenと紐づいたnameか確認
	if result.Name != name {
		t.Errorf("name must be %s but %s", name, result.Name)
	}
}

func Test_E2E_UpdateUser(t *testing.T) {
	//DBに接続する
	db, err := sql.Open("mysql", config.Config().GenerateDSN())
	if err != nil {
		log.Fatal(err)
	}
	defer func() {
		// DBのcleanを行う
		db.Exec("set foreign_key_checks = 0")
		db.Exec("truncate table user")
		db.Exec("set foreign_key_checks = 1")
		db.Close()
	}()

	update_name := "update_name"

	if _, err := db.Exec("insert into user(name, token) values (?, ?)", name, token); err != nil {
		t.Fatal(err)
	}

	userUsecase := usecase.NewUser(repository.NewUser(db))

	var body bytes.Buffer
	if err := json.NewEncoder(&body).Encode(&handler.ReqUpdateUserJSON{
		Name: update_name,
	}); err != nil {
		t.Fatal(err)
	}

	// requestをシュミレートする
	req := httptest.NewRequest(http.MethodPut, "/", &body)
	req.Header.Set("x-token", token)
	rec := httptest.NewRecorder()
	http.HandlerFunc(handler.UpdateUser(userUsecase)).ServeHTTP(rec, req)

	// responseのStatus Codeをチェックする
	if rec.Code != http.StatusNoContent {
		t.Errorf("status code must be 204 but: %d", rec.Code)
		t.Fatalf("body: %s", rec.Body.String())
	}

	// nameがupdateされているかチェックする
	var actual string
	if err := db.QueryRow("select name from user where token = ?", token).Scan(&actual); err != nil {
		t.Fatal(err)
	}
	if actual != update_name {
		t.Errorf("update_name must be %s but %s", update_name, actual)
	}
}
