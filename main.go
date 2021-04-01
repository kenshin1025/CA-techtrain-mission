package main

import (
	"database/sql"
	"fmt"
	"log"

	"net/http"

	"ca-mission/internal/cache"
	"ca-mission/internal/config"
	"ca-mission/internal/handler"
	"ca-mission/internal/infrastructure/mysql/repository"
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

	// Repositoryの初期化
	userRepository := repository.NewUserRepository(db)
	charaRepository := repository.NewCharaRepository(db)
	userCharaPossessionRepository := repository.NewUserCharaPossessionRepository(db)

	// Cacheの初期化
	GachaConfigGenerater := cache.NewGachaConfigGenerater(charaRepository)
	gachaConfig, err := GachaConfigGenerater.GenerateGachaConfig()
	if err != nil {
		log.Fatal(err)
	}

	// Usecaseの初期化
	userUsecase := usecase.NewUserUsecase(userRepository)
	gachaUsecase := usecase.NewGachaUsecase(userRepository, userCharaPossessionRepository, gachaConfig)
	characterUsecase := usecase.NewCharacterUsecase(userRepository, userCharaPossessionRepository)

	// // Middlewareの初期化
	// auth.SetUserRepository(userRepository)

	r.HandleFunc("/user/create", handler.CreateUser(userUsecase)).Methods("POST")
	r.HandleFunc("/user/get", handler.GetUser(userUsecase)).Methods("GET")
	r.HandleFunc("/user/update", handler.UpdateUser(userUsecase)).Methods("PUT")

	r.HandleFunc("/gacha/draw", handler.Gacha(gachaUsecase)).Methods("POST")

	r.HandleFunc("/character/list", handler.GetUsersCharacterList(characterUsecase)).Methods("GET")

	http.ListenAndServe(":8080", r)
}
