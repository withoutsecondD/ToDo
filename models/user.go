package models

// User struct represents a user in application. Bcrypt is used to hash passwords
type User struct {
	ID        int64  `db:"id" json:"id"`
	Age       int32  `db:"age" json:"age" validate:"min=0,max=100"`
	FirstName string `db:"first_name" json:"firstName" validate:"required,max=100"`
	LastName  string `db:"last_name" json:"lastName" validate:"required,max=100"`
	City      string `db:"city" json:"city" validate:"max=40"`
	Email     string `db:"email" json:"email" validate:"email"`
	Password  string `db:"password" json:"password" validate:"max=50"`
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
