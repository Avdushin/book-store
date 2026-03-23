package services

import (
	"context"
	"errors"
	"strings"

	"bookstore/backend/internal/models"
	"bookstore/backend/internal/repository"
)

type ReferenceService struct {
	repo *repository.ReferenceRepository
}

func NewReferenceService(repo *repository.ReferenceRepository) *ReferenceService {
	return &ReferenceService{repo: repo}
}

func (s *ReferenceService) ListAuthors(ctx context.Context) ([]models.Author, error) {
	return s.repo.ListAuthors(ctx)
}

func (s *ReferenceService) ListCategories(ctx context.Context) ([]models.Category, error) {
	return s.repo.ListCategories(ctx)
}

func (s *ReferenceService) CreateAuthor(ctx context.Context, req models.CreateAuthorRequest) (*models.Author, error) {
	req.FullName = strings.TrimSpace(req.FullName)
	if req.FullName == "" {
		return nil, errors.New("full_name is required")
	}

	return s.repo.CreateAuthor(ctx, req.FullName)
}

func (s *ReferenceService) CreateCategory(ctx context.Context, req models.CreateCategoryRequest) (*models.Category, error) {
	req.Name = strings.TrimSpace(req.Name)
	if req.Name == "" {
		return nil, errors.New("name is required")
	}

	return s.repo.CreateCategory(ctx, req.Name)
}
