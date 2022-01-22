package storage

import (
	"time"

	"github.com/jmoiron/sqlx"
)

type AuthStorage interface {
	Insert(token string, id int) error
	GetCustomerByToken(token string) (Customer, error)
	GetCustomerIDByToken(token string) (int, error)
}

type Auths struct {
	*sqlx.DB
}

type Auth struct {
	ID      int
	Token   string
	Expired time.Time
}

func (db *Auths) GetCustomerByToken(token string) (customer Customer, err error) {
	row := db.QueryRow("SELECT u.id, u.username, u.password, u.role, u.is_admin FROM public.customer AS u JOIN public.security AS t ON t.customer_id=u.id WHERE t.token=($1)", token)
	err = row.Scan(&customer.ID, &customer.Name, &customer.password, &customer.Role, &customer.IsAdmin)
	if err != nil { //TODO no rows
		return Customer{}, err
	}
	return
}

func (db *Auths) Insert(token string, id int) error {
	_, err := db.Exec("INSERT INTO security (token, customer_id) VALUES ($1, $2)", token, id)
	return err
}

func (db *Auths) GetCustomerIDByToken(token string) (id int, err error) {
	row := db.QueryRow("SELECT security.customer_id FROM security WHERE token=($1)", token)
	err = row.Scan(&id)
	if err != nil { //TODO no rows
		return 0, err
	}
	return
}
