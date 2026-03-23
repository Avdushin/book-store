package services

import (
	"context"
	"errors"
	"time"

	"bookstore/backend/internal/models"
	"bookstore/backend/internal/repository"
)

type RentalService struct {
	rentalRepo *repository.RentalRepository
	bookRepo   *repository.BookRepository
}

func NewRentalService(
	rentalRepo *repository.RentalRepository,
	bookRepo *repository.BookRepository,
) *RentalService {
	return &RentalService{
		rentalRepo: rentalRepo,
		bookRepo:   bookRepo,
	}
}

func (s *RentalService) Create(ctx context.Context, userID int64, req models.CreateRentalRequest) (*models.Rental, error) {
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
		return nil, errors.New("book is not available for rent")
	}

	startDate := time.Now()
	endDate, err := calculateEndDate(startDate, req.Tariff)
	if err != nil {
		return nil, err
	}

	rental, err := s.rentalRepo.Create(ctx, userID, req.BookID, req.Tariff, startDate, endDate, "active")
	if err != nil {
		return nil, err
	}

	rental.BookTitle = book.Title

	if err := s.bookRepo.UpdateAvailabilityAndStatus(ctx, req.BookID, false, "rented"); err != nil {
		return nil, err
	}

	return rental, nil
}

func (s *RentalService) ListMy(ctx context.Context, userID int64) ([]models.Rental, error) {
	return s.rentalRepo.ListByUserID(ctx, userID)
}

func calculateEndDate(start time.Time, tariff string) (time.Time, error) {
	switch tariff {
	case "2_weeks":
		return start.AddDate(0, 0, 14), nil
	case "1_month":
		return start.AddDate(0, 1, 0), nil
	case "3_months":
		return start.AddDate(0, 3, 0), nil
	default:
		return time.Time{}, errors.New("invalid tariff")
	}
}
