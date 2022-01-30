package http

import (
	"database/sql"
	"net/http"
	"strconv"
	"time"

	"github.com/hararudoka/blog/internal/storage"
	"github.com/labstack/echo/v4"
)

type PostStorage struct {
	handler
}

func (s *PostStorage) Register(h handler, g *echo.Group) {
	s.handler = h

	g.GET("/", func(ctx echo.Context) error {
		return ctx.Redirect(http.StatusFound, "/feed")
	})

	g.GET("contacts", s.Contacts)
	g.GET("feed", s.Feed)
	g.GET("posts/:id", s.Post)

	m := s.Middleware(*s.db)

	g.GET("addPost", s.WriteForm, m)
	g.POST("addPost", s.WriteFromForm, m)
}

func (s *PostStorage) Feed(ctx echo.Context) error {
	// posts logic
	var posts []storage.Post
	n, err := s.db.Posts.Count()
	if err != nil {
		return err
	}
	for i := 1; i <= n; i++ {
		post, err := s.post(i)
		if err == sql.ErrNoRows {
		} else if err != nil {
			return err
		}
		posts = append(posts, post)
	}
	for i, j := 0, len(posts)-1; i < j; i, j = i+1, j-1 {
		posts[i], posts[j] = posts[j], posts[i]
	}

	var temp struct {
		Temp
		Posts []storage.Post
	}
	err = temp.DefaultTemp(ctx, s.db)
	if err != nil {
		return err
	}

	temp.Posts = posts
	temp.PageTitle = "haraldka's blog"

	return ctx.Render(http.StatusOK, "feed", temp)
}

func (s *PostStorage) WriteForm(ctx echo.Context) error {
	var temp struct {
		Temp
	}
	err := temp.DefaultTemp(ctx, s.db)
	if err != nil {
		return err
	}
	return ctx.Render(http.StatusOK, "addPost", temp)
}

func (s *PostStorage) WriteFromForm(ctx echo.Context) error {
	title := ctx.FormValue("title")
	content := ctx.FormValue("content")

	cok, err := ctx.Cookie("token")

	customer, err := s.db.Auths.GetCustomerByToken(cok.Value)
	if err != nil {
		return err
	}
	err = s.db.Posts.Insert(storage.Post{CustomerID: customer.ID, Title: title, Content: content, CreatedAt: time.Now()})
	if err != nil {
		return err
	}

	return ctx.Redirect(http.StatusFound, "feed")
}

func (s *PostStorage) Post(c echo.Context) error {
	var temp struct {
		Temp
		storage.Post
	}
	err := temp.DefaultTemp(c, s.db)
	if err != nil {
		return err
	}

	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return err
	}

	post, err := s.post(id)
	if err != nil {
		return err
	}

	temp.Post = post

	return c.Render(http.StatusOK, "post", temp)
}

func (s *PostStorage) post(id int) (storage.Post, error) {
	post, err := s.db.Posts.GetByID(id)
	if err == sql.ErrNoRows {
		return storage.Post{Title: "удалён"}, nil
	} else if err != nil {
		return storage.Post{}, err
	}

	post.Customer, err = s.db.Customers.GetByID(post.CustomerID)
	if err != nil {
		return storage.Post{}, err
	}

	post.Comments, err = s.db.Comments.SliceByPostID(post.ID)
	if err != nil {
		return storage.Post{}, err
	}

	for i, e := range post.Comments {
		user, err := s.db.Customers.GetByID(e.CustomerID)
		if err != nil && err != sql.ErrNoRows {
			return storage.Post{}, err
		}

		post.Comments[i].Customer = user
	}

	return post, nil
}

func (s *PostStorage) Contacts(c echo.Context) error {
	var temp struct {
		Temp
	}
	err := temp.DefaultTemp(c, s.db)
	if err != nil {
		return err
	}

	return c.Render(http.StatusOK, "contacts", temp)
}
