package error

import (
	"encoding/json"
	"log"
	"net/http"
)

type messageJSON struct {
	Massage string `json:"massage"`
}

func Respons405(w http.ResponseWriter) http.ResponseWriter{
	// レスポンス用のjson生成
	m := messageJSON{"そのメソッドは定義されていません"}
	res, err := json.Marshal(m)
	if err != nil {
		log.Fatal(err)
	}

	w.WriteHeader(http.StatusMethodNotAllowed)
	w.Write(res)

	return w
}