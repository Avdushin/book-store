package server

import (
	"database/sql"
	"net/http"

	"bookstore/backend/internal/config"
	"bookstore/backend/internal/handlers"
	appmiddleware "bookstore/backend/internal/middleware"
	"bookstore/backend/internal/repository"
	"bookstore/backend/internal/services"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
)

func NewRouter(db *sql.DB, cfg *config.Config) http.Handler {
	r := chi.NewRouter()

	r.Use(middleware.RequestID)
	r.Use(middleware.RealIP)
	r.Use(middleware.Logger)
	r.Use(middleware.Recoverer)

	healthHandler := handlers.NewHealthHandler(db)

	bookRepo := repository.NewBookRepository(db)
	bookService := services.NewBookService(bookRepo)
	bookHandler := handlers.NewBookHandler(bookService)

	userRepo := repository.NewUserRepository(db)
	authService := services.NewAuthService(userRepo, cfg.JWTSecret)
	authHandler := handlers.NewAuthHandler(authService)

	purchaseRepo := repository.NewPurchaseRepository(db)
	purchaseService := services.NewPurchaseService(purchaseRepo, bookRepo)
	purchaseHandler := handlers.NewPurchaseHandler(purchaseService)

	rentalRepo := repository.NewRentalRepository(db)
	rentalService := services.NewRentalService(rentalRepo, bookRepo)
	rentalHandler := handlers.NewRentalHandler(rentalService)

	r.Get("/health", healthHandler.Health)
	r.Handle("/covers/*", http.StripPrefix("/covers/", http.FileServer(http.Dir("./static/covers"))))

	r.Route("/api", func(api chi.Router) {
		api.Route("/auth", func(auth chi.Router) {
			auth.Post("/register", authHandler.Register)
			auth.Post("/login", authHandler.Login)

			auth.Group(func(private chi.Router) {
				private.Use(appmiddleware.Auth(authService))
				private.Get("/me", authHandler.Me)
			})
		})

		api.Get("/books", bookHandler.List)
		api.Get("/books/{id}", bookHandler.GetByID)

		api.Group(func(private chi.Router) {
			private.Use(appmiddleware.Auth(authService))

			private.Post("/purchases", purchaseHandler.Create)
			private.Get("/purchases/my", purchaseHandler.ListMy)

			private.Post("/rentals", rentalHandler.Create)
			private.Get("/rentals/my", rentalHandler.ListMy)
		})
	})

	return r
}
