package repositories

import (
	"api/src/models"
	"database/sql"
	"fmt"
)

type users struct {
	db *sql.DB
}

func UsersRepository(db *sql.DB) *users {
	return &users{db}
}

func (repository users) Create(user models.Users) (uint64, error) {
	statement, err := repository.db.Prepare(
		"INSERT INTO users (name, nickname, email, password) values(?, ?, ?, ?)",
	)
	if err != nil {
		return 0, err
	}

	defer statement.Close()

	result, err := statement.Exec(user.Name, user.Nickname, user.Email, user.Password)
	if err != nil {
		return 0, err
	}

	lastIDInserted, err := result.LastInsertId()
	if err != nil {
		return 0, err
	}

	return uint64(lastIDInserted), nil
}

func (repository users) Find(nameOrNickname string) ([]models.Users, error) {
	nameOrNickname = fmt.Sprintf("%%%s%%", nameOrNickname)

	rows, err := repository.db.Query(
		"SELECT id, name, nickname, email, createdAt FROM users WHERE name LIKE ? or nickname LIKE ?",
		nameOrNickname, nameOrNickname,
	)

	if err != nil {
		return nil, err
	}

	defer rows.Close()

	var users []models.Users

	for rows.Next() {
		var user models.Users

		if err = rows.Scan(
			&user.ID,
			&user.Name,
			&user.Nickname,
			&user.Email,
			&user.CreatedAt,
		); err != nil {
			return nil, err
		}

		users = append(users, user)
	}

	return users, nil
}
