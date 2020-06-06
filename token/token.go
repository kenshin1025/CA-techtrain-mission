package token

import (
	"github.com/google/uuid"
	"log"
)

func CreateToken() string {
	// tokenとしてuuid作成
	uuid, err := uuid.NewRandom()
	if err != nil {
		log.Fatal(err)
	}
	return uuid.String()
}
