package main

import (
	"fmt"
	"log"

	"net/http"

	"ca-mission/api/handler/gacha"
	"ca-mission/api/handler/user"
	"ca-mission/db"

	_ "github.com/go-sql-driver/mysql"
)

func test(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "ok")
}

func main() {
	fmt.Printf("Starting server at 'http://localhost:8080'\n")

	//DBに接続する
	db, err := db.ConnectSQL()
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
