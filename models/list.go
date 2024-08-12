package models

type List struct {
	ID     int64  `db:"id" json:"id"`
	UserID int64  `db:"user_id" json:"userId"`
	Title  string `db:"title" json:"title"`
}

type ListCreateDto struct {
	Title string `json:"title" validate:"min=1,max=30"`
}
