package model

import "time"

type User struct {
	ID        int       `db:"id"`
	Name      string    `db:"name"`
	Token     string    `db:"token"`
	CreatedAt time.Time `db:"created_at"`
	UpdatedAt time.Time `db:"updated_at"`
}
