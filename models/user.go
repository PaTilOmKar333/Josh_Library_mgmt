package models

import (
	"time"

	"github.com/golang-jwt/jwt"
)

// User struct declaration
type User struct {
	ID        int        `json:"id,omitempty" db:"id"`
	Name      string     `json:"name,omitempty" db:"name"`
	Email     string     `json:"email,omitempty" db:"email"`
	Password  string     `json:"password,omitempty" db:"password"`
	RoleType  string     `json:"role_type,omitempty" db:"role_type"`
	CreatedBy string     `json:"created_by,omitempty" db:"created_by"`
	CreatedAT *time.Time `json:"created_at,omitempty" db:"created_at"`
	UpdatedBy string     `json:"updated_by,omitempty" db:"updated_by"`
	UpdatedAt *time.Time `json:"updated_at,omitempty" db:"updated_at"`
}

type UserAuth struct {
	ID       int    `json:"id ,omitempty" db:"id"`
	Name     string `json:"name,omitempty" db:"name"`
	Email    string `json:"email,omitempty" db:"email"`
	Password string `json:"password,omitempty" db:"password"`
	RoleType string `json:"role_type,omitempty" db:"role_type"`
}

type Authentication struct {
	ID       int    `json:"id ,omitempty"`
	Email    string `json:"email"`
	Password string `json:"password"`
	RoleType string `json:"role_type,omitempty"`
}

type Claims struct {
	ID       int    `json:"id"`
	RoleType string `json:"role_type"`
	Email    string `json:"email"`
	jwt.StandardClaims
}

type Token struct {
	Token string `json:"login_token,omitempty"`
}

type TokenResponse struct {
	StatusCode int    `json:"error_code"`
	Message    string `json:"error_description"`
	Token      Token  `json:"data,omitempty"`
}

type CreateUserResponse struct {
	StatusCode int    `json:"error_code"`
	Message    string `json:"error_description"`
	CreateUser User   `json:"data,omitempty"`
}

type UserListResponse struct {
	StatusCode int    `json:"error_code"`
	Message    string `json:"error_description"`
	UserList   []User `json:"data,omitempty"`
}

type UpdateUserResponse struct {
	StatusCode  int    `json:"error_code"`
	Message     string `json:"error_description"`
	UpdatedUser User   `json:"data,omitempty"`
}

type GetUserResponse struct {
	StatusCode int    `json:"error_code"`
	Message    string `json:"error_description"`
	GotUser    User   `json:"data,omitempty"`
}

type AuthenticationResponse struct {
	StatusCode int    `json:"status_code"`
	Message    string `json:"message"`
}

type UserDeleteResponse struct {
	StatusCode     int              `json:"error_code"`
	Message        string           `json:"error_description"`
	BookReportList []BookReportList `json:"book_issued"`
	DeletedUser    User             `json:"data,omitempty"`
}
