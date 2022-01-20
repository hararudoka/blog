package web

import (
	"github.com/hararudoka/blog/internal/storage"
	"html/template"
	"io"
	"net/http"

	"github.com/labstack/echo/v4"
)

type Template struct {
	Templates *template.Template
}

func (t *Template) Render(w io.Writer, name string, data interface{}, c echo.Context) error {
	return t.Templates.ExecuteTemplate(w, name, data)
}

type Temp struct {
	PageTitle string
	Customer  storage.Customer
	IsAuth    bool
	Error     error
}

func (temp *Temp) DefaultTemp(c echo.Context, s *storage.DB) error {
	cok, err := c.Cookie("token")
	if err == http.ErrNoCookie {
		temp.Customer = storage.Customer{ID: 0}

		temp.IsAuth = false
	} else if err == nil {
		user, err := s.Auths.GetCustomerByToken(cok.Value)
		if err != nil {
			return err
		}

		temp.Customer = user

		temp.IsAuth = true
	}
	return err
}
