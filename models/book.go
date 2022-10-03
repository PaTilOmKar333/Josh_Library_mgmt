package models

import "time"

type Book struct {
	BookID     int        `json:"id" db:"book_id"`
	BookName   string     `json:"name" db:"book_name"`
	BookStatus string     `json:"status" db:"status"`
	Price      int        `json:"price" db:"price"`
	Category   string     `json:"category" db:"category"`
	CreatedBy  string     `json:"created_by,omitempty" db:"created_by"`
	CreatedAT  *time.Time `json:"created_at,omitempty" db:"created_at"`
	UpdatedBy  string     `json:"updated_by,omitempty" db:"updated_by"`
	UpdatedAt  *time.Time `json:"updated_at,omitempty" db:"updated_at"`
}

type BookListResponse struct {
	StatusCode int    `json:"error_code"`
	Message    string `json:"error_description"`
	BookList   []Book `json:"data,omitempty"`
}

type BookDeleteResponse struct {
	StatusCode     int              `json:"error_code"`
	Message        string           `json:"error_description,omitempty"`
	BookReportList []BookReportList `json:"book_issued"`
	DeleteBook     Book             `json:"data"`
}

type BookCreatedResponse struct {
	StatusCode  int    `json:"error_code"`
	Message     string `json:"error_description,omitempty"`
	CreatedBook Book   `json:"data"`
}

type UpdateBookResponse struct {
	StatusCode int    `json:"error_code"`
	Message    string `json:"error_description,omitempty"`
	UpdateBook Book   `json:"data,omitempty"`
}
