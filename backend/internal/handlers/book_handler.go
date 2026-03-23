package handlers

import (
	"encoding/json"
	"net/http"
	"strconv"

	"bookstore/backend/internal/models"
	"bookstore/backend/internal/services"

	"github.com/go-chi/chi/v5"
)

type BookHandler struct {
	service *services.BookService
}

func NewBookHandler(service *services.BookService) *BookHandler {
	return &BookHandler{service: service}
}

func (h *BookHandler) List(w http.ResponseWriter, r *http.Request) {
	year, _ := strconv.Atoi(r.URL.Query().Get("year"))

	params := models.ListBooksParams{
		Category: r.URL.Query().Get("category"),
		Author:   r.URL.Query().Get("author"),
		Year:     year,
		SortBy:   r.URL.Query().Get("sort_by"),
		Order:    r.URL.Query().Get("order"),
	}

	resp, err := h.service.List(r.Context(), params)
	if err != nil {
		writeError(w, http.StatusInternalServerError, "failed to list books")
		return
	}

	writeJSON(w, http.StatusOK, resp)
}

func (h *BookHandler) GetByID(w http.ResponseWriter, r *http.Request) {
	idParam := chi.URLParam(r, "id")
	id, err := strconv.ParseInt(idParam, 10, 64)
	if err != nil {
		writeError(w, http.StatusBadRequest, "invalid book id")
		return
	}

	book, err := h.service.GetByID(r.Context(), id)
	if err != nil {
		writeError(w, http.StatusInternalServerError, "failed to get book")
		return
	}
	if book == nil {
		writeError(w, http.StatusNotFound, "book not found")
		return
	}

	writeJSON(w, http.StatusOK, book)
}

func writeJSON(w http.ResponseWriter, status int, data any) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	_ = json.NewEncoder(w).Encode(data)
}

func writeError(w http.ResponseWriter, status int, message string) {
	writeJSON(w, status, map[string]string{
		"error": message,
	})
}
