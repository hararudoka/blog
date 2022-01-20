package storage

import (
	"time"

	"github.com/jmoiron/sqlx"
)

type (
	AuthStorage interface {
		Insert(token string, userid int) error
		UserByToken(token string) (User, error)
		UserID(token string) (int, error)
	}

	Auths struct {
		*sqlx.DB
	}

	Auth struct {
		ID      int
		Token   string
		Expired time.Time
	}
)

func (db *Auths) UserByToken(token string) (User, error) {
	var user User

	row := db.QueryRow("SELECT u.id, u.username, u.password, u.role, u.is_admin FROM public.customer AS u JOIN public.security AS t ON t.customer_id=u.id WHERE t.token=($1)", token)

	err := row.Scan(&user.ID, &user.Name, &user.password, &user.Role, &user.IsAdmin)
	if err != nil {
		return User{}, err
	}
	return user, nil
}

func (db *Auths) Insert(token string, userID int) error {
	_, err := db.Exec("INSERT INTO security (token, customer_id) VALUES ($1, $2)", token, userID)
	return err
}

func (db *Auths) UserID(token string) (int, error) {
	var id int
	row := db.QueryRow("SELECT security.customer_id FROM security WHERE token=($1)", token)
	err := row.Scan(&id)
	if err != nil {
		return 0, err
	}
	return id, nil
}
