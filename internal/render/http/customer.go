package http

import (
	"database/sql"
	"net/http"
	"strconv"

	"github.com/hararudoka/blog/internal/storage"
	"github.com/labstack/echo/v4"
)

type CustomerStorage struct {
	handler
}

func (s *CustomerStorage) Register(h handler, g *echo.Group) {
	s.handler = h

	g.GET("/:id", s.GetUser)
}

func (s *CustomerStorage) GetUser(c echo.Context) error {
	var temp struct {
		Temp
		storage.Customer
	}

	err := temp.DefaultTemp(c, s.db)
	if err != nil {
		return err
	}

	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return err
	}

	customer, err := s.db.Customers.GetByID(id)
	if err == sql.ErrNoRows {
		return c.Render(http.StatusOK, "404", temp)
	} else if err != nil {
		return err
	}

	temp.Customer = customer

	return c.Render(http.StatusOK, "user", temp)
}
