package models

type CreateBookRequest struct {
	Title            string  `json:"title"`
	Description      string  `json:"description"`
	AuthorID         int64   `json:"author_id"`
	CategoryID       int64   `json:"category_id"`
	YearWritten      int     `json:"year_written"`
	PurchasePrice    float64 `json:"purchase_price"`
	RentPrice2Weeks  float64 `json:"rent_price_2_weeks"`
	RentPrice1Month  float64 `json:"rent_price_1_month"`
	RentPrice3Months float64 `json:"rent_price_3_months"`
	Status           string  `json:"status"`
	IsAvailable      bool    `json:"is_available"`
	CoverURL         string  `json:"cover_url"`
}

type UpdateBookRequest = CreateBookRequest

type UpdateBookStatusRequest struct {
	Status string `json:"status"`
}

type UpdateBookAvailabilityRequest struct {
	IsAvailable bool `json:"is_available"`
}
