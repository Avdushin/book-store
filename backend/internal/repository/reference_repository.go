package repository

import (
	"context"
	"database/sql"
	"fmt"

	"bookstore/backend/internal/models"
)

type ReferenceRepository struct {
	db *sql.DB
}

func NewReferenceRepository(db *sql.DB) *ReferenceRepository {
	return &ReferenceRepository{db: db}
}

func (r *ReferenceRepository) ListAuthors(ctx context.Context) ([]models.Author, error) {
	rows, err := r.db.QueryContext(ctx, `
		SELECT id, full_name
		FROM authors
		ORDER BY full_name ASC
	`)
	if err != nil {
		return nil, fmt.Errorf("list authors: %w", err)
	}
	defer rows.Close()

	items := make([]models.Author, 0)
	for rows.Next() {
		var item models.Author
		if err := rows.Scan(&item.ID, &item.FullName); err != nil {
			return nil, fmt.Errorf("scan author: %w", err)
		}
		items = append(items, item)
	}

	if err := rows.Err(); err != nil {
		return nil, fmt.Errorf("rows authors: %w", err)
	}

	return items, nil
}

func (r *ReferenceRepository) ListCategories(ctx context.Context) ([]models.Category, error) {
	rows, err := r.db.QueryContext(ctx, `
		SELECT id, name
		FROM categories
		ORDER BY name ASC
	`)
	if err != nil {
		return nil, fmt.Errorf("list categories: %w", err)
	}
	defer rows.Close()

	items := make([]models.Category, 0)
	for rows.Next() {
		var item models.Category
		if err := rows.Scan(&item.ID, &item.Name); err != nil {
			return nil, fmt.Errorf("scan category: %w", err)
		}
		items = append(items, item)
	}

	if err := rows.Err(); err != nil {
		return nil, fmt.Errorf("rows categories: %w", err)
	}

	return items, nil
}

func (r *ReferenceRepository) CreateAuthor(ctx context.Context, fullName string) (*models.Author, error) {
	var item models.Author

	err := r.db.QueryRowContext(
		ctx,
		`
		INSERT INTO authors (full_name)
		VALUES ($1)
		RETURNING id, full_name
		`,
		fullName,
	).Scan(&item.ID, &item.FullName)
	if err != nil {
		return nil, fmt.Errorf("create author: %w", err)
	}

	return &item, nil
}

func (r *ReferenceRepository) CreateCategory(ctx context.Context, name string) (*models.Category, error) {
	var item models.Category

	err := r.db.QueryRowContext(
		ctx,
		`
		INSERT INTO categories (name)
		VALUES ($1)
		RETURNING id, name
		`,
		name,
	).Scan(&item.ID, &item.Name)
	if err != nil {
		return nil, fmt.Errorf("create category: %w", err)
	}

	return &item, nil
}
