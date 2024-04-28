package create

import (
	"github.com/go-chi/chi/v5/middleware"
	"github.com/samurenkoroma/lha/internal/models"
	"log/slog"
	"net/http"
)

type BookRequest struct {
	Title   string   `json:"title" validate:"required"`
	Authors []string `json:"authors,omitempty"`
	Tags    []string `json:"tags"`
}

type BookStorage interface {
	SaveBook(book models.Book) (int64, error)
}

func NewBook(log *slog.Logger, storage BookStorage) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		const op = "handlers.library.create.New"
		log = log.With(
			slog.String("op", op),
			slog.String("request_id", middleware.GetReqID(r.Context())),
		)

	}
}
