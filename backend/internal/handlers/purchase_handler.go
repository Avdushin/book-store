package handlers

import (
	"encoding/json"
	"net/http"

	appmiddleware "bookstore/backend/internal/middleware"
	"bookstore/backend/internal/models"
	"bookstore/backend/internal/services"
)

type PurchaseHandler struct {
	service *services.PurchaseService
}

func NewPurchaseHandler(service *services.PurchaseService) *PurchaseHandler {
	return &PurchaseHandler{service: service}
}

func (h *PurchaseHandler) Create(w http.ResponseWriter, r *http.Request) {
	userID, ok := r.Context().Value(appmiddleware.UserIDKey).(int64)
	if !ok {
		writeError(w, http.StatusUnauthorized, "unauthorized")
		return
	}

	var req models.CreatePurchaseRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		writeError(w, http.StatusBadRequest, "invalid request body")
		return
	}

	purchase, err := h.service.Create(r.Context(), userID, req)
	if err != nil {
		writeError(w, http.StatusBadRequest, err.Error())
		return
	}

	writeJSON(w, http.StatusCreated, purchase)
}

func (h *PurchaseHandler) ListMy(w http.ResponseWriter, r *http.Request) {
	userID, ok := r.Context().Value(appmiddleware.UserIDKey).(int64)
	if !ok {
		writeError(w, http.StatusUnauthorized, "unauthorized")
		return
	}

	items, err := h.service.ListMy(r.Context(), userID)
	if err != nil {
		writeError(w, http.StatusInternalServerError, "failed to list purchases")
		return
	}

	writeJSON(w, http.StatusOK, map[string]any{
		"items": items,
		"total": len(items),
	})
}
