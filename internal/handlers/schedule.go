package handlers

import (
	"net/http"
	"webSocketEg/internal/helpers"
)

// ListEntries lists schedule entries
func (repo *DBRepo) ListEntries(w http.ResponseWriter, r *http.Request) {
	err := helpers.RenderPage(w, r, "schedule", nil, nil)
	if err != nil {
		printTemplateError(w, err)
	}
}
