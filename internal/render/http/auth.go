package http

import (
	"database/sql"
	"math/rand"
	"net/http"
	"time"

	"github.com/hararudoka/blog/internal/storage"
	"github.com/labstack/echo/v4"
)

type AuthService struct {
	handler
}

func (s *AuthService) Register(h handler, g *echo.Group) {
	s.handler = h

	g.GET("/login", s.Login)
	g.POST("/login", s.RenderLogin)

	g.GET("/register", s.Reg)
	g.POST("/register", s.RenderReg)
}

func (s *AuthService) Login(ctx echo.Context) error {
	var temp Temp
	err := temp.DefaultTemp(ctx, s.db)
	if err != nil {
		return err
	}
	return ctx.Render(http.StatusOK, "login", temp)
}

func (s *AuthService) Reg(ctx echo.Context) error {
	var temp Temp
	err := temp.DefaultTemp(ctx, s.db)
	if err != nil {
		return err
	}
	return ctx.Render(http.StatusOK, "register", temp)
}

func (s *AuthService) RenderLogin(ctx echo.Context) error {
	username := ctx.FormValue("username")
	password := ctx.FormValue("password")

	user, err := s.db.Customers.GetByName(username)
	if err != nil {
		return ctx.Redirect(http.StatusFound, "/404")
	}

	if password != user.Password() {
		return ctx.Redirect(http.StatusFound, "/404")
	}

	token := generateToken()

	err = s.db.Auths.Insert(token, user.ID)
	if err != nil {
		return err
	}

	cookie := new(http.Cookie)
	cookie.Name = "token"
	cookie.Value = token
	cookie.Expires = time.Now().Add(72 * time.Hour)
	ctx.SetCookie(cookie)

	return ctx.Redirect(http.StatusFound, "/feed")
}

func generateToken() string {
	rand.Seed(time.Now().UnixNano())

	var letterRunes = []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789")
	b := make([]rune, 20)
	for i := range b {
		b[i] = letterRunes[rand.Intn(len(letterRunes))]
	}
	return string(b)
}

func (s *AuthService) RenderReg(ctx echo.Context) error {
	name := ctx.FormValue("username")
	password := ctx.FormValue("password")

	_, err := s.db.Customers.GetByName(name)
	if err == sql.ErrNoRows {
		var c storage.Customer
		c.Name = name
		c.SetPassword(password)
		c.Role = "user"
		err = s.db.Customers.Insert(c)
		if err != nil {
			return err
		}
	}
	if err != nil {
		return err
	}
	return ctx.Redirect(http.StatusFound, "/feed")
}
