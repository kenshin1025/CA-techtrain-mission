package user

import (
	"bytes"
	"database/sql"
	"encoding/json"
	"io"
	"log"

	"net/http"
	"time"

	"ca-mission/api/middleware/myError"
	"ca-mission/api/middleware/token"

	"github.com/go-playground/validator"
	_ "github.com/go-sql-driver/mysql"
)

type User struct {
	ID        int       `db:"id"`
	Name      string    `db:"name"`
	Token     string    `db:"token"`
	CreatedAt time.Time `db:"created_at"`
	UpdatedAt time.Time `db:"updated_at"`
}

// user情報のjson用の構造体
type UserName struct {
	Name string `validate:"required",json:"name"`
}

// userのtokenのjson用の構造体
type UserToken struct {
	Token string `json:"token"`
}

// ユーザー作成
func Create(w http.ResponseWriter, r *http.Request, db *sql.DB) {
	switch r.Method {
	case "POST":
		// リクエストbody(json)を受け取る
		body := r.Body
		defer body.Close()

		// byte配列に変換するためにcopy
		buf := new(bytes.Buffer)
		io.Copy(buf, body)

		// byte配列にしたbody内のjsonをgoで扱えるようにobjectに変換
		var name UserName
		err := json.Unmarshal(buf.Bytes(), &name)
		if err != nil {
			log.Fatal(err)
		}

		validate := validator.New()
		err = validate.Struct(name)
		if err != nil {
			log.Fatal(err)
			// レスポンス
			err := myError.ErrorResponse(w, http.StatusBadRequest)
			if err != nil {
				log.Fatal(err)
			}
			return
		}

		// tokenとしてuuid作成
		token, err := token.CreateToken()
		if err != nil {
			log.Fatal(err)
		}

		// トランザクション開始
		tx, err := db.Begin()
		if err != nil {
			log.Fatal(err)
		}

		// DBに追加
		//レコードを取得する必要のない、クエリはExecメソッドを使う。
		_, execErr := tx.Exec("INSERT INTO user(name, token) VALUES(?,?)", name.Name, token)
		//エラーが起きたらロールバック
		if execErr != nil {
			_ = tx.Rollback()
			log.Fatal(execErr)
		}
		// エラーが起きなければコミット
		err = tx.Commit()
		if err != nil {
			log.Fatal(err)
		}

		// レスポンス用のjson生成
		t := UserToken{token}
		res, err := json.Marshal(t)
		if err != nil {
			log.Fatal(err)
		}

		// レスポンス
		w.WriteHeader(http.StatusOK)
		w.Write(res)

	default:
		err := myError.ErrorResponse(w, http.StatusMethodNotAllowed)
		if err != nil {
			log.Fatal(err)
		}
	}
}

func Get(w http.ResponseWriter, r *http.Request, db *sql.DB) {
	switch r.Method {
	case "GET":
		// リクエストheaderを受け取る
		header := r.Header

		// トランザクション開始
		tx, err := db.Begin()
		if err != nil {
			log.Fatal(err)
		}

		// tokenを元にユーザーのnameを取得
		var name string
		execErr := tx.QueryRow("SELECT name FROM user WHERE token = ?", header.Get("x-token")).Scan(&name)
		if execErr != nil {
			_ = tx.Rollback()
			log.Fatal(execErr)
		}
		// エラーが起きなければコミット
		err = tx.Commit()
		if err != nil {
			log.Fatal(err)
		}

		// レスポンス用のjson生成
		u := UserName{name}
		res, err := json.Marshal(u)
		if err != nil {
			log.Fatal(err)
		}

		// レスポンス
		w.WriteHeader(http.StatusOK)
		w.Write(res)
	default:
		err := myError.ErrorResponse(w, http.StatusMethodNotAllowed)
		if err != nil {
			log.Fatal(err)
		}
	}
}

func Update(w http.ResponseWriter, r *http.Request, db *sql.DB) {
	switch r.Method {
	case "PUT":
		// リクエストを受け取る
		header := r.Header
		body := r.Body
		defer body.Close()

		// byteに変換するためにcopy
		buf := new(bytes.Buffer)
		io.Copy(buf, body)

		// byte配列にしたbody内のjsonをgoで扱えるようにobjectに変換
		var name UserName
		json.Unmarshal(buf.Bytes(), &name)

		// トランザクション開始
		tx, err := db.Begin()
		if err != nil {
			log.Fatal(err)
		}

		// tokenを元にユーザーのnameを更新
		_, execErr := tx.Exec("UPDATE user SET name = ? WHERE token = ?", name.Name, header.Get("x-token"))
		if execErr != nil {
			_ = tx.Rollback()
			log.Fatal(execErr)
		}
		// エラーが起きなければコミット
		err = tx.Commit()
		if err != nil {
			log.Fatal(err)
		}

		// レスポンス
		w.WriteHeader(http.StatusOK)
	default:
		err := myError.ErrorResponse(w, http.StatusMethodNotAllowed)
		if err != nil {
			log.Fatal(err)
		}
	}
}