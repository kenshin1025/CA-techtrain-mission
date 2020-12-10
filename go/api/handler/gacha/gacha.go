package gacha

import (
	"database/sql"
	"fmt"
	"net/http"
)

func Draw(w http.ResponseWriter, r *http.Request, db *sql.DB) {
	character := oneDraw(0.01)
	fmt.Fprint(w, character)
}

func oneDraw(n float64) string {
	chara := ""
	if n < 0.01 {
		chara = "大当たり"
	} else if n < 0.11 {
		chara = "当たり"
	} else {
		chara = "ハズレ"
	}
	return chara
}
