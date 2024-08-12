package models

type User struct {
	UserId       uint64 `db:"user_id"`
	UserName     string `json:"username" db:"username"`
	Name         string `json:"name" db:"name"`
	Email        string `json:"email" db:"email"`
	Phone        string `json:"phone" db:"phone"`
	PasswordHash string `json:"password,omitempty" db:"password_hash"`
}
