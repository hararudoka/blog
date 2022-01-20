package handler

import (
	"fmt"
	"github.com/hararudoka/blog/internal/storage"
	"github.com/hararudoka/blog/web"
	"net/http"

	"github.com/labstack/echo/v4/middleware"

	"github.com/labstack/echo/v4"
)

type Handler struct {
	DB *storage.DB
}

type handler struct {
	db *storage.DB
}

func NewHandler(h Handler) *handler {
	return &handler{
		db: h.DB,
	}
}

type Service interface {
	REGISTER(h handler, g *echo.Group)
}

func (h handler) Register(group *echo.Group, service Service) {
	service.REGISTER(h, group)
}

func (h handler) CustomHTTPErrorHandler(err error, c echo.Context) {
	code := http.StatusInternalServerError
	if he, ok := err.(*echo.HTTPError); ok {
		code = he.Code
	}

	c.Logger().Error(err)

	var temp struct {
		web.Temp
		Error error
	}

	e := temp.DefaultTemp(c, h.db)
	if e != nil {
		c.Logger().Error(err)
	}
	temp.Error = err

	if code == http.StatusNotFound {
		err = c.Render(http.StatusOK, "404", temp)
	} else {
		err = c.Render(http.StatusOK, "error", temp)
	}

	if err != nil {
		c.Logger().Error(err)
	}
}

func Middleware(db storage.DB) func(echo.HandlerFunc) echo.HandlerFunc {
	return middleware.KeyAuthWithConfig(middleware.KeyAuthConfig{
		KeyLookup: "cookie:token",

		Validator: func(key string, c echo.Context) (bool, error) {
			_, err := db.Auths.UserByToken(key)
			if err != nil {
				return false, err
			}
			return true, nil
		},

		ErrorHandler: func(err error, c echo.Context) error {
			if err != fmt.Errorf("missing key in cookies: %w", http.ErrNoCookie) {
				return c.Redirect(http.StatusFound, "/../login")
			}

			var temp web.Temp

			e := temp.DefaultTemp(c, &db)
			if e != nil {
				return e
			}

			temp.Error = err
			return c.Render(http.StatusOK, "error", temp)
		},

		//Skipper: func(c echo.Context) bool {
		//	return true
		//},
	})
}
