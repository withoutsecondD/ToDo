package models

type User struct {
	ID        int64  `db:"id" json:"id"`
	Age       int32  `db:"age" json:"age"`
	FirstName string `db:"first_name" json:"firstName"`
	LastName  string `db:"last_name" json:"lastName"`
	City      string `db:"city" json:"city"`
	Email     string `db:"email" json:"email"`
}
