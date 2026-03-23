package handlers

import (
	"encoding/json"
	"net/http"

	appmiddleware "bookstore/backend/internal/middleware"
	"bookstore/backend/internal/models"
	"bookstore/backend/internal/services"
)

type RentalHandler struct {
	service *services.RentalService
}

func NewRentalHandler(service *services.RentalService) *RentalHandler {
	return &RentalHandler{service: service}
}

func (h *RentalHandler) Create(w http.ResponseWriter, r *http.Request) {
	userID, ok := r.Context().Value(appmiddleware.UserIDKey).(int64)
	if !ok {
		writeError(w, http.StatusUnauthorized, "unauthorized")
		return
	}

	var req models.CreateRentalRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		writeError(w, http.StatusBadRequest, "invalid request body")
		return
	}

	rental, err := h.service.Create(r.Context(), userID, req)
	if err != nil {
		writeError(w, http.StatusBadRequest, err.Error())
		return
	}

	writeJSON(w, http.StatusCreated, rental)
}

func (h *RentalHandler) ListMy(w http.ResponseWriter, r *http.Request) {
	userID, ok := r.Context().Value(appmiddleware.UserIDKey).(int64)
	if !ok {
		writeError(w, http.StatusUnauthorized, "unauthorized")
		return
	}

	items, err := h.service.ListMy(r.Context(), userID)
	if err != nil {
		writeError(w, http.StatusInternalServerError, "failed to list rentals")
		return
	}

	writeJSON(w, http.StatusOK, map[string]any{
		"items": items,
		"total": len(items),
	})
}
