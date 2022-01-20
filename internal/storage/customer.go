package storage

import (
	"time"

	"github.com/jmoiron/sqlx"
)

type (
	CustomerStorage interface {
		UserByID(id int) (user Customer, err error)
		Exists(username string) (user Customer, err error)
		Insert(user Customer) error
	}

	Customers struct {
		*sqlx.DB
	}

	Customer struct {
		ID        int
		Name      string
		password  string
		Role      string
		IsAdmin   bool
		CreatedAt time.Time
	}
)

func (db *Customers) Insert(user Customer) error {
	_, err := db.Exec("INSERT INTO customer (username, password, role, is_admin, created_at) VALUES ($1, $2, $3, $4, $5)", user.Name, user.password, user.Role, user.IsAdmin, time.Now())
	return err
}

func (db *Customers) UserByID(id int) (user Customer, err error) {
	row := db.QueryRow("SELECT * FROM customer WHERE id=($1)", id)
	err = row.Scan(&user.ID, &user.Name, &user.password, &user.Role, &user.IsAdmin, &user.CreatedAt)
	return user, err
}

func (db *Customers) Exists(username string) (user Customer, err error) {
	row := db.QueryRow("SELECT * FROM customer WHERE customer.username=($1)", username)
	err = row.Scan(&user.ID, &user.Name, &user.password, &user.Role, &user.IsAdmin, &user.CreatedAt)
	return user, err
}

func (u *Customer) Password() string {
	return u.password
}

func (u *Customer) SetPassword(password string) {
	u.password = password
}
