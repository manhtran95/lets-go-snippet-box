package main

import (
	// New import

	"snippetbox.mtran.io/internal/models"
)

type templateData struct {
	Snippet  *models.Snippet
	Snippets []*models.Snippet
}
