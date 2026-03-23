package services

import (
	"context"

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
