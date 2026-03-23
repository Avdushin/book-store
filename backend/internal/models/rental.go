package models

type CreateRentalRequest struct {
	BookID int64  `json:"book_id"`
	Tariff string `json:"tariff"`
}

type Rental struct {
	ID        int64  `json:"id"`
	UserID    int64  `json:"user_id"`
	BookID    int64  `json:"book_id"`
	BookTitle string `json:"book_title"`
	Tariff    string `json:"tariff"`
	StartDate string `json:"start_date"`
	EndDate   string `json:"end_date"`
	Status    string `json:"status"`
}
