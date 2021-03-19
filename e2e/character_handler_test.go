package e2e

import (
	"ca-mission/internal/config"
	"ca-mission/internal/handler"
	"ca-mission/internal/infrastructure/mysql/repository"
	"ca-mission/internal/usecase"
	"database/sql"
	"encoding/json"
	"log"
	"net/http"
	"net/http/httptest"
	"testing"
)

func Test_E2E_Character(t *testing.T) {
	//DBに接続する
	db, err := sql.Open("mysql", config.Config().GenerateDSN())
	if err != nil {
		log.Fatal(err)
	}
	defer func() {
		// DBのcleanを行う
		db.Exec("set foreign_key_checks = 0")
		db.Exec("truncate table user")
		db.Exec("truncate table user_chara_possession")
		db.Exec("set foreign_key_checks = 1")
		db.Close()
	}()

	// テスト用のユーザを作成
	userInsertResult, err := db.Exec("insert into user(name, token) values (?, ?)", name, token)
	if err != nil {
		t.Fatal(err)
	}

	user_id, err := userInsertResult.LastInsertId()
	if err != nil {
		t.Fatal(err)
	}

	//ユーザの所持キャラクターを設定
	if _, err := db.Exec("insert into user_chara_possession(user_id, chara_id) values (?, 1),(?, 4),(?, 7),(?, 8)", user_id, user_id, user_id, user_id); err != nil {
		t.Fatal(err)
	}

	const haveCharaLength = 4

	characterUsecase := usecase.NewCharacter(repository.NewCharacter(db))

	// requestをシュミレートする
	req := httptest.NewRequest(http.MethodGet, "/", nil)
	req.Header.Set("x-token", token)
	rec := httptest.NewRecorder()
	http.HandlerFunc(handler.GetUserCharacterList(characterUsecase)).ServeHTTP(rec, req)

	// responseのStatus Codeをチェックする
	if rec.Code != http.StatusOK {
		t.Errorf("status code must be 200 but: %d", rec.Code)
		t.Fatalf("body: %s", rec.Body.String())
	}

	var result handler.ResUserCharacterListJSON
	if err := json.NewDecoder(rec.Body).Decode(&result); err != nil {
		t.Fatal(err)
	}

	// responseで返ってきたnameがinsertしたtokenと紐づいたnameか確認
	if len(result.Characters) != haveCharaLength {
		t.Errorf("resultListLength must be %d but %d", haveCharaLength, len(result.Characters))
	}
}
