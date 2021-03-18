package handler

import (
	"ca-mission/internal/apierr"
	"ca-mission/internal/domain/usecase"
	"ca-mission/internal/model"
	"encoding/json"
	"log"
	"net/http"

	"github.com/go-playground/validator"
)

type ReqCreateUserJSON struct {
	Name string `json:"name" validate:"required"`
}

type ResCreateUserJSON struct {
	Token string `json:"token"`
}

type ResGetUserJSON struct {
	Name string `json:"name"`
}

type ReqUpdateUserJSON struct {
	Name string `json:"name" validate:"required"`
}

func CreateUser(userUsecase usecase.UserInterface) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

		//jsonからgoの構造体にデコードする
		var user ReqCreateUserJSON
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
			writeError(w, http.StatusInternalServerError, apierr.ErrInternalServerError)
			return
		}

		w.Header().Set("Content-Type", "application/json;charset=utf-8")
		w.WriteHeader(http.StatusCreated)

		if err := json.NewEncoder(w).Encode(&ResCreateUserJSON{
			Token: m.Token,
		}); err != nil {
			log.Fatal(err)
		}
	}
}

func GetUser(userUsecase usecase.UserInterface) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

		m := &model.User{
			Token: r.Header.Get("x-token"),
		}

		if err := userUsecase.Get(m); err != nil {
			log.Fatal(err)
			writeError(w, http.StatusInternalServerError, apierr.ErrInternalServerError)
			return
		}

		w.Header().Set("Content-Type", "application/json;charset=utf-8")
		w.WriteHeader(http.StatusOK)

		if err := json.NewEncoder(w).Encode(&ResGetUserJSON{
			Name: m.Name,
		}); err != nil {
			log.Fatal(err)
		}
	}
}

func UpdateUser(userUsecase usecase.UserInterface) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

		//jsonからgoの構造体にデコードする
		var user ReqUpdateUserJSON
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
			Name:  user.Name,
			Token: r.Header.Get("x-token"),
		}

		if err := userUsecase.Update(m); err != nil {
			log.Fatal(err)
			writeError(w, http.StatusInternalServerError, apierr.ErrInternalServerError)
			return
		}

		w.Header().Set("Content-Type", "application/json;charset=utf-8")
		w.WriteHeader(http.StatusNoContent)
	}
}
