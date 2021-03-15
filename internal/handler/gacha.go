package handler

import (
	"ca-mission/internal/apierr"
	"ca-mission/internal/usecase"
	"encoding/json"
	"log"
	"net/http"

	"github.com/go-playground/validator"
)

type ReqGachaJSON struct {
	Times int `json:"times" validate:"required"`
}

type ResGachaJSON struct {
	Results []Result `json:"results"`
}

type Result struct {
	CharacterID int    `json:"characterID"`
	Name        string `json:"name"`
}

func Gacha(gachaUsecase *usecase.Gacha) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		//jsonからgoの構造体にデコードする
		var gacha ReqGachaJSON
		//http通信などのストリームデータをデコードする際はNewDecoderが使える
		if err := json.NewDecoder(r.Body).Decode(&gacha); err != nil {
			log.Fatal(err)
			return
		}

		//バリデーション
		validate := validator.New()
		if err := validate.Struct(&gacha); err != nil {
			log.Fatal(err)
			return
		}

		charas, err := gachaUsecase.Draw(gacha.Times, r.Header.Get("x-token"))
		if err != nil {
			log.Fatal(err)
			writeError(w, http.StatusInternalServerError, apierr.ErrInternalServerError)
			return
		}

		var results []Result
		for i := 0; i < len(charas); i++ {
			result := Result{
				CharacterID: charas[i].ID,
				Name:        charas[i].Name,
			}
			results = append(results, result)
		}

		w.Header().Set("Content-Type", "application/json;charset=utf-8")
		w.WriteHeader(http.StatusOK)

		if err := json.NewEncoder(w).Encode(&ResGachaJSON{
			Results: results,
		}); err != nil {
			log.Fatal(err)
		}
	}
}
