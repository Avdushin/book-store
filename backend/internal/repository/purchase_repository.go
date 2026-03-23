package repository

import (
	"context"
	"database/sql"
	"fmt"

	"bookstore/backend/internal/models"
)

type PurchaseRepository struct {
	db *sql.DB
}

func NewPurchaseRepository(db *sql.DB) *PurchaseRepository {
	return &PurchaseRepository{db: db}
}

func (r *PurchaseRepository) Create(ctx context.Context, userID, bookID int64, price float64) (*models.Purchase, error) {
	query := `
		INSERT INTO purchases (user_id, book_id, price)
		VALUES ($1, $2, $3)
		RETURNING id, user_id, book_id, price, purchased_at
	`

	var p models.Purchase
	err := r.db.QueryRowContext(ctx, query, userID, bookID, price).Scan(
		&p.ID,
		&p.UserID,
		&p.BookID,
		&p.Price,
		&p.PurchasedAt,
	)
	if err != nil {
		return nil, fmt.Errorf("create purchase: %w", err)
	}

	return &p, nil
}

func (r *PurchaseRepository) ListByUserID(ctx context.Context, userID int64) ([]models.Purchase, error) {
	query := `
		SELECT
			p.id,
			p.user_id,
			p.book_id,
			b.title,
			p.price,
			p.purchased_at
		FROM purchases p
		JOIN books b ON b.id = p.book_id
		WHERE p.user_id = $1
		ORDER BY p.purchased_at DESC
	`

	rows, err := r.db.QueryContext(ctx, query, userID)
	if err != nil {
		return nil, fmt.Errorf("list purchases by user: %w", err)
	}
	defer rows.Close()

	items := make([]models.Purchase, 0)
	for rows.Next() {
		var p models.Purchase
		if err := rows.Scan(
			&p.ID,
			&p.UserID,
			&p.BookID,
			&p.BookTitle,
			&p.Price,
			&p.PurchasedAt,
		); err != nil {
			return nil, fmt.Errorf("scan purchase: %w", err)
		}
		items = append(items, p)
	}

	if err := rows.Err(); err != nil {
		return nil, fmt.Errorf("rows purchases: %w", err)
	}

	return items, nil
}
