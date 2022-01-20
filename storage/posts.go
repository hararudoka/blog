package storage

import (
	"time"

	"github.com/jmoiron/sqlx"
)

type (
	PostStorage interface {
		Insert(data Post) error
		PostCounting() (int, error)
		ByID(id int) (Post, error)
	}

	Posts struct {
		*sqlx.DB
	}

	Post struct {
		User
		Comments  []Comment
		ID        int
		UserID    int
		Title     string
		Content   string
		CreatedAt time.Time
		HumanTime string
	}
)

func (db *Posts) ByID(id int) (Post, error) {
	var post Post

	row := db.QueryRow("SELECT * FROM post WHERE id=($1)", id)
	err := row.Scan(&post.ID, &post.UserID, &post.Title, &post.Content, &post.CreatedAt)
	if err != nil {
		return Post{}, err
	}

	post.HumanTime = post.CreatedAt.Format("January 2, 2006")

	return post, nil
}

func (db *Posts) Insert(post Post) error {
	_, err := db.Exec("INSERT INTO post (customer_id, title, content, created_at) VALUES ($1, $2, $3, $4)", post.UserID, post.Title, post.Content, time.Now())
	return err
}

func (db *Posts) PostCounting() (int, error) {
	var n int
	row := db.QueryRow(
		"SELECT MAX(id) FROM post")
	err := row.Scan(&n)
	return n, err
}
