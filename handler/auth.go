package handler

import (
	"blog/storage"
	"blog/web"
	"database/sql"
	"fmt"
	"math/rand"
	"net/http"
	"time"

	"github.com/labstack/echo/v4"
)

type AuthService struct {
	handler
}

func (s *AuthService) REGISTER(h handler, g *echo.Group) {
	s.handler = h

	g.GET("/login", s.Login)
	g.POST("/login", s.RenderLogin)

	g.GET("/register", s.Reg)
	g.POST("/register", s.RenderReg)
}

func (s *AuthService) Login(c echo.Context) error {
	var temp web.Temp
	err := temp.DefaultTemp(c, s.db)
	if err != nil {
		return err
	}
	return c.Render(http.StatusOK, "login", temp)
}

func (s *AuthService) Reg(c echo.Context) error {
	var temp web.Temp
	err := temp.DefaultTemp(c, s.db)
	if err != nil {
		return err
	}
	return c.Render(http.StatusOK, "register", temp)
}

func (s *AuthService) RenderLogin(c echo.Context) error {
	username := c.FormValue("username")
	password := c.FormValue("password")

	user, err := s.db.Users.Exists(username)
	if err != nil {
		fmt.Println(err)
		return c.Redirect(http.StatusFound, "/404")
	}

	if password != user.Password() {
		fmt.Println("неверный пароль")
		return c.Redirect(http.StatusFound, "/404")
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
	c.SetCookie(cookie)

	return c.Redirect(http.StatusFound, "/feed")
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

func (s *AuthService) RenderReg(c echo.Context) error {
	username := c.FormValue("username")
	password := c.FormValue("password")

	_, err := s.db.Users.Exists(username)
	if err == sql.ErrNoRows {
		var u storage.User
		u.Name = username
		u.SetPassword(password)
		u.Role = "user"
		err = s.db.Users.Insert(u)
		if err != nil {
			return err
		}
	}
	if err != nil {
		return err
	}
	return c.Redirect(http.StatusFound, "/feed")
}
