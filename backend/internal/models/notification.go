package models

type Notification struct {
	ID       int64  `json:"id"`
	UserID   int64  `json:"user_id"`
	RentalID int64  `json:"rental_id"`
	Type     string `json:"type"`
	Message  string `json:"message"`
	SentAt   string `json:"sent_at,omitempty"`
	Status   string `json:"status"`
}
