package main

import (
	"bytes"
	"database/sql"
	"fmt"
	"encoding/json"
	_ "github.com/go-sql-driver/mysql"
	"io"
	"log"
	"net/http"

	"github.com/dgrijalva/jwt-go"
)

// dbを全ての関数で使いたいので宣言
var db *sql.DB

// user情報のjson用の構造体
type userJSON struct {
	Name string `json:"name"`
}

// userのtokenのjson用の構造体
type tokenJSON struct {
	Token string `json:"token"`
}

// Token 作成関数
func createToken(user userJSON) (string, error) {
	var err error

	// 鍵となる文字列(多分なんでもいい)
	secret := "secret"

	// Token を作成
	// jwt -> JSON Web Token - JSON をセキュアにやり取りするための仕様
	// jwtの構造 -> {Base64 encoded Header}.{Base64 encoded Payload}.{Signature}
	// HS254 -> 証明生成用(https://ja.wikipedia.org/wiki/JSON_Web_Token)
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"name": user.Name,
		"iss":   "__init__", // JWT の発行者が入る(文字列(__init__)は任意)
	})

	tokenString, err := token.SignedString([]byte(secret))

	if err != nil {
		log.Fatal(err)
	}

	return tokenString, nil
}

// ユーザー作成
func create(w http.ResponseWriter, r *http.Request) {
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
		json.Unmarshal(buf.Bytes(), &user)

		// jwtでtoken作成
		token, err := createToken(user)

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
		fmt.Fprint(w, "Method not allowed.\n")
	}
}

func get(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case "GET":
		// リクエストheaderを受け取る
		header := r.Header

		// tokenを元にユーザーのnameを取得
		var name string
		if err := db.QueryRow("SELECT name FROM user WHERE token = ?", header.Get("x-token")).Scan(&name); err != nil {
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
		fmt.Fprint(w, "Method not allowed.\n")
	}
}

func update(w http.ResponseWriter, r *http.Request) {
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
		fmt.Fprint(w, "Method not allowed.\n")
	}
}

func main() {
	var err error
	fmt.Printf("Starting server at 'localhost:8080'\n")
	db, err = sql.Open("mysql", "root:@/ca_mission")
	if err != nil {
		panic(err.Error()) // Just for example purpose. You should use proper error handling instead of panic
	}
	defer db.Close()

	http.HandleFunc("/user/create", create)
	http.HandleFunc("/user/get", get)
	http.HandleFunc("/user/update", update)
	http.ListenAndServe(":8080", nil)
}
