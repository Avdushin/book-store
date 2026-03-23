package handlers

import (
	"database/sql"
	"encoding/json"
	"errors"
	"net/http"
	"strconv"

	"bookstore/backend/internal/models"
	"bookstore/backend/internal/services"

	"github.com/go-chi/chi/v5"
)

type AdminBookHandler struct {
	service *services.AdminBookService
}

func NewAdminBookHandler(service *services.AdminBookService) *AdminBookHandler {
	return &AdminBookHandler{service: service}
}

func (h *AdminBookHandler) Create(w http.ResponseWriter, r *http.Request) {
	var req models.CreateBookRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		writeError(w, http.StatusBadRequest, "invalid request body")
		return
	}

	book, err := h.service.Create(r.Context(), req)
	if err != nil {
		writeError(w, http.StatusBadRequest, err.Error())
		return
	}

	writeJSON(w, http.StatusCreated, book)
}

func (h *AdminBookHandler) Update(w http.ResponseWriter, r *http.Request) {
	id, err := parseBookID(r)
	if err != nil {
		writeError(w, http.StatusBadRequest, "invalid book id")
		return
	}

	var req models.UpdateBookRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		writeError(w, http.StatusBadRequest, "invalid request body")
		return
	}

	book, err := h.service.Update(r.Context(), id, req)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			writeError(w, http.StatusNotFound, "book not found")
			return
		}
		writeError(w, http.StatusBadRequest, err.Error())
		return
	}

	writeJSON(w, http.StatusOK, book)
}

func (h *AdminBookHandler) Delete(w http.ResponseWriter, r *http.Request) {
	id, err := parseBookID(r)
	if err != nil {
		writeError(w, http.StatusBadRequest, "invalid book id")
		return
	}

	err = h.service.Delete(r.Context(), id)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			writeError(w, http.StatusNotFound, "book not found")
			return
		}
		writeError(w, http.StatusBadRequest, err.Error())
		return
	}

	writeJSON(w, http.StatusOK, map[string]string{
		"message": "book deleted",
	})
}

func (h *AdminBookHandler) UpdateStatus(w http.ResponseWriter, r *http.Request) {
	id, err := parseBookID(r)
	if err != nil {
		writeError(w, http.StatusBadRequest, "invalid book id")
		return
	}

	var req models.UpdateBookStatusRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		writeError(w, http.StatusBadRequest, "invalid request body")
		return
	}

	err = h.service.UpdateStatus(r.Context(), id, req.Status)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			writeError(w, http.StatusNotFound, "book not found")
			return
		}
		writeError(w, http.StatusBadRequest, err.Error())
		return
	}

	writeJSON(w, http.StatusOK, map[string]string{
		"message": "book status updated",
	})
}

func (h *AdminBookHandler) UpdateAvailability(w http.ResponseWriter, r *http.Request) {
	id, err := parseBookID(r)
	if err != nil {
		writeError(w, http.StatusBadRequest, "invalid book id")
		return
	}

	var req models.UpdateBookAvailabilityRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		writeError(w, http.StatusBadRequest, "invalid request body")
		return
	}

	err = h.service.UpdateAvailability(r.Context(), id, req.IsAvailable)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			writeError(w, http.StatusNotFound, "book not found")
			return
		}
		writeError(w, http.StatusBadRequest, err.Error())
		return
	}

	writeJSON(w, http.StatusOK, map[string]string{
		"message": "book availability updated",
	})
}

func parseBookID(r *http.Request) (int64, error) {
	return strconv.ParseInt(chi.URLParam(r, "id"), 10, 64)
}
