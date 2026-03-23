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
