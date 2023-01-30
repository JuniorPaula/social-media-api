package repositories

import (
	"api/src/models"
	"database/sql"
)

type Posts struct {
	db *sql.DB
}

func NewPostsRepository(db *sql.DB) *Posts {
	return &Posts{db}
}

func (repository Posts) Create(post models.Post) (uint64, error) {
	statement, err := repository.db.Prepare(
		"INSERT INTO posts (title, content, author_id) values (?,?,?)",
	)
	if err != nil {
		return 0, err
	}
	defer statement.Close()

	result, err := statement.Exec(post.Title, post.Content, post.AuthorID)
	if err != nil {
		return 0, err
	}

	insertedID, err := result.LastInsertId()
	if err != nil {
		return 0, err
	}

	return uint64(insertedID), nil
}

func (repository Posts) FindById(postID uint64) (models.Post, error) {
	row, err := repository.db.Query(
		"SELECT p.*, u.nickname FROM posts p INNER JOIN users u ON u.id = p.author_id WHERE p.id = ?",
		postID,
	)
	if err != nil {
		return models.Post{}, err
	}
	defer row.Close()

	var post models.Post

	if row.Next() {
		if err = row.Scan(
			&post.ID,
			&post.Title,
			&post.Content,
			&post.AuthorID,
			&post.Likes,
			&post.CreatedAt,
			&post.AuthorNickname,
		); err != nil {
			return models.Post{}, err
		}
	}

	return post, nil
}

func (repository Posts) FindAll(userId uint64) ([]models.Post, error) {
	rows, err := repository.db.Query(`
		SELECT distinct p.*, u.nickname FROM posts p LEFT JOIN users u ON u.id = p.author_id
		LEFT JOIN followers f ON p.author_id = f.user_id WHERE u.id = ? or f.follower_id = ?
		ORDER BY 1 DESC
	`, userId, userId)

	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var posts []models.Post

	for rows.Next() {
		var post models.Post

		if err = rows.Scan(
			&post.ID,
			&post.Title,
			&post.Content,
			&post.AuthorID,
			&post.Likes,
			&post.CreatedAt,
			&post.AuthorNickname,
		); err != nil {
			return nil, err
		}

		posts = append(posts, post)

	}
	return posts, nil
}

func (repository Posts) Update(postID uint64, post models.Post) error {
	statement, err := repository.db.Prepare("UPDATE posts set title = ?, content = ? WHERE id = ?")
	if err != nil {
		return err
	}
	defer statement.Close()

	if _, err = statement.Exec(post.Title, post.Content, postID); err != nil {
		return err
	}

	return nil
}

func (repository Posts) Delete(postID uint64) error {
	statement, err := repository.db.Prepare("DELETE FROM posts WHERE id = ?")
	if err != nil {
		return err
	}

	defer statement.Close()

	if _, err = statement.Exec(postID); err != nil {
		return err
	}

	return nil
}

func (repository Posts) FindPostsFromUser(userId uint64) ([]models.Post, error) {
	rows, err := repository.db.Query(`
		SELECT p.*, u.nickname FROM posts p JOIN users u ON u.id = p.author_id WHERE p.author_id = ? 
	`, userId)
	if err != nil {
		return nil, err
	}

	defer rows.Close()

	var posts []models.Post

	for rows.Next() {
		var post models.Post

		if err = rows.Scan(
			&post.ID,
			&post.Title,
			&post.Content,
			&post.AuthorID,
			&post.Likes,
			&post.CreatedAt,
			&post.AuthorNickname,
		); err != nil {
			return nil, err
		}
		posts = append(posts, post)
	}
	return posts, nil
}

func (repository Posts) Like(postId uint64) error {
	statement, err := repository.db.Prepare("UPDATE posts SET likes = likes + 1 WHERE id = ?")
	if err != nil {
		return err
	}
	defer statement.Close()

	if _, err = statement.Exec(postId); err != nil {
		return err
	}

	return nil
}
