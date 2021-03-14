package main

import (
	"database/sql"
	"fmt"
	"log"

	"net/http"

	"ca-mission/api/handler/gacha"
	"ca-mission/api/handler/user"
	"ca-mission/internal/config"
	"ca-mission/internal/handler"
	"ca-mission/internal/repository"
	"ca-mission/internal/usecase"

	_ "github.com/go-sql-driver/mysql"
	"github.com/gorilla/mux"
)

func test(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "ok")
}

func main() {
	fmt.Printf("Starting server at 'http://localhost:8080'\n")

	db, err := sql.Open("mysql", config.Config().GenerateDSN())
	if err != nil {
		log.Fatal("open failed")
	}
	defer db.Close()

	r := mux.NewRouter()

	r.HandleFunc("/", test)

	userUsecase := usecase.NewUser(repository.NewUser(), db)
	r.HandleFunc("/user/create", handler.CreateUser(userUsecase)).Methods("POST")

	r.HandleFunc("/user/get", handler.GetUser(userUsecase)).Methods("GET")
	r.HandleFunc("/user/update", func(w http.ResponseWriter, r *http.Request) {
		user.Update(w, r, db)
	})
	r.HandleFunc("/gacha/draw", func(w http.ResponseWriter, r *http.Request) {
		gacha.Draw(w, r, db)
	})
	http.ListenAndServe(":8080", r)
}
