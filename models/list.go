package models

type List struct {
	ID     int64  `db:"id" json:"id"`
	UserID int64  `db:"user_id" json:"userId"`
	Title  string `db:"title" json:"title"`
}
