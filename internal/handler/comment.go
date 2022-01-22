package handler

import (
	"errors"
	"github.com/hararudoka/blog/internal/storage"
	"github.com/hararudoka/blog/web"
	"log"
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
)

type CommentStorage struct {
	handler
}

func (s *CommentStorage) REGISTER(h handler, g *echo.Group) {
	s.handler = h

	g.GET("", s.LastComments)

	m := s.Middleware(*s.db)

	g.GET("/addComment", s.AddComment, m)
	g.POST("/addComment", s.Comment, m)

}

func (s *CommentStorage) Comment(c echo.Context) error {
	var temp struct {
		web.Temp
		Error error
	}
	err := temp.DefaultTemp(c, s.db)
	if err != nil {
		return err
	}

	content := c.FormValue("content")
	if len([]byte(content)) > 200 {
		temp.Error = errors.New("too long com")
		return c.Render(http.StatusFound, "error", temp)
	}

	postID, err := strconv.Atoi(c.QueryParam("postID"))
	if err != nil {
		return err
	}

	err = s.db.Comments.Insert(storage.Comment{
		CustomerID: temp.Customer.ID,
		PostID:     postID,
		Content:    content,
	})

	return c.Redirect(http.StatusFound, "/posts/"+c.QueryParam("postID"))
}

func (s *CommentStorage) AddComment(c echo.Context) error {
	var temp struct {
		web.Temp
		PostID string
	}
	err := temp.DefaultTemp(c, s.db)
	if err != nil {
		return err
	}

	temp.PostID = c.QueryParam("postID")

	return c.Render(http.StatusFound, "addComment", temp)
}

func (s *CommentStorage) LastComments(c echo.Context) error {
	var temp struct {
		web.Temp
		Comments []storage.Comment
	}
	err := temp.DefaultTemp(c, s.db)
	if err != nil {
		return err
	}


	temp.Comments, err = s.db.Comments.GetAll()
	if err != nil {
		return err
	}

	log.Println(temp.Comments[0].Name)

	return c.Render(http.StatusOK, "comments", temp)
}
