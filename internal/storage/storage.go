package storage

import (
	"fmt"
	"log"
	"os"

	"github.com/jmoiron/sqlx"
	"github.com/joho/godotenv"

	_ "github.com/lib/pq"
)

type DB struct {
	*sqlx.DB
	Customers CustomerStorage
	Posts     PostStorage
	Comments  CommentStorage
	Auths     AuthStorage
}

func Open() (*DB, error) {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	name := os.Getenv("DB_USERNAME")
	password := os.Getenv("DB_PASSWORD")
	host := os.Getenv("DB_HOSTNAME")
	dbName := os.Getenv("DB_NAME")
	mode := os.Getenv("DB_MODE")

	connStr := fmt.Sprintf("postgres://%s:%s@%s/%s?sslmode=%s",
		name, password, host, dbName, mode)

	db, err := sqlx.Open("postgres", connStr)
	if err != nil {
		return nil, err
	}
	err = db.Ping()
	if err != nil {
		return nil, err
	}
	log.Println("connected")
	return &DB{
		DB:        db,
		Customers: &Customers{DB: db},
		Posts:     &Posts{DB: db},
		Comments:  &Comments{DB: db},
		Auths:     &Auths{DB: db},
	}, nil
}
