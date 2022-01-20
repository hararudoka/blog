package storage

import (
	"time"

	"github.com/jmoiron/sqlx"
)

type (
	CommentStorage interface {
		ByPost(postID int) ([]Comment, error)
		Insert(Comment) error
	}

	Comments struct {
		*sqlx.DB
	}

	Comment struct {
		Customer
		ID         int
		CustomerID int
		PostID     int
		Content   string
		CreatedAt time.Time
	}
)

func (db *Comments) ByPost(postID int) ([]Comment, error) {
	var comments []Comment

	rows, err := db.Query("SELECT * FROM comment WHERE post_id=($1)", postID)
	if err != nil {
		return []Comment{}, err
	}

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
