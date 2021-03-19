package main

import (
	"database/sql"
	"fmt"
	"log"

	"net/http"

	"ca-mission/internal/config"
	"ca-mission/internal/domain/usecase"
	"ca-mission/internal/handler"
	"ca-mission/internal/infrastructure/mysql/repository"

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

	gachaConfigUsecase := usecase.NewGachaConfig(repository.NewGachaConfig(db))
	gachaConfig, err := gachaConfigUsecase.GenerateGachaConfig()
	if err != nil {
		log.Fatal(err)
	}

	r := mux.NewRouter()

	r.HandleFunc("/", test)

	userUsecase := usecase.NewUser(repository.NewUser(db))
	r.HandleFunc("/user/create", handler.CreateUser(userUsecase)).Methods("POST")
	r.HandleFunc("/user/get", handler.GetUser(userUsecase)).Methods("GET")
	r.HandleFunc("/user/update", handler.UpdateUser(userUsecase)).Methods("PUT")

	gachaUsecase := usecase.NewGacha(repository.NewGacha(db), gachaConfig)
	r.HandleFunc("/gacha/draw", handler.Gacha(gachaUsecase)).Methods("POST")

	characterUsecase := usecase.NewCharacter(repository.NewCharacter(db))
	r.HandleFunc("/character/list", handler.GetUserCharacterList(characterUsecase)).Methods("GET")
	http.ListenAndServe(":8080", r)
}
