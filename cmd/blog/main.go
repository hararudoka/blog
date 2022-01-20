package main

import (
	"github.com/hararudoka/blog/internal/handler"
	"github.com/hararudoka/blog/internal/storage"
	"github.com/hararudoka/blog/web"
	"html/template"

	"github.com/labstack/echo/v4/middleware"

	"github.com/labstack/echo/v4"
)

func main() {
	t := &web.Template{
		Templates: template.Must(template.ParseGlob("web/*.html")),
	}

	db, err := storage.Open()
	if err != nil {
		panic(err)
	}

	h := handler.New(handler.Handler{DB: db})

	e := echo.New()
	e.Renderer = t
	e.HTTPErrorHandler = h.CustomHTTPErrorHandler

	e.Use(middleware.Logger())
	//e.Use(middleware.Recover())

	h.Register(e.Group(""), &handler.PostStorage{})

	h.Register(e.Group("/users"), &handler.CustomerStorage{})
	h.Register(e.Group("/comments"), &handler.CommentStorage{})
	h.Register(e.Group(""), &handler.AuthService{})

	e.Logger.Fatal(e.Start(":80"))
}
