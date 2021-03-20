package model

type Chara struct {
	ID          int    `db:"id"`
	Name        string `db:"name"`
	Probability int    `db:"probability"`
}
