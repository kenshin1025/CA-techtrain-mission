package myError

import (
	"encoding/json"
	"net/http"
)

type messageJSON struct {
	Massage string `json:"massage"`
}

func ErrorResponse(w http.ResponseWriter, sc int) (error){
	// レスポンス用のjson生成
	m := messageJSON{http.StatusText(sc)}
	res, err := json.Marshal(m)

	w.WriteHeader(sc)
	w.Write(res)

	return  err
}