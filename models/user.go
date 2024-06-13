package models

// User struct represents a user in application. Bcrypt is used to hash passwords
type User struct {
	ID        int64  `db:"id" json:"id"`
	Age       int32  `db:"age" json:"age"`
	FirstName string `db:"first_name" json:"firstName"`
	LastName  string `db:"last_name" json:"lastName"`
	City      string `db:"city" json:"city"`
	Email     string `db:"email" json:"email"`
	Password  string `db:"password" json:"password"`
}

// UserResponse struct is a User's ready-to-send variant (i.e. without any sensitive information)
type UserResponse struct {
	ID        int64  `db:"id" json:"id"`
	Age       int32  `db:"age" json:"age"`
	FirstName string `db:"first_name" json:"firstName"`
	LastName  string `db:"last_name" json:"lastName"`
	City      string `db:"city" json:"city"`
	Email     string `db:"email" json:"email"`
}
