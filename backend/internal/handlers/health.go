package handlers

import (
	"database/sql"
	"encoding/json"
	"net/http"
	"time"
)

type HealthHandler struct {
	DB *sql.DB
}

type healthResponse struct {
	Status    string `json:"status"`
	App       string `json:"app"`
	Database  string `json:"database"`
	Timestamp string `json:"timestamp"`
}

func NewHealthHandler(db *sql.DB) *HealthHandler {
	return &HealthHandler{DB: db}
}

func (h *HealthHandler) Health(w http.ResponseWriter, r *http.Request) {
	dbStatus := "up"
	if err := h.DB.PingContext(r.Context()); err != nil {
		dbStatus = "down"
	}

	resp := healthResponse{
		Status:    "ok",
		App:       "bookstore-backend",
		Database:  dbStatus,
		Timestamp: time.Now().Format(time.RFC3339),
	}

	w.Header().Set("Content-Type", "application/json")
	_ = json.NewEncoder(w).Encode(resp)
}
