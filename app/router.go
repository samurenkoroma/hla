package app

import (
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/samurenkoroma/lha/internal/http-server/handlers/book"
	"github.com/samurenkoroma/lha/internal/http-server/handlers/file/upload"
	"github.com/samurenkoroma/lha/internal/http-server/handlers/url/redirect"
	"github.com/samurenkoroma/lha/internal/http-server/handlers/url/save"
)

func (a *Application) configureRouter() {
	a.Router.Route("/file", func(r chi.Router) {
		r.Post("/upload", upload.New(a.Log))
	})

	a.Router.Route("/books", func(r chi.Router) {
		r.Get("/", book.List)
		r.Get("/{id}", book.ById)
		r.Post("/", book.Create)
	})

	a.Router.Route("/url", func(r chi.Router) {
		r.Use(middleware.BasicAuth("hla", map[string]string{
			a.Cfg.HTTPServer.User: a.Cfg.HTTPServer.Password,
		}))
		r.Post("/", save.New(a.Log, a.Storage))
		r.Get("/{alias}", redirect.New(a.Log, a.Storage))
	})
}
