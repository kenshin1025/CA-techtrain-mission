package token

import (
	"github.com/google/uuid"
)

func CreateToken() (string, error) {
	// tokenとしてuuid作成
	uuid, err := uuid.NewRandom()
	return uuid.String(), err
}
