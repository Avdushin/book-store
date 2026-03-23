package handlers

import (
	"encoding/json"
	"net/http"

	"bookstore/backend/internal/models"
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

func (h *ReferenceHandler) CreateAuthor(w http.ResponseWriter, r *http.Request) {
	var req models.CreateAuthorRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		writeError(w, http.StatusBadRequest, "invalid request body")
		return
	}

	item, err := h.service.CreateAuthor(r.Context(), req)
	if err != nil {
		writeError(w, http.StatusBadRequest, err.Error())
		return
	}

	writeJSON(w, http.StatusCreated, item)
}

func (h *ReferenceHandler) CreateCategory(w http.ResponseWriter, r *http.Request) {
	var req models.CreateCategoryRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		writeError(w, http.StatusBadRequest, "invalid request body")
		return
	}

	item, err := h.service.CreateCategory(r.Context(), req)
	if err != nil {
		writeError(w, http.StatusBadRequest, err.Error())
		return
	}

	writeJSON(w, http.StatusCreated, item)
}
