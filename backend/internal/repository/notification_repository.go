package repository

import (
	"context"
	"database/sql"
	"fmt"

	"bookstore/backend/internal/models"
)

type NotificationRepository struct {
	db *sql.DB
}

func NewNotificationRepository(db *sql.DB) *NotificationRepository {
	return &NotificationRepository{db: db}
}

func (r *NotificationRepository) ExistsByRentalAndType(ctx context.Context, rentalID int64, notificationType string) (bool, error) {
	query := `
		SELECT EXISTS(
			SELECT 1
			FROM notifications
			WHERE rental_id = $1 AND type = $2
		)
	`

	var exists bool
	err := r.db.QueryRowContext(ctx, query, rentalID, notificationType).Scan(&exists)
	if err != nil {
		return false, fmt.Errorf("check notification exists: %w", err)
	}

	return exists, nil
}

func (r *NotificationRepository) Create(
	ctx context.Context,
	userID, rentalID int64,
	notificationType, message, status string,
) (*models.Notification, error) {
	query := `
		INSERT INTO notifications (user_id, rental_id, type, message, status)
		VALUES ($1, $2, $3, $4, $5)
		RETURNING id, user_id, rental_id, type, message, status
	`

	var n models.Notification
	err := r.db.QueryRowContext(ctx, query, userID, rentalID, notificationType, message, status).Scan(
		&n.ID,
		&n.UserID,
		&n.RentalID,
		&n.Type,
		&n.Message,
		&n.Status,
	)
	if err != nil {
		return nil, fmt.Errorf("create notification: %w", err)
	}

	return &n, nil
}
