package models

type UserDBModel struct {
	Username string `db:"username"`
	Email    string `db:"email"`
}

type UserJsonModel struct {
}
