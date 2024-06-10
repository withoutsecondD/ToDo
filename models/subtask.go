package models

type SubTask struct {
	ID     int64
	TaskID int64
	Title  string
	IsDone bool
}
