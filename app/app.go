package app

import (
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/samurenkoroma/lha/internal/config"
	wmLogger "github.com/samurenkoroma/lha/internal/http-server/middleware/logger"
	"github.com/samurenkoroma/lha/internal/lib/logger/sl"
	"github.com/samurenkoroma/lha/internal/storage/sqlite"
	"log/slog"
	"net/http"
	"os"
)

const (
	envLocal = "local"
	envDev   = "dev"
	envProd  = "prod"
)

type Application struct {
	Log     *slog.Logger
	Cfg     *config.Config
	Router  *chi.Mux
	Storage *sqlite.Storage
}

func NewApplication(cfg *config.Config) *Application {
	return &Application{
		Cfg:    cfg,
		Router: chi.NewRouter(),
	}
}

func (a *Application) initDB() {
	storage, err := sqlite.New(a.Cfg.Storage.Path)
	if err != nil {
		a.Log.Error("Ошибка подключения к БД.", sl.Err(err))
		os.Exit(1)
	}
	a.Storage = storage
}

func (a *Application) Configure() *Application {
	a.initDB()
	a.initLog()
	a.Router.Use(middleware.RequestID)
	//router.Use(middleware.Logger)
	a.Router.Use(wmLogger.New(a.Log))
	a.Router.Use(middleware.Recoverer)
	a.Router.Use(middleware.URLFormat)
	a.configureRouter()
	return a
}

func (a *Application) Run() {
	srv := &http.Server{
		Addr:         a.Cfg.HTTPServer.Address,
		Handler:      a.Router,
		ReadTimeout:  a.Cfg.HTTPServer.Timeout,
		WriteTimeout: a.Cfg.HTTPServer.Timeout,
		IdleTimeout:  a.Cfg.HTTPServer.IdleTimeout,
	}

	if err := srv.ListenAndServe(); err != nil {
		a.Log.Error("failed to start server")
	}
}
