package storage

import (
	"database/sql"
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
	ID         int `db:"id"`
	CustomerID int `db:"customer_id"`
	Title      string `db:"title"`
	Content    string `db:"content"`
	CreatedAt  time.Time `db:"created_at"`
	HumanTime  string
}

func (db *Posts) GetByID(id int) (Post, error) {
	var post Post

	err := db.QueryRow("SELECT * FROM post WHERE id=($1)", id).Scan(&post.ID, &post.CustomerID, &post.Title, &post.Content, &post.CreatedAt)
	if err != nil { //TODO no rows
		return Post{}, err
	}
	db.DB.Begin()


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
	if err == sql.ErrNoRows {

	}
	return n, err
}
