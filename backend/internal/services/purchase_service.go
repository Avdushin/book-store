package services

import (
	"context"
	"errors"

	"bookstore/backend/internal/models"
	"bookstore/backend/internal/repository"
)

type PurchaseService struct {
	purchaseRepo *repository.PurchaseRepository
	bookRepo     *repository.BookRepository
}

func NewPurchaseService(
	purchaseRepo *repository.PurchaseRepository,
	bookRepo *repository.BookRepository,
) *PurchaseService {
	return &PurchaseService{
		purchaseRepo: purchaseRepo,
		bookRepo:     bookRepo,
	}
}

func (s *PurchaseService) Create(ctx context.Context, userID int64, req models.CreatePurchaseRequest) (*models.Purchase, error) {
	if req.BookID <= 0 {
		return nil, errors.New("invalid book_id")
	}

	book, err := s.bookRepo.GetByID(ctx, req.BookID)
	if err != nil {
		return nil, err
	}
	if book == nil {
		return nil, errors.New("book not found")
	}
	if !book.IsAvailable || book.Status != "available" {
		return nil, errors.New("book is not available for purchase")
	}

	purchase, err := s.purchaseRepo.Create(ctx, userID, req.BookID, book.PurchasePrice)
	if err != nil {
		return nil, err
	}

	purchase.BookTitle = book.Title

	if err := s.bookRepo.UpdateAvailabilityAndStatus(ctx, req.BookID, false, "sold_out"); err != nil {
		return nil, err
	}

	return purchase, nil
}

func (s *PurchaseService) ListMy(ctx context.Context, userID int64) ([]models.Purchase, error) {
	return s.purchaseRepo.ListByUserID(ctx, userID)
}
