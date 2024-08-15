package models

type Tag struct {
	ID     int64  `db:"id" json:"id"`
	UserID int64  `db:"user_id" json:"user_id"`
	Title  string `db:"title" json:"title"`
	Color  string `db:"color" json:"color"`
}

type TagCreateDto struct {
	Title string `json:"title" validate:"required"`
	Color string `json:"color" validate:"hexcolor"`
}
