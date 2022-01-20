package storage

import (
	"time"

	"github.com/jmoiron/sqlx"
)

type CustomerStorage interface {
	GetByID(id int) (customer Customer, err error)
	GetByName(name string) (customer Customer, err error)
	Insert(customer Customer) error
}

type Customers struct {
	*sqlx.DB
}

type Customer struct {
	ID        int
	Name      string
	password  string
	Role      string
	IsAdmin   bool
	CreatedAt time.Time
}

func (db *Customers) Insert(customer Customer) error {
	_, err := db.Exec("INSERT INTO customer (username, password, role, is_admin, created_at) VALUES ($1, $2, $3, $4, $5)", customer.Name, customer.password, customer.Role, customer.IsAdmin, time.Now())
	return err
}

func (db *Customers) GetByID(id int) (customer Customer, err error) {
	row := db.QueryRow("SELECT * FROM customer WHERE id=($1)", id)
	err = row.Scan(&customer.ID, &customer.Name, &customer.password, &customer.Role, &customer.IsAdmin, &customer.CreatedAt)
	return
}

func (db *Customers) GetByName(username string) (customer Customer, err error) {
	row := db.QueryRow("SELECT * FROM customer WHERE customer.username=($1)", username)
	err = row.Scan(&customer.ID, &customer.Name, &customer.password, &customer.Role, &customer.IsAdmin, &customer.CreatedAt)
	return
}

func (u *Customer) Password() string {
	return u.password
}

func (u *Customer) SetPassword(password string) {
	u.password = password
}
