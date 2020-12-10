package main

import (
	"database/sql"
	"fmt"
	"log"

	"net/http"
	"os"

	"ca-mission/api/handler/gacha"
	"ca-mission/api/handler/user"

	_ "github.com/go-sql-driver/mysql"

	"github.com/joho/godotenv"
)

func test(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "ok")
}

func main() {
	fmt.Printf("Starting server at 'http://localhost:8080'\n")

	err := godotenv.Load("./.env")
	if err != nil {
		log.Fatal(err)
	}

	//.envファイルからdataSourceNameを作成してDBに接続する
	dsn := fmt.Sprintf("%s:%s@tcp(db:3306)/%s", os.Getenv("DB_ROOT_USER"), os.Getenv("DB_ROOT_PASS"), os.Getenv("DB_DATABASE"))
	db, err := sql.Open("mysql", dsn)
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	http.HandleFunc("/", test)

	http.HandleFunc("/user/create", func(w http.ResponseWriter, r *http.Request) {
		user.Create(w, r, db)
	})
	http.HandleFunc("/user/get", func(w http.ResponseWriter, r *http.Request) {
		user.Get(w, r, db)
	})
	http.HandleFunc("/user/update", func(w http.ResponseWriter, r *http.Request) {
		user.Update(w, r, db)
	})
	http.HandleFunc("/gacha/draw", func(w http.ResponseWriter, r *http.Request) {
		gacha.Draw(w, r, db)
	})
	http.ListenAndServe(":8080", nil)
}
