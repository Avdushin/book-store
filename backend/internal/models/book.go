package models

type Book struct {
	ID               int64   `json:"id"`
	Title            string  `json:"title"`
	Description      string  `json:"description"`
	AuthorID         int64   `json:"author_id"`
	AuthorName       string  `json:"author_name"`
	CategoryID       int64   `json:"category_id"`
	CategoryName     string  `json:"category_name"`
	YearWritten      int     `json:"year_written"`
	PurchasePrice    float64 `json:"purchase_price"`
	RentPrice2Weeks  float64 `json:"rent_price_2_weeks"`
	RentPrice1Month  float64 `json:"rent_price_1_month"`
	RentPrice3Months float64 `json:"rent_price_3_months"`
	Status           string  `json:"status"`
	IsAvailable      bool    `json:"is_available"`
	CoverURL         string  `json:"cover_url"`
	CreatedAt        string  `json:"created_at,omitempty"`
	UpdatedAt        string  `json:"updated_at,omitempty"`
}

type ListBooksParams struct {
	Category string
	Author   string
	Year     int
	SortBy   string
	Order    string
}

type BookListResponse struct {
	Items []Book `json:"items"`
	Total int    `json:"total"`
}
