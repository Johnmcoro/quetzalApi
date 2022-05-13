package postgres

import (
	"github.com/jmoiron/sqlx"
)

func New() (*sqlx.DB, error) {
	db, err := sqlx.Connect("postgres", "user=johncoronel dbname=quetzal sslmode=disable")
	if err != nil {
		return nil, err
	}
	return db, nil
}
