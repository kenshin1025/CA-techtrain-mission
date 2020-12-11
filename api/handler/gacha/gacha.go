package gacha

import (
	"database/sql"
	"fmt"
	"net/http"
	"strings"
)

type Rarity struct {
	probability float64
	charas      []Character
}

type Character struct {
	id   string
	name string
}

func Draw(w http.ResponseWriter, r *http.Request, db *sql.DB) {
	character, err := oneDraw(0.01, db)
	if err != nil {
		fmt.Fprint(w, "エラーやんけ！")
		return
	}
	fmt.Println(character.name)
	fmt.Fprint(w, character.name)
}

func oneDraw(n float64, db *sql.DB) (*Character, error) {
	config, err := getConfig(db)
	if err != nil {
		return nil, err
	}
	boundary := 0.0
	for _, raritiy := range config {
		for _, chara := range raritiy.charas {
			boundary += raritiy.probability / float64(len(raritiy.charas))
			if n <= boundary {
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
		if err := rows.Scan(&r.probability, &stringCharaIDs, &stringCharaNames); err != nil {
			return nil, err
		}

		sliceCharaIDs := strings.Split(stringCharaIDs, ",")
		sliceCharaNames := strings.Split(stringCharaNames, ",")
		var charas []Character
		for i, CharaID := range sliceCharaIDs {
			charas = append(charas, Character{id: CharaID, name: sliceCharaNames[i]})
		}
		r.charas = charas
		config = append(config, r)
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	rows.Close()

	return config, nil
}
