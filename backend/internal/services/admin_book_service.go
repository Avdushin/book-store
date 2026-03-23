package services

import (
	"context"
	"database/sql"
	"errors"
	"strings"

	"bookstore/backend/internal/models"
	"bookstore/backend/internal/repository"
)

type AdminBookService struct {
	repo *repository.BookRepository
}

func NewAdminBookService(repo *repository.BookRepository) *AdminBookService {
	return &AdminBookService{repo: repo}
}

func (s *AdminBookService) Create(ctx context.Context, req models.CreateBookRequest) (*models.Book, error) {
	req.Title = strings.TrimSpace(req.Title)
	req.Description = strings.TrimSpace(req.Description)
	req.Status = strings.TrimSpace(req.Status)
	req.CoverURL = strings.TrimSpace(req.CoverURL)

	if err := validateBookPayload(req.Title, req.AuthorID, req.CategoryID, req.YearWritten, req.PurchasePrice, req.RentPrice2Weeks, req.RentPrice1Month, req.RentPrice3Months, req.Status); err != nil {
		return nil, err
	}

	return s.repo.Create(ctx, req)
}

func (s *AdminBookService) Update(ctx context.Context, id int64, req models.UpdateBookRequest) (*models.Book, error) {
	req.Title = strings.TrimSpace(req.Title)
	req.Description = strings.TrimSpace(req.Description)
	req.Status = strings.TrimSpace(req.Status)
	req.CoverURL = strings.TrimSpace(req.CoverURL)

	if id <= 0 {
		return nil, errors.New("invalid book id")
	}
	if err := validateBookPayload(req.Title, req.AuthorID, req.CategoryID, req.YearWritten, req.PurchasePrice, req.RentPrice2Weeks, req.RentPrice1Month, req.RentPrice3Months, req.Status); err != nil {
		return nil, err
	}

	return s.repo.Update(ctx, id, req)
}

func (s *AdminBookService) Delete(ctx context.Context, id int64) error {
	if id <= 0 {
		return errors.New("invalid book id")
	}
	return s.repo.Delete(ctx, id)
}

func (s *AdminBookService) UpdateStatus(ctx context.Context, id int64, status string) error {
	if id <= 0 {
		return errors.New("invalid book id")
	}

	status = strings.TrimSpace(status)
	switch status {
	case "available", "rented", "sold_out", "inactive":
		return s.repo.UpdateStatus(ctx, id, status)
	default:
		return errors.New("invalid status")
	}
}

func (s *AdminBookService) UpdateAvailability(ctx context.Context, id int64, isAvailable bool) error {
	if id <= 0 {
		return errors.New("invalid book id")
	}
	return s.repo.UpdateAvailability(ctx, id, isAvailable)
}

func validateBookPayload(
	title string,
	authorID, categoryID int64,
	yearWritten int,
	purchasePrice, rent2w, rent1m, rent3m float64,
	status string,
) error {
	if title == "" {
		return errors.New("title is required")
	}
	if authorID <= 0 {
		return errors.New("author_id is required")
	}
	if categoryID <= 0 {
		return errors.New("category_id is required")
	}
	if yearWritten <= 0 {
		return errors.New("year_written must be greater than 0")
	}
	if purchasePrice < 0 || rent2w < 0 || rent1m < 0 || rent3m < 0 {
		return errors.New("prices must be greater than or equal to 0")
	}
	switch status {
	case "available", "rented", "sold_out", "inactive":
		return nil
	default:
		return errors.New("invalid status")
	}
}

var _ = sql.ErrNoRows
