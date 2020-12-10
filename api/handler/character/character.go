package character

func Gacha(n float64) string {
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
