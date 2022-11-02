package models

type UserModels struct {
	Email    string `db:"email"`
	Password string `db:"password"`
}
