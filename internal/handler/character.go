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

		charas, err := characterUsecase.GetUserCharacterList(u)
		if err != nil {
			log.Fatal(err)
			writeError(w, http.StatusInternalServerError, apierr.ErrInternalServerError)
			return
		}

		var userCharacters []UserCharacter
		for _, chara := range charas {
			uc := UserCharacter{
				UserCharacterID: 0,
				CharacterID:     chara.ID,
				Name:            chara.Name,
			}
			userCharacters = append(userCharacters, uc)
		}

		w.Header().Set("Content-Type", "application/json;charset=utf-8")
		w.WriteHeader(http.StatusOK)

		if err := json.NewEncoder(w).Encode(&ResUserCharacterListJSON{
			Characters: userCharacters,
		}); err != nil {
			log.Fatal(err)
		}
	}
}
