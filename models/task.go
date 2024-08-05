package models

type Task struct {
	ID          int64  `db:"id" json:"id"`
	ListID      int64  `db:"list_id" json:"list_id"`
	Title       string `db:"title" json:"title"`
	Description string `db:"description" json:"description"`
	Status      string `db:"status" json:"status"`
	Deadline    string `db:"deadline" json:"deadline"`
	Tags        []Tag  `json:"tags"`
}
