package handler

import (
	"blog/storage"
	"blog/web"
	"database/sql"
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
)

type UserStorage struct {
	handler
}

func (s *UserStorage) REGISTER(h handler, g *echo.Group) {
	s.handler = h

	g.GET("/:id", s.GetUser)
}

func (s *UserStorage) GetUser(c echo.Context) error {
	var temp struct {
		web.Temp
		storage.User
	}

	err := temp.DefaultTemp(c, s.db)
	if err != nil {
		return err
	}

	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return err
	}

	user, err := s.db.Users.UserByID(id)
	if err == sql.ErrNoRows {
		return c.Render(http.StatusOK, "404", temp)
	} else if err != nil {
		return err
	}

	temp.User = user

	return c.Render(http.StatusOK, "user", temp)
}
