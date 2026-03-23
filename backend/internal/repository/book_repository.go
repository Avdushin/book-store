package repository

import (
	"context"
	"database/sql"
	"fmt"
	"strings"

	"bookstore/backend/internal/models"
)

type BookRepository struct {
	db *sql.DB
}

func NewBookRepository(db *sql.DB) *BookRepository {
	return &BookRepository{db: db}
}

func (r *BookRepository) List(ctx context.Context, params models.ListBooksParams) ([]models.Book, error) {
	baseQuery := `
SELECT
	b.id,
	b.title,
	COALESCE(b.description, ''),
	b.author_id,
	a.full_name,
	b.category_id,
	c.name,
	b.year_written,
	b.purchase_price,
	b.rent_price_2_weeks,
	b.rent_price_1_month,
	b.rent_price_3_months,
	b.status,
	b.is_available,
	COALESCE(b.cover_url, '')
FROM books b
JOIN authors a ON a.id = b.author_id
JOIN categories c ON c.id = b.category_id
WHERE 1=1
`
	args := make([]any, 0)
	argPos := 1

	if params.Category != "" {
		baseQuery += fmt.Sprintf(" AND c.name ILIKE $%d", argPos)
		args = append(args, params.Category)
		argPos++
	}

	if params.Author != "" {
		baseQuery += fmt.Sprintf(" AND a.full_name ILIKE $%d", argPos)
		args = append(args, params.Author)
		argPos++
	}

	if params.Year != 0 {
		baseQuery += fmt.Sprintf(" AND b.year_written = $%d", argPos)
		args = append(args, params.Year)
		argPos++
	}

	sortColumn := "b.id"
	switch strings.ToLower(params.SortBy) {
	case "title":
		sortColumn = "b.title"
	case "year":
		sortColumn = "b.year_written"
	case "price":
		sortColumn = "b.purchase_price"
	case "author":
		sortColumn = "a.full_name"
	case "category":
		sortColumn = "c.name"
	}

	order := "ASC"
	if strings.ToUpper(params.Order) == "DESC" {
		order = "DESC"
	}

	baseQuery += fmt.Sprintf(" ORDER BY %s %s", sortColumn, order)

	rows, err := r.db.QueryContext(ctx, baseQuery, args...)
	if err != nil {
		return nil, fmt.Errorf("query books: %w", err)
	}
	defer rows.Close()

	books := make([]models.Book, 0)
	for rows.Next() {
		var b models.Book
		if err := rows.Scan(
			&b.ID,
			&b.Title,
			&b.Description,
			&b.AuthorID,
			&b.AuthorName,
			&b.CategoryID,
			&b.CategoryName,
			&b.YearWritten,
			&b.PurchasePrice,
			&b.RentPrice2Weeks,
			&b.RentPrice1Month,
			&b.RentPrice3Months,
			&b.Status,
			&b.IsAvailable,
			&b.CoverURL,
		); err != nil {
			return nil, fmt.Errorf("scan book: %w", err)
		}
		books = append(books, b)
	}

	if err := rows.Err(); err != nil {
		return nil, fmt.Errorf("rows books: %w", err)
	}

	return books, nil
}

func (r *BookRepository) GetByID(ctx context.Context, id int64) (*models.Book, error) {
	query := `
SELECT
	b.id,
	b.title,
	COALESCE(b.description, ''),
	b.author_id,
	a.full_name,
	b.category_id,
	c.name,
	b.year_written,
	b.purchase_price,
	b.rent_price_2_weeks,
	b.rent_price_1_month,
	b.rent_price_3_months,
	b.status,
	b.is_available,
	COALESCE(b.cover_url, '')
FROM books b
JOIN authors a ON a.id = b.author_id
JOIN categories c ON c.id = b.category_id
WHERE b.id = $1
`
	var b models.Book
	err := r.db.QueryRowContext(ctx, query, id).Scan(
		&b.ID,
		&b.Title,
		&b.Description,
		&b.AuthorID,
		&b.AuthorName,
		&b.CategoryID,
		&b.CategoryName,
		&b.YearWritten,
		&b.PurchasePrice,
		&b.RentPrice2Weeks,
		&b.RentPrice1Month,
		&b.RentPrice3Months,
		&b.Status,
		&b.IsAvailable,
		&b.CoverURL,
	)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, nil
		}
		return nil, fmt.Errorf("get book by id: %w", err)
	}

	return &b, nil
}

func (r *BookRepository) UpdateAvailabilityAndStatus(ctx context.Context, id int64, isAvailable bool, status string) error {
	query := `
		UPDATE books
		SET is_available = $1, status = $2
		WHERE id = $3
	`

	result, err := r.db.ExecContext(ctx, query, isAvailable, status, id)
	if err != nil {
		return fmt.Errorf("update book availability and status: %w", err)
	}

	rowsAffected, err := result.RowsAffected()
	if err != nil {
		return fmt.Errorf("rows affected: %w", err)
	}
	if rowsAffected == 0 {
		return sql.ErrNoRows
	}

	return nil
}

func (r *BookRepository) Create(ctx context.Context, req models.CreateBookRequest) (*models.Book, error) {
	query := `
		INSERT INTO books (
			title,
			description,
			author_id,
			category_id,
			year_written,
			purchase_price,
			rent_price_2_weeks,
			rent_price_1_month,
			rent_price_3_months,
			status,
			is_available,
			cover_url
		)
		VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11, $12)
		RETURNING
			id,
			title,
			COALESCE(description, ''),
			author_id,
			category_id,
			year_written,
			purchase_price,
			rent_price_2_weeks,
			rent_price_1_month,
			rent_price_3_months,
			status,
			is_available,
			COALESCE(cover_url, '')
	`

	var book models.Book
	err := r.db.QueryRowContext(
		ctx,
		query,
		req.Title,
		req.Description,
		req.AuthorID,
		req.CategoryID,
		req.YearWritten,
		req.PurchasePrice,
		req.RentPrice2Weeks,
		req.RentPrice1Month,
		req.RentPrice3Months,
		req.Status,
		req.IsAvailable,
		req.CoverURL,
	).Scan(
		&book.ID,
		&book.Title,
		&book.Description,
		&book.AuthorID,
		&book.CategoryID,
		&book.YearWritten,
		&book.PurchasePrice,
		&book.RentPrice2Weeks,
		&book.RentPrice1Month,
		&book.RentPrice3Months,
		&book.Status,
		&book.IsAvailable,
		&book.CoverURL,
	)
	if err != nil {
		return nil, fmt.Errorf("create book: %w", err)
	}

	return r.GetByID(ctx, book.ID)
}

func (r *BookRepository) Update(ctx context.Context, id int64, req models.UpdateBookRequest) (*models.Book, error) {
	query := `
		UPDATE books
		SET
			title = $1,
			description = $2,
			author_id = $3,
			category_id = $4,
			year_written = $5,
			purchase_price = $6,
			rent_price_2_weeks = $7,
			rent_price_1_month = $8,
			rent_price_3_months = $9,
			status = $10,
			is_available = $11,
			cover_url = $12
		WHERE id = $13
	`

	result, err := r.db.ExecContext(
		ctx,
		query,
		req.Title,
		req.Description,
		req.AuthorID,
		req.CategoryID,
		req.YearWritten,
		req.PurchasePrice,
		req.RentPrice2Weeks,
		req.RentPrice1Month,
		req.RentPrice3Months,
		req.Status,
		req.IsAvailable,
		req.CoverURL,
		id,
	)
	if err != nil {
		return nil, fmt.Errorf("update book: %w", err)
	}

	rowsAffected, err := result.RowsAffected()
	if err != nil {
		return nil, fmt.Errorf("rows affected: %w", err)
	}
	if rowsAffected == 0 {
		return nil, sql.ErrNoRows
	}

	return r.GetByID(ctx, id)
}

func (r *BookRepository) Delete(ctx context.Context, id int64) error {
	result, err := r.db.ExecContext(ctx, `DELETE FROM books WHERE id = $1`, id)
	if err != nil {
		return fmt.Errorf("delete book: %w", err)
	}

	rowsAffected, err := result.RowsAffected()
	if err != nil {
		return fmt.Errorf("rows affected: %w", err)
	}
	if rowsAffected == 0 {
		return sql.ErrNoRows
	}

	return nil
}

func (r *BookRepository) UpdateStatus(ctx context.Context, id int64, status string) error {
	result, err := r.db.ExecContext(ctx, `UPDATE books SET status = $1 WHERE id = $2`, status, id)
	if err != nil {
		return fmt.Errorf("update book status: %w", err)
	}

	rowsAffected, err := result.RowsAffected()
	if err != nil {
		return fmt.Errorf("rows affected: %w", err)
	}
	if rowsAffected == 0 {
		return sql.ErrNoRows
	}

	return nil
}

func (r *BookRepository) UpdateAvailability(ctx context.Context, id int64, isAvailable bool) error {
	result, err := r.db.ExecContext(ctx, `UPDATE books SET is_available = $1 WHERE id = $2`, isAvailable, id)
	if err != nil {
		return fmt.Errorf("update book availability: %w", err)
	}

	rowsAffected, err := result.RowsAffected()
	if err != nil {
		return fmt.Errorf("rows affected: %w", err)
	}
	if rowsAffected == 0 {
		return sql.ErrNoRows
	}

	return nil
}
