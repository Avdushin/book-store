package models

type CreatePurchaseRequest struct {
	BookID int64 `json:"book_id"`
}

type Purchase struct {
	ID          int64   `json:"id"`
	UserID      int64   `json:"user_id"`
	BookID      int64   `json:"book_id"`
	BookTitle   string  `json:"book_title"`
	Price       float64 `json:"price"`
	PurchasedAt string  `json:"purchased_at"`
}
