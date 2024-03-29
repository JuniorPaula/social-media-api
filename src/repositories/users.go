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

func (repository users) FindById(ID uint64) (models.Users, error) {
	rows, err := repository.db.Query(
		"SELECT id, name, nickname, email, createdAt FROM users WHERE id = ?",
		ID,
	)

	if err != nil {
		return models.Users{}, err
	}

	defer rows.Close()

	var user models.Users

	if rows.Next() {
		if err = rows.Scan(
			&user.ID,
			&user.Name,
			&user.Nickname,
			&user.Email,
			&user.CreatedAt,
		); err != nil {
			return models.Users{}, err
		}
	}

	return user, nil
}

func (repository users) Update(ID uint64, user models.Users) error {
	statement, err := repository.db.Prepare(
		"UPDATE users set name = ?, nickname = ?, email = ? WHERE id = ?",
	)
	if err != nil {
		return err
	}

	defer statement.Close()

	if _, err := statement.Exec(user.Name, user.Nickname, user.Email, ID); err != nil {
		return err
	}

	return nil
}

func (repository users) Delete(ID uint64) error {
	statement, err := repository.db.Prepare(
		"DELETE from users WHERE id = ?",
	)
	if err != nil {
		return err
	}

	defer statement.Close()

	if _, err = statement.Exec(ID); err != nil {
		return err
	}

	return nil
}

func (repository users) FindByEmail(email string) (models.Users, error) {
	row, err := repository.db.Query(
		"SELECT id, password FROM users WHERE email = ?", email,
	)
	if err != nil {
		return models.Users{}, err
	}
	defer row.Close()

	var user models.Users

	if row.Next() {
		if err = row.Scan(&user.ID, &user.Password); err != nil {
			return models.Users{}, err
		}
	}

	return user, nil
}

func (repository users) Follower(userID, followerID uint64) error {
	statement, err := repository.db.Prepare(
		"INSERT ignore INTO followers (user_id, follower_id) values (?,?)",
	)
	if err != nil {
		return err
	}
	defer statement.Close()

	if _, err := statement.Exec(userID, followerID); err != nil {
		return err
	}

	return nil
}

func (repository users) Unfollower(userID, followerID uint64) error {
	statement, err := repository.db.Prepare(
		"DELETE from followers WHERE user_id = ? AND follower_id = ?",
	)
	if err != nil {
		return err
	}

	defer statement.Close()

	if _, err := statement.Exec(userID, followerID); err != nil {
		return err
	}

	return nil
}

func (repository users) FindUsersFollowers(userID uint64) ([]models.Users, error) {
	rows, err := repository.db.Query(
		`SELECT u.id, u.name, u.nickname, u.email FROM users u INNER JOIN
		 followers s ON u.id = s.follower_id WHERE s.user_id = ?`,
		userID,
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
		); err != nil {
			return nil, err
		}

		users = append(users, user)
	}

	return users, nil
}

func (repository users) FindFollowing(userID uint64) ([]models.Users, error) {
	rows, err := repository.db.Query(`
		SELECT u.id, u.name, u.nickname, u.email FROM users u INNER JOIN 
		followers s ON u.id = s.user_id WHERE s.follower_id = ?
	`, userID)
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
		); err != nil {
			return nil, err
		}

		users = append(users, user)
	}

	return users, nil
}

func (repository users) FindPasswordByUSerID(userID uint64) (string, error) {
	row, err := repository.db.Query("SELECT password FROM users WHERE id = ?", userID)
	if err != nil {
		return "", err
	}

	defer row.Close()

	var user models.Users

	if row.Next() {
		if err = row.Scan(&user.Password); err != nil {
			return "", err
		}
	}

	return user.Password, nil
}

func (repository users) UppdatePassword(userID uint64, password string) error {
	statement, err := repository.db.Prepare("UPDATE users SET password = ? where id = ?")
	if err != nil {
		return err
	}

	defer statement.Close()

	if _, err = statement.Exec(password, userID); err != nil {
		return err
	}

	return nil
}
