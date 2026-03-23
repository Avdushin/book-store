package server

import (
	"database/sql"
	"net/http"

	"bookstore/backend/internal/handlers"
	"bookstore/backend/internal/repository"
	"bookstore/backend/internal/services"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
)

func NewRouter(db *sql.DB) http.Handler {
	r := chi.NewRouter()

	r.Use(middleware.RequestID)
	r.Use(middleware.RealIP)
	r.Use(middleware.Logger)
	r.Use(middleware.Recoverer)

	healthHandler := handlers.NewHealthHandler(db)

	bookRepo := repository.NewBookRepository(db)
	bookService := services.NewBookService(bookRepo)
	bookHandler := handlers.NewBookHandler(bookService)

	r.Get("/health", healthHandler.Health)
	r.Handle("/covers/*", http.StripPrefix("/covers/", http.FileServer(http.Dir("./static/covers"))))

	r.Route("/api", func(api chi.Router) {
		api.Get("/books", bookHandler.List)
		api.Get("/books/{id}", bookHandler.GetByID)
	})

	return r
}
