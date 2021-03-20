package model

type UserCharaPossession struct {
	ID        int `db:"id"`
	User      User
	Chara Chara
}
