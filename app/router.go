package app

import (
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/samurenkoroma/lha/internal/http-server/handlers/file/upload"
	"github.com/samurenkoroma/lha/internal/http-server/handlers/library"
	"github.com/samurenkoroma/lha/internal/http-server/handlers/url/redirect"
	"github.com/samurenkoroma/lha/internal/http-server/handlers/url/save"
)

func (a *Application) configureRouter() {
	a.Router.Use(middleware.BasicAuth("hla", map[string]string{
		a.Cfg.HTTPServer.User: a.Cfg.HTTPServer.Password,
	}))
	a.Router.Route("/file", func(r chi.Router) {
		r.Post("/upload", upload.New(a.Log))
	})

	a.Router.Route("/books", func(r chi.Router) {
		r.Get("/", library.List)
		r.Get("/{id}", library.ById)
		r.Post("/", library.Create)
	})

	a.Router.Route("/url", func(r chi.Router) {
		r.Post("/", save.New(a.Log, a.Storage))
		r.Get("/{alias}", redirect.New(a.Log, a.Storage))
	})
}
