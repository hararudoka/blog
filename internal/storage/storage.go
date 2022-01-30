package storage

import (
	"fmt"
	"log"

	"github.com/hararudoka/blog/internal/config"
	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
)

type DB struct {
	*sqlx.DB
	Customers CustomerStorage
	Posts     PostStorage
	Comments  CommentStorage
	Auths     AuthStorage
}

func Open(e config.Env) (*DB, error) {
	conn := fmt.Sprintf("postgres://%s:%s@%s/%s?sslmode=%s",
		e.Username, e.Password, e.Hostname, e.DBName, e.Mode)

	log.Println(conn)

	db, err := sqlx.Connect("postgres", conn)
	if err != nil {
		return nil, err
	}

	return &DB{DB: db}, nil
}

