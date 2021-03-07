package handler

import (
	"ca-mission/internal/model"
	"ca-mission/internal/usecase"
	"encoding/json"
	"log"
	"net/http"

	"github.com/go-playground/validator"
)

type ReqUserCreateJSON struct {
	Name string `json:"name" validate:"required"`
}

type ResUserCreateJSON struct {
	Token string `json:"token"`
}

func UserCreate(userUsecase *usecase.User) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

		//jsonからgoの構造体にデコードする
		var user ReqUserCreateJSON
		//http通信などのストリームデータをデコードする際はNewDecoderが使える
		if err := json.NewDecoder(r.Body).Decode(&user); err != nil {
			log.Fatal(err)
			return
		}

		//バリデーション
		validate := validator.New()
		if err := validate.Struct(&user); err != nil {
			log.Fatal(err)
			return
		}

		m := &model.User{
			Name: user.Name,
		}
		if err := userUsecase.Create(m); err != nil {
			log.Fatal(err)
			return
		}

		w.Header().Set("Content-Type", "application/json;charset=utf-8")
		w.WriteHeader(http.StatusCreated)

		if err := json.NewEncoder(w).Encode(&ResUserCreateJSON{
			Token: m.Token,
		}); err != nil {
			log.Fatal(err)
		}
	}
}
