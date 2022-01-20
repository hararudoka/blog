package storage

import (
	"log"
	"time"

	"github.com/jmoiron/sqlx"
)

type (
	UserStorage interface {
		UserByID(id int) (user User, err error)
		Exists(username string) (user User, err error)
		Insert(user User) error
	}

	Users struct {
		*sqlx.DB
	}

	User struct {
		ID        int
		Name      string
		password  string
		Role      string
		IsAdmin   bool
		CreatedAt time.Time
	}
)

func (db *Users) Insert(user User) error {
	_, err := db.Exec("INSERT INTO customer (username, password, role, is_admin, created_at) VALUES ($1, $2, $3, $4, $5)", user.Name, user.password, user.Role, user.IsAdmin, time.Now())
	return err
}

func (db *Users) UserByID(id int) (user User, err error) {
	row := db.QueryRow("SELECT * FROM customer WHERE id=($1)", id)
	err = row.Scan(&user.ID, &user.Name, &user.password, &user.Role, &user.IsAdmin, &user.CreatedAt)
	log.Println(err)
	return user, err
}

func (db *Users) Exists(username string) (user User, err error) {
	row := db.QueryRow("SELECT * FROM customer WHERE customer.username=($1)", username)
	err = row.Scan(&user.ID, &user.Name, &user.password, &user.Role, &user.IsAdmin, &user.CreatedAt)
	return user, err
}

func (u *User) Password() string {
	return u.password
}

func (u *User) SetPassword(password string) {
	u.password = password
}
