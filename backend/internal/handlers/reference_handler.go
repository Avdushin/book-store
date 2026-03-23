package handlers

import (
	"net/http"

	"bookstore/backend/internal/services"
)

type ReferenceHandler struct {
	service *services.ReferenceService
}

func NewReferenceHandler(service *services.ReferenceService) *ReferenceHandler {
	return &ReferenceHandler{service: service}
}

func (h *ReferenceHandler) ListAuthors(w http.ResponseWriter, r *http.Request) {
	items, err := h.service.ListAuthors(r.Context())
	if err != nil {
		writeError(w, http.StatusInternalServerError, "failed to list authors")
		return
	}

	writeJSON(w, http.StatusOK, map[string]any{
		"items": items,
		"total": len(items),
	})
}

func (h *ReferenceHandler) ListCategories(w http.ResponseWriter, r *http.Request) {
	items, err := h.service.ListCategories(r.Context())
	if err != nil {
		writeError(w, http.StatusInternalServerError, "failed to list categories")
		return
	}

	writeJSON(w, http.StatusOK, map[string]any{
		"items": items,
		"total": len(items),
	})
}
