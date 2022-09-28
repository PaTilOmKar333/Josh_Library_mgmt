package models

import (
	"time"
)

type BookReport struct {
	BookReportID     int        `json:"book_report_id" db:"bkreport_id"`
	BookID           int        `json:"book_id" db:"book_id"`
	UserID           int        `json:"user_id" db:"user_id"`
	IssueDate        time.Time  `json:"issue_date" db:"issue_date"`
	ReturnDate       time.Time  `json:"return_date" db:"return_date"`
	ActualReturnDate *time.Time `json:"actual_return_date,omitempty" db:"actual_retun_date"`
	IssuedBy         string     `json:"issued_by,omitempty" db:"issued_by"`
}

type BookReportResponse struct {
	ID         int    `json:"book_report_id,omitempty"`
	Message    string `json:"message"`
	StatusCode int    `json:"status_code"`
}

type IssueBook struct {
	BookReportID int        `json:"Id,omitempty"`
	BookName     string     `json:"name,omitempty"`
	Price        int        `json:"prize,omitempty"`
	Category     string     `json:"category,omitempty"`
	Status       string     `json:"status,omitempty"`
	IssuedTo     string     `json:"issued_to,omitempty"`
	IssuedBy     string     `json:"issued_by,omitempty"`
	ReturnDate   *time.Time `json:"return_date,omitempty"`
}

type IssueBookResponse struct {
	StatusCode int       `json:"error_code"`
	Message    string    `json:"error_description,omitempty"`
	IssueBook  IssueBook `json:"data"`
}

type BookReportList struct {
	BookReportID     int       `json:"Id,omitempty"`
	BookName         string    `json:"name,omitempty"`
	Price            int       `json:"prize,omitempty"`
	Category         string    `json:"category,omitempty"`
	BookStatus       string    `json:"status,omitempty"`
	UserName         string    `json:"issued_to,omitempty"`
	IssuedBy         string    `json:"issued_by,omitempty"`
	ReturnDate       time.Time `json:"return_date,omitempty"`
	IssueDate        time.Time `json:"issue_date,omitempty"`
	ActualReturnDate time.Time `json:"actual_return_date,omitempty"`
}

type BookReportListResponse struct {
	StatusCode     int              `json:"error_code"`
	Message        string           `json:"error_description,omitempty"`
	BookReportList []BookReportList `json:"data"`
}

type ReturnBookResponse struct {
	StatusCode     int            `json:"error_code"`
	Message        string         `json:"error_description,omitempty"`
	BookReportList BookReportList `json:"return_book,omitempty"`
}
