package handler

import (
	"ca-mission/internal/apierr"
	"ca-mission/internal/domain/model"
	"ca-mission/internal/usecase"
	"encoding/json"
	"log"
	"net/http"
)

type ResUserCharacterListJSON struct {
	Characters []UserCharacter `json:"characters"`
}

type UserCharacter struct {
	UserCharacterID int    `json:"userCharacterID"`
	CharacterID     int    `json:"characterID"`
	Name            string `json:"name"`
}

func GetUserCharacterList(characterUsecase usecase.CharacterLister) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		u := &model.User{
			Token: r.Header.Get("x-token"),
		}

		userCharas, err := characterUsecase.GetUserCharacterList(u)
		if err != nil {
			log.Fatal(err)
			writeError(w, http.StatusInternalServerError, apierr.ErrInternalServerError)
			return
		}

		var characters []UserCharacter
		for _, userChara := range userCharas {
			uc := UserCharacter{
				UserCharacterID: userChara.ID,
				CharacterID:     userChara.Chara.ID,
				Name:            userChara.Chara.Name,
			}
			characters = append(characters, uc)
		}

		w.Header().Set("Content-Type", "application/json;charset=utf-8")
		w.WriteHeader(http.StatusOK)

		if err := json.NewEncoder(w).Encode(&ResUserCharacterListJSON{
			Characters: characters,
		}); err != nil {
			log.Fatal(err)
		}
	}
}
