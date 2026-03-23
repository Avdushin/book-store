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
	"github.com/go-chi/cors"
)

func NewRouter(db *sql.DB, cfg *config.Config) http.Handler {
	r := chi.NewRouter()

	r.Use(cors.Handler(cors.Options{
		AllowedOrigins:   []string{"http://localhost:5173"},
		AllowedMethods:   []string{"GET", "POST", "PUT", "PATCH", "DELETE", "OPTIONS"},
		AllowedHeaders:   []string{"Accept", "Authorization", "Content-Type"},
		AllowCredentials: true,
	}))

	r.Use(middleware.RequestID)
	r.Use(middleware.RealIP)
	r.Use(middleware.Logger)
	r.Use(middleware.Recoverer)

	healthHandler := handlers.NewHealthHandler(db)

	bookRepo := repository.NewBookRepository(db)
	bookService := services.NewBookService(bookRepo)
	bookHandler := handlers.NewBookHandler(bookService)

	adminBookService := services.NewAdminBookService(bookRepo)
	adminBookHandler := handlers.NewAdminBookHandler(adminBookService)

	userRepo := repository.NewUserRepository(db)
	authService := services.NewAuthService(userRepo, cfg.JWTSecret)
	authHandler := handlers.NewAuthHandler(authService)

	purchaseRepo := repository.NewPurchaseRepository(db)
	purchaseService := services.NewPurchaseService(purchaseRepo, bookRepo)
	purchaseHandler := handlers.NewPurchaseHandler(purchaseService)

	rentalRepo := repository.NewRentalRepository(db)
	rentalService := services.NewRentalService(rentalRepo, bookRepo)
	rentalHandler := handlers.NewRentalHandler(rentalService)

	referenceRepo := repository.NewReferenceRepository(db)
	referenceService := services.NewReferenceService(referenceRepo)
	referenceHandler := handlers.NewReferenceHandler(referenceService)

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

		api.Get("/authors", referenceHandler.ListAuthors)
		api.Get("/categories", referenceHandler.ListCategories)

		api.Group(func(private chi.Router) {
			private.Use(appmiddleware.Auth(authService))

			private.Post("/purchases", purchaseHandler.Create)
			private.Get("/purchases/my", purchaseHandler.ListMy)

			private.Post("/rentals", rentalHandler.Create)
			private.Get("/rentals/my", rentalHandler.ListMy)
		})

		api.Route("/admin", func(admin chi.Router) {
			admin.Use(appmiddleware.Auth(authService))
			admin.Use(appmiddleware.AdminOnly)

			admin.Post("/books", adminBookHandler.Create)
			admin.Put("/books/{id}", adminBookHandler.Update)
			admin.Delete("/books/{id}", adminBookHandler.Delete)
			admin.Patch("/books/{id}/status", adminBookHandler.UpdateStatus)
			admin.Patch("/books/{id}/availability", adminBookHandler.UpdateAvailability)

			admin.Post("/authors", referenceHandler.CreateAuthor)
			admin.Post("/categories", referenceHandler.CreateCategory)
		})
	})

	return r
}
