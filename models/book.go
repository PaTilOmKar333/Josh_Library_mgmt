package models

import "time"

// User struct declaration
type BookStatus struct {
	StatusID int    `json:"status_id,omitempty" db:"status_id"`
	Status   string `json:"status" db:"status"`
}

type Book struct {
	BookID          int        `json:"id" db:"book_id"`
	BookName        string     `json:"name" db:"book_name"`
	AuthorName      string     `json:"author_name" db:"author_name"`
	AvailableCopies int        `json:"no_of_copies_available" db:"available_book_copies"`
	BookStatusID    int        `json:"status_id" db:"status_id"`
	Price           int        `json:"price" db:"price"`
	Category        string     `json:"category" db:"category"`
	CreatedBy       string     `json:"created_by,omitempty" db:"created_by"`
	CreatedAT       *time.Time `json:"created_at,omitempty" db:"created_at"`
	UpdatedBy       string     `json:"updated_by,omitempty" db:"updated_by"`
	UpdatedAt       *time.Time `json:"updated_at,omitempty" db:"updated_at"`
}

type CreateBook struct {
	BookID          int        `json:"book_id" db:"book_id"`
	BookName        string     `json:"book_name" db:"book_name"`
	AuthorName      string     `json:"author_name" db:"author_name"`
	AvailableCopies int        `json:"no_of_copies_available" db:"available_book_copies"`
	BookStatus      string     `json:"status" db:"status_id"`
	Price           int        `json:"price" db:"price"`
	Category        string     `json:"category" db:"category"`
	CreatedBy       string     `json:"created_by,omitempty" db:"created_by"`
	CreatedAT       *time.Time `json:"created_at,omitempty" db:"created_at"`
	UpdatedBy       string     `json:"updated_by,omitempty" db:"updated_by"`
	UpdatedAt       *time.Time `json:"updated_at,omitempty" db:"updated_at"`
}

type BookList struct {
	BookID          int    `json:"book_id"`
	BookName        string `json:"book_name"`
	AuthorName      string `json:"author_name"`
	AvailableCopies int    `json:"no_of_copies_available"`
	Status          string `json:"status"`
}

type BookListResponse struct {
	StatusCode int        `json:"error_code"`
	Message    string     `json:"error_description"`
	BookList   []BookList `json:"data,omitempty"`
}

type BookResponse struct {
	ID         int    `json:"book_id,omitempty"`
	Message    string `json:"message"`
	StatusCode int    `json:"status_code"`
}

type DeleteBook struct {
	BookID   int    `json:"id,omitempty"`
	BookName string `json:"name,omitempty"`
}

type BookDeleteResponse struct {
	StatusCode     int              `json:"error_code"`
	Message        string           `json:"error_description,omitempty"`
	BookReportList []BookReportList `json:"book_issued"`
	DeleteBook     DeleteBook       `json:"data"`
}

type BookCreatedResponse struct {
	StatusCode  int        `json:"error_code"`
	Message     string     `json:"error_description,omitempty"`
	CreatedBook CreateBook `json:"data"`
}

type UpdateBook struct {
	BookID          int        `json:"id" db:"book_id"`
	BookName        string     `json:"name" db:"book_name"`
	AuthorName      string     `json:"author_name" db:"author_name"`
	AvailableCopies int        `json:"no_of_copies_available" db:"available_book_copies"`
	BookStatus      string     `json:"status"`
	Price           int        `json:"price" db:"price"`
	Category        string     `json:"category" db:"category"`
	CreatedBy       string     `json:"created_by,omitempty" db:"created_by"`
	CreatedAT       *time.Time `json:"created_at,omitempty" db:"created_at"`
	UpdatedBy       string     `json:"updated_by,omitempty" db:"updated_by"`
	UpdatedAt       *time.Time `json:"updated_at,omitempty" db:"updated_at"`
}

type UpdateBookResponse struct {
	StatusCode int        `json:"error_code"`
	Message    string     `json:"error_description,omitempty"`
	UpdateBook UpdateBook `json:"data,omitempty"`
}
