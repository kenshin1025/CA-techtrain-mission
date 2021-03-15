package main

import (
	"database/sql"
	"fmt"
	"log"

	"net/http"

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

	gachaConfig, err := config.GenerateGachaConfig(db)
	if err != nil {
		log.Fatal(err)
	}

	r := mux.NewRouter()

	r.HandleFunc("/", test)

	userUsecase := usecase.NewUser(repository.NewUser(), db)
	r.HandleFunc("/user/create", handler.CreateUser(userUsecase)).Methods("POST")
	r.HandleFunc("/user/get", handler.GetUser(userUsecase)).Methods("GET")
	r.HandleFunc("/user/update", handler.UpdateUser(userUsecase)).Methods("PUT")

	gachaUsecase := usecase.NewGacha(repository.NewGacha(), db, gachaConfig)
	r.HandleFunc("/gacha/draw", handler.Gacha(gachaUsecase)).Methods("POST")
	http.ListenAndServe(":8080", r)
}
