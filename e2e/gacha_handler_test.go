package e2e

import (
	"bytes"
	"ca-mission/internal/config"
	"ca-mission/internal/handler"
	"ca-mission/internal/repository"
	"ca-mission/internal/usecase"
	"database/sql"
	"encoding/json"
	"log"
	"net/http"
	"net/http/httptest"
	"testing"
)

func Test_E2E_Gacha(t *testing.T) {
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

	gachaConfig, err := config.GenerateGachaConfig(db)
	if err != nil {
		log.Fatal(err)
	}

	if _, err := db.Exec("insert into user(name, token) values (?, ?)", name, token); err != nil {
		t.Fatal(err)
	}

	const times = 1

	gachaUsecase := usecase.NewGacha(repository.NewGacha(), db, gachaConfig)
	var body bytes.Buffer
	if err := json.NewEncoder(&body).Encode(&handler.ReqGachaJSON{
		Times: times,
	}); err != nil {
		t.Fatal(err)
	}

	// requestをシュミレートする
	req := httptest.NewRequest(http.MethodPost, "/", &body)
	req.Header.Set("x-token", token)
	rec := httptest.NewRecorder()
	http.HandlerFunc(handler.Gacha(gachaUsecase)).ServeHTTP(rec, req)

	// responseのStatus Codeをチェックする
	if rec.Code != http.StatusOK {
		t.Errorf("status code must be 200 but: %d", rec.Code)
		t.Fatalf("body: %s", rec.Body.String())
	}

	var result handler.ResGachaJSON
	if err := json.NewDecoder(rec.Body).Decode(&result); err != nil {
		t.Fatal(err)
	}

	if len(result.Results) != times {
		t.Fatalf("result length must be %d but: %d", times, len(result.Results))
	}

	// charaがinsertされているかチェックする
	var actual int
	if err := db.QueryRow("select chara_id from user_chara_possession limit 1").Scan(&actual); err != nil {
		t.Fatal(err)
	}
	if actual != result.Results[0].CharacterID {
		t.Errorf("character_id must be %d but %d", result.Results[0].CharacterID, actual)
	}
}
