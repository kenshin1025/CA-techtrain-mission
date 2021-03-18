package gacha

import (
	"bytes"
	"database/sql"
	"encoding/json"
	"fmt"
	"io"
	"math/rand"
	"net/http"
	"strings"
	"time"
)

type Rarity struct {
	Probability float64
	Charas      []Character
}

type Character struct {
	ID   string `json:"characterID"`
	Name string `json:"name"`
}

type Response struct {
	Results []Character `json:"results"`
}

type DrawTimes struct {
	Times int `json:"times"`
}

func Draw(w http.ResponseWriter, r *http.Request, db *sql.DB) {
	switch r.Method {
	case "POST":
		// リクエストbody(json)を受け取る
		body := r.Body
		defer body.Close()

		// byte配列に変換するためにcopy
		buf := new(bytes.Buffer)
		io.Copy(buf, body)

		//繰り返す回数
		var times DrawTimes
		err := json.Unmarshal(buf.Bytes(), &times)

		userToken := r.Header.Get("x-token")

		rand.Seed(time.Now().UnixNano())
		var res Response
		for i := 0; i < times.Times; i++ {
			character, err := oneDraw(userToken, rand.Float64(), db)
			if err != nil {
				fmt.Fprint(w, err)
				return
			}
			res.Results = append(res.Results, *character)
		}

		resJSON, err := json.Marshal(res)
		if err != nil {
			fmt.Fprint(w, err)
			return
		}

		// レスポンス
		w.Header().Add("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		w.Write(resJSON)
	}
}

func oneDraw(userToken string, randN float64, db *sql.DB) (*Character, error) {
	//これoneDrawのたびに呼び出すの効率悪くない？
	config, err := getConfig(db)
	if err != nil {
		return nil, err
	}
	boundary := 0.0
	for _, raritiy := range config {
		for _, chara := range raritiy.Charas {
			boundary += raritiy.Probability / float64(len(raritiy.Charas))
			if randN <= boundary {

				return &chara, nil
			}
		}
	}
	return nil, err
}

func getConfig(db *sql.DB) ([]Rarity, error) {
	var config []Rarity
	rows, err := db.Query("SELECT rarity.probability, GROUP_CONCAT(chara.id), GROUP_CONCAT(chara.name) FROM rarity JOIN chara ON rarity.id = chara.rarity_id GROUP BY rarity.id ORDER BY rarity.probability")
	if err != nil {
		return nil, err
	}
	for rows.Next() {
		var r Rarity
		var stringCharaIDs string
		var stringCharaNames string
		if err := rows.Scan(&r.Probability, &stringCharaIDs, &stringCharaNames); err != nil {
			return nil, err
		}

		sliceCharaIDs := strings.Split(stringCharaIDs, ",")
		sliceCharaNames := strings.Split(stringCharaNames, ",")
		var charas []Character
		for i, CharaID := range sliceCharaIDs {
			charas = append(charas, Character{ID: CharaID, Name: sliceCharaNames[i]})
		}
		r.Charas = charas
		config = append(config, r)
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	rows.Close()

	return config, nil
}
