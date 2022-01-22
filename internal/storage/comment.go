package storage

import (
	"time"

	"github.com/jmoiron/sqlx"
)

type CommentStorage interface {
	SliceByPostID(postID int) ([]Comment, error)
	Insert(Comment) error
	GetAll() ([]Comment, error)
}

type Comments struct {
	*sqlx.DB
}

type Comment struct {
	Customer
	ID         int `db:"id"`
	CustomerID int `db:"customer_id"`
	PostID     int `db:"post_id"`
	Content    string `db:"content"`
	CreatedAt  time.Time `db:"created_at"`
}

func (db *Comments) SliceByPostID(postID int) ([]Comment, error) {
	var comments []Comment

	rows, err := db.Query("SELECT * FROM comment WHERE post_id=($1)", postID)
	if err != nil {
		return []Comment{}, err
	}

	defer rows.Close()

	for rows.Next() {
		var temp Comment
		err = rows.Scan(&temp.ID, &temp.PostID, &temp.CustomerID, &temp.Content, &temp.CreatedAt)
		if err != nil {
			return []Comment{}, err
		}
		comments = append(comments, temp)
	}

	return comments, err
}

func (db *Comments) Insert(comment Comment) error {
	_, err := db.Exec("INSERT INTO comment (customer_id, post_id, content, created_at) VALUES ($1, $2, $3, $4)", comment.CustomerID, comment.PostID, comment.Content, time.Now())
	return err
}

func (db *Comments) GetAll() ([]Comment, error) {
	var comments []Comment

	err := db.Select(&comments, "SELECT * FROM comment JOIN customer ON customer.id=comment.post_id")
	if err != nil {
		return []Comment{}, err
	}

	return comments, err
}
