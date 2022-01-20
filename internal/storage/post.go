package storage

import (
	"time"

	"github.com/jmoiron/sqlx"
)

type PostStorage interface {
	Insert(data Post) error
	Count() (int, error)
	GetByID(id int) (Post, error)
}

type Posts struct {
	*sqlx.DB
}

type Post struct {
	Customer
	Comments   []Comment
	ID         int
	CustomerID int
	Title      string
	Content    string
	CreatedAt  time.Time
	HumanTime  string
}

func (db *Posts) GetByID(id int) (Post, error) {
	var post Post

	row := db.QueryRow("SELECT * FROM post WHERE id=($1)", id)
	err := row.Scan(&post.ID, &post.CustomerID, &post.Title, &post.Content, &post.CreatedAt)
	if err != nil {
		return Post{}, err
	}

	post.HumanTime = post.CreatedAt.Format("January 2, 2006")

	return post, nil
}

func (db *Posts) Insert(post Post) error {
	_, err := db.Exec("INSERT INTO post (customer_id, title, content, created_at) VALUES ($1, $2, $3, $4)", post.CustomerID, post.Title, post.Content, time.Now())
	return err
}

func (db *Posts) Count() (int, error) {
	var n int
	row := db.QueryRow(
		"SELECT MAX(id) FROM post")
	err := row.Scan(&n)
	return n, err
}
