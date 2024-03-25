package app

import (
	"github.com/samurenkoroma/lha/internal/lib/logger/handlers/slogpretty"
	"log/slog"
	"os"
)

func (a *Application) initLog() {
	a.Log = setupLogger(a.Cfg.Env)
	a.Log = a.Log.With(slog.String("env", a.Cfg.Env)) // к каждому сообщению будет добавляться поле с информацией о текущем окружении

	a.Log.Info("initializing server", slog.String("address", a.Cfg.HTTPServer.Address)) // Помимо сообщения выведем параметр с адресом
	a.Log.Debug("logger debug mode enabled")

}

func setupLogger(env string) *slog.Logger {
	var log *slog.Logger

	switch env {
	case envLocal:
		log = setupPrettySlog()
	case envDev:
		log = slog.New(
			slog.NewJSONHandler(os.Stdout, &slog.HandlerOptions{Level: slog.LevelDebug}),
		)
	case envProd:
		log = slog.New(
			slog.NewJSONHandler(os.Stdout, &slog.HandlerOptions{Level: slog.LevelInfo}),
		)
	}

	return log
}
func setupPrettySlog() *slog.Logger {
	opts := slogpretty.PrettyHandlerOptions{
		SlogOpts: &slog.HandlerOptions{
			Level: slog.LevelDebug,
		},
	}

	handler := opts.NewPrettyHandler(os.Stdout)

	return slog.New(handler)
}
