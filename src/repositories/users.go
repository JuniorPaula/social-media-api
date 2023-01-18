package repositories

import (
	"api/src/models"
	"database/sql"
)

type users struct {
	db *sql.DB
}

func CreateUsersRepository(db *sql.DB) *users {
	return &users{db}
}

func (u users) Create(user models.Users) (uint64, error) {
	return 0, nil
}
