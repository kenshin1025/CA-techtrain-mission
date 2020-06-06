package main

import (
	"bytes"
	"database/sql"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"

	"./token"
	"./error"

	_ "github.com/go-sql-driver/mysql"

	"github.com/joho/godotenv"
)

// user情報のjson用の構造体
type userJSON struct {
	Name string `json:"name"`
}

// userのtokenのjson用の構造体
type tokenJSON struct {
	Token string `json:"token"`
}

// ユーザー作成
func create(w http.ResponseWriter, r *http.Request, db *sql.DB) {
	switch r.Method {
	case "POST":
		// リクエストbody(json)を受け取る
		body := r.Body
		defer body.Close()

		// byte配列に変換するためにcopy
		buf := new(bytes.Buffer)
		io.Copy(buf, body)

		// byte配列にしたbody内のjsonをgoで扱えるようにobjectに変換
		var user userJSON
		err := json.Unmarshal(buf.Bytes(), &user)
		if err != nil {
			log.Fatal(err)
		}
		// tokenとしてuuid作成
		token := token.CreateToken()

		// DBに追加
		//レコードを取得する必要のない、クエリはExecメソッドを使う。
		_, err = db.Exec("INSERT INTO user(name, token) VALUES(?,?)", user.Name, token)
		if err != nil {
			log.Fatal(err)
		}

		// レスポンス用のjson生成
		t := tokenJSON{token}
		res, err := json.Marshal(t)
		if err != nil {
			log.Fatal(err)
		}

		// レスポンス
		w.WriteHeader(http.StatusOK)
		w.Write(res)

	default:
		w = error.Respons405(w)
	}
}

func get(w http.ResponseWriter, r *http.Request, db *sql.DB) {
	switch r.Method {
	case "GET":
		// リクエストheaderを受け取る
		header := r.Header

		// tokenを元にユーザーのnameを取得
		var name string
		err := db.QueryRow("SELECT name FROM user WHERE token = ?", header.Get("x-token")).Scan(&name)
		if err != nil {
			log.Fatal(err)
		}

		// レスポンス用のjson生成
		u := userJSON{name}
		res, err := json.Marshal(u)
		if err != nil {
			log.Fatal(err)
		}

		// レスポンス
		w.WriteHeader(http.StatusOK)
		w.Write(res)
	default:
		w = error.Respons405(w)
	}
}

func update(w http.ResponseWriter, r *http.Request, db *sql.DB) {
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
		var user userJSON
		json.Unmarshal(buf.Bytes(), &user)

		// tokenを元にユーザーのnameを更新
		_, err := db.Exec("UPDATE user SET name = ? WHERE token = ?", user.Name, header.Get("x-token"))
		if err != nil {
			log.Fatal(err)
		}

		// レスポンス
		w.WriteHeader(http.StatusOK)
	default:
		w = error.Respons405(w)
	}
}

func main() {
	fmt.Printf("Starting server at 'localhost:8080'\n")

	err := godotenv.Load("./.env")
	if err != nil {
		log.Fatal(err)
	}

	//.envファイルからdataSourceNameを作成してDBに接続する
	dsn := fmt.Sprintf("%s:@%s/%s", os.Getenv("DB_USER"), os.Getenv("DB_PASS"), os.Getenv("DB_NAME"))
	db, err := sql.Open("mysql", dsn)
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	http.HandleFunc("/user/create", func(w http.ResponseWriter, r *http.Request) {
		create(w, r, db)
	})
	http.HandleFunc("/user/get", func(w http.ResponseWriter, r *http.Request) {
		get(w, r, db)
	})
	http.HandleFunc("/user/update", func(w http.ResponseWriter, r *http.Request) {
		update(w, r, db)
	})
	http.ListenAndServe(":8080", nil)
}
