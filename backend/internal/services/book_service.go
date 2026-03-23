package services

import (
	"context"
	"strings"

	"bookstore/backend/internal/models"
	"bookstore/backend/internal/repository"
)

type BookService struct {
	repo *repository.BookRepository
}

func NewBookService(repo *repository.BookRepository) *BookService {
	return &BookService{repo: repo}
}

func (s *BookService) List(ctx context.Context, params models.ListBooksParams) (*models.BookListResponse, error) {
	params.Category = strings.TrimSpace(params.Category)
	params.Author = strings.TrimSpace(params.Author)
	params.SortBy = strings.TrimSpace(params.SortBy)
	params.Order = strings.TrimSpace(params.Order)

	books, err := s.repo.List(ctx, params)
	if err != nil {
		return nil, err
	}

	return &models.BookListResponse{
		Items: books,
		Total: len(books),
	}, nil
}

func (s *BookService) GetByID(ctx context.Context, id int64) (*models.Book, error) {
	return s.repo.GetByID(ctx, id)
}
