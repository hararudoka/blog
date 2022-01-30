package execute

import (
	"github.com/hararudoka/blog/internal/config"
	"github.com/hararudoka/blog/internal/render/http"
	"github.com/hararudoka/blog/internal/storage"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"github.com/sirupsen/logrus"
)

func Execute() {
	log := logrus.New()

	//conf, err := config.LoadConfig()
	//if err != nil {
	//	log.Fatal("Config loading fails")
	//}

	dbConfig := config.LoadEnv()

	log.Println(dbConfig)

	db, err := storage.Open(dbConfig)
	if err != nil {
		log.Fatal("DB connection fails: ", err)
	}

	h := http.New(db)

	e := echo.New()

	e.HTTPErrorHandler = h.CustomHTTPErrorHandler

	e.Use(middleware.Logger())
	e.File("/favicon.ico", "view/favicon.ico")

	h.NewGroup(e.Group(""), &http.PostStorage{})

	h.NewGroup(e.Group("/users"), &http.CustomerStorage{})
	h.NewGroup(e.Group("/comments"), &http.CommentStorage{})
	h.NewGroup(e.Group(""), &http.AuthService{})

	e.Logger.Fatal(e.Start(":80"))
}
