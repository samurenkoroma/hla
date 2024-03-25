package adapters

import (
	"context"
	"github.com/samurenkoroma/lha/internal/models"
)

type BookRepository interface {
	List(ctx context.Context) ([]models.Book, error)
	ById(ctx context.Context, id string) (models.Book, error)
	Create(ctx context.Context) error
}

var Books = []models.Book{
	{
		ID:      "aeSESEz",
		Title:   "Azbooka",
		Authors: []models.Author{{Name: "awdawd"}},
		Tags:    []string{"python", "devops", "flask"},
		Resources: []models.File{{
			Path: "./awdawdw",
			Type: "pdf",
		}},
	},
	{
		ID:      "AEdrvz",
		Title:   "Azbooka",
		Authors: []models.Author{{Name: "awdawd"}},
		Tags:    []string{"python", "devops", "flask"},
		Resources: []models.File{{
			Path: "./awdawdw",
			Type: "pdf",
		}},
	},
}
