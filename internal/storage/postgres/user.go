package postgres

import (
	"log"

	"github.com/jmoiron/sqlx"
	"github.com/johnmcoro/quetzalapi/models"
)

type UserStorage interface {
	GetUsers() ([]models.UserDBModel, error)
}

type userStorage struct {
	DB *sqlx.DB
}

func NewUserStorage(db *sqlx.DB) *userStorage {
	return &userStorage{
		DB: db,
	}
}

func (us userStorage) GetUsers() ([]models.UserDBModel, error) {
	users := []models.UserDBModel{}
	err := us.DB.Select(&users, "SELECT username, email FROM users")
	if err != nil {
		log.Println("postgres error", err)
		return nil, err
	}
	return users, nil
}
