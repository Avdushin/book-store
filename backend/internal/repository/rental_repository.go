package repository

import (
	"context"
	"database/sql"
	"fmt"
	"time"

	"bookstore/backend/internal/models"
)

type RentalRepository struct {
	db *sql.DB
}

func NewRentalRepository(db *sql.DB) *RentalRepository {
	return &RentalRepository{db: db}
}

func (r *RentalRepository) Create(
	ctx context.Context,
	userID, bookID int64,
	tariff string,
	startDate, endDate time.Time,
	status string,
) (*models.Rental, error) {
	query := `
		INSERT INTO rentals (user_id, book_id, tariff, start_date, end_date, status)
		VALUES ($1, $2, $3, $4, $5, $6)
		RETURNING id, user_id, book_id, tariff, start_date, end_date, status, created_at
	`

	var rental models.Rental
	err := r.db.QueryRowContext(ctx, query, userID, bookID, tariff, startDate, endDate, status).Scan(
		&rental.ID,
		&rental.UserID,
		&rental.BookID,
		&rental.Tariff,
		&rental.StartDate,
		&rental.EndDate,
		&rental.Status,
		&rental.CreatedAt,
	)
	if err != nil {
		return nil, fmt.Errorf("create rental: %w", err)
	}

	return &rental, nil
}

func (r *RentalRepository) ListByUserID(ctx context.Context, userID int64) ([]models.Rental, error) {
	query := `
		SELECT
			r.id,
			r.user_id,
			r.book_id,
			b.title,
			r.tariff,
			r.start_date,
			r.end_date,
			r.status,
			r.created_at
		FROM rentals r
		JOIN books b ON b.id = r.book_id
		WHERE r.user_id = $1
		ORDER BY r.created_at DESC
	`

	rows, err := r.db.QueryContext(ctx, query, userID)
	if err != nil {
		return nil, fmt.Errorf("list rentals by user: %w", err)
	}
	defer rows.Close()

	items := make([]models.Rental, 0)
	for rows.Next() {
		var rental models.Rental
		if err := rows.Scan(
			&rental.ID,
			&rental.UserID,
			&rental.BookID,
			&rental.BookTitle,
			&rental.Tariff,
			&rental.StartDate,
			&rental.EndDate,
			&rental.Status,
			&rental.CreatedAt,
		); err != nil {
			return nil, fmt.Errorf("scan rental: %w", err)
		}
		items = append(items, rental)
	}

	if err := rows.Err(); err != nil {
		return nil, fmt.Errorf("rows rentals: %w", err)
	}

	return items, nil
}

func (r *RentalRepository) FindExpiringSoon(ctx context.Context) ([]models.Rental, error) {
	query := `
		SELECT
			r.id,
			r.user_id,
			r.book_id,
			b.title,
			r.tariff,
			r.start_date,
			r.end_date,
			r.status,
			r.created_at
		FROM rentals r
		JOIN books b ON b.id = r.book_id
		WHERE r.status = 'active'
		  AND r.end_date > NOW()
		  AND r.end_date <= NOW() + INTERVAL '1 day'
		ORDER BY r.end_date ASC
	`

	rows, err := r.db.QueryContext(ctx, query)
	if err != nil {
		return nil, fmt.Errorf("find expiring rentals: %w", err)
	}
	defer rows.Close()

	items := make([]models.Rental, 0)
	for rows.Next() {
		var rental models.Rental
		if err := rows.Scan(
			&rental.ID,
			&rental.UserID,
			&rental.BookID,
			&rental.BookTitle,
			&rental.Tariff,
			&rental.StartDate,
			&rental.EndDate,
			&rental.Status,
			&rental.CreatedAt,
		); err != nil {
			return nil, fmt.Errorf("scan expiring rental: %w", err)
		}
		items = append(items, rental)
	}

	if err := rows.Err(); err != nil {
		return nil, fmt.Errorf("rows expiring rentals: %w", err)
	}

	return items, nil
}

func (r *RentalRepository) FindExpiredActive(ctx context.Context) ([]models.Rental, error) {
	query := `
		SELECT
			r.id,
			r.user_id,
			r.book_id,
			b.title,
			r.tariff,
			r.start_date,
			r.end_date,
			r.status,
			r.created_at
		FROM rentals r
		JOIN books b ON b.id = r.book_id
		WHERE r.status = 'active'
		  AND r.end_date <= NOW()
		ORDER BY r.end_date ASC
	`

	rows, err := r.db.QueryContext(ctx, query)
	if err != nil {
		return nil, fmt.Errorf("find expired rentals: %w", err)
	}
	defer rows.Close()

	items := make([]models.Rental, 0)
	for rows.Next() {
		var rental models.Rental
		if err := rows.Scan(
			&rental.ID,
			&rental.UserID,
			&rental.BookID,
			&rental.BookTitle,
			&rental.Tariff,
			&rental.StartDate,
			&rental.EndDate,
			&rental.Status,
			&rental.CreatedAt,
		); err != nil {
			return nil, fmt.Errorf("scan expired rental: %w", err)
		}
		items = append(items, rental)
	}

	if err := rows.Err(); err != nil {
		return nil, fmt.Errorf("rows expired rentals: %w", err)
	}

	return items, nil
}

func (r *RentalRepository) MarkExpired(ctx context.Context, rentalID int64) error {
	_, err := r.db.ExecContext(ctx, `
		UPDATE rentals
		SET status = 'expired'
		WHERE id = $1
	`, rentalID)
	if err != nil {
		return fmt.Errorf("mark rental expired: %w", err)
	}

	return nil
}
