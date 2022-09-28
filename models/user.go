package models

import (
	"time"

	"github.com/golang-jwt/jwt"
)

// User struct declaration
type User struct {
	User_ID   int        `json:"u_id ,omitempty" db:"user_id"`
	FirstName string     `json:"u_firstname,omitempty" db:"first_name"`
	LastName  string     `json:"u_lastname,omitempty" db:"last_name"`
	Age       int        `json:"age,omitempty" db:"age"`
	Email     string     `json:"email,omitempty" db:"email"`
	Password  string     `json:"password,omitempty" db:"password"`
	Address   string     `json:"address,omitempty" db:"address"`
	Role_ID   int        `json:"role_id,omitempty" db:"role_id"`
	CreatedBy string     `json:"created_by,omitempty" db:"created_by"`
	CreatedAT *time.Time `json:"created_at,omitempty" db:"created_at"`
	UpdatedBy string     `json:"updated_by,omitempty" db:"updated_by"`
	UpdatedAt *time.Time `json:"updated_at,omitempty" db:"updated_at"`
}

type UserAuth struct {
	User_ID   int    `json:"u_id"`
	FirstName string `json:"u_firstname"`
	LastName  string `json:"u_lastname"`
	Age       int    `json:"age"`
	Email     string `json:"email"`
	Password  string `json:"password,omitempty"`
	Address   string `json:"address"`
	Role      string `json:"role"`
}

// type UpdateUser struct {
// 	OldFirstName string `json:"old_firstname,omitempty"`
// 	OldLastName  string `json:"old_lastname,omitempty"`
// 	OldPassword  string `json:"old_password,omitempty"`
// 	NewFirstName string `json:"new_firstname,omitempty"`
// 	NewLastName  string `json:"new_lastname,omitempty"`
// 	NewPassword  string `json:"new_password,omitempty"`
// }

type UpdateUser struct {
	User_ID   int        `json:"id ,omitempty" db:"user_id"`
	FirstName string     `json:"firstname,omitempty" db:"first_name"`
	LastName  string     `json:"lastname,omitempty" db:"last_name"`
	Age       int        `json:"age,omitempty" db:"age"`
	Email     string     `json:"email,omitempty" db:"email"`
	Password  string     `json:"password,omitempty" db:"password"`
	Address   string     `json:"address,omitempty" db:"address"`
	Role      string     `json:"role_type,omitempty" db:"role_id"`
	CreatedBy string     `json:"created_by,omitempty" db:"created_by"`
	CreatedAT *time.Time `json:"created_at,omitempty" db:"created_at"`
	UpdatedBy string     `json:"updated_by,omitempty" db:"updated_by"`
	UpdatedAt *time.Time `json:"updated_at,omitempty" db:"updated_at"`
}

type CreateUser struct {
	User_ID   int        `json:"id ,omitempty" db:"user_id"`
	FirstName string     `json:"firstname,omitempty" db:"first_name"`
	LastName  string     `json:"lastname,omitempty" db:"last_name"`
	Age       int        `json:"age,omitempty" db:"age"`
	Email     string     `json:"email,omitempty" db:"email"`
	Password  string     `json:"password,omitempty" db:"password"`
	Address   string     `json:"address,omitempty" db:"address"`
	Role      string     `json:"role_type,omitempty" db:"role_id"`
	CreatedBy string     `json:"createdb_y,omitempty" db:"created_by"`
	CreatedAT *time.Time `json:"created_at,omitempty" db:"created_at"`
	UpdatedBy string     `json:"updated_by,omitempty" db:"updated_by"`
	UpdatedAt *time.Time `json:"updated_at,omitempty" db:"updated_at"`
}

type UserRole struct {
	RoleID int    `json:"role_id" db:"role_id"`
	Role   string `json:"role_type" db:"role_name"`
}

type UserList struct {
	User_ID   int        `json:"id"`
	FirstName string     `json:"firstname"`
	LastName  string     `json:"lastname"`
	Age       int        `json:"age"`
	Email     string     `json:"email"`
	Address   string     `json:"address"`
	Role      string     `json:"role_type"`
	CreatedBy string     `json:"createdb_y,omitempty"`
	CreatedAT *time.Time `json:"created_at,omitempty"`
	UpdatedBy string     `json:"updated_by,omitempty"`
	UpdatedAt *time.Time `json:"updated_at,omitempty"`
}

type Authentication struct {
	UserID   int    `json:"id"`
	Email    string `json:"email"`
	Password string `json:"password"`
	Role     string `json:"role_type"`
}

type Claims struct {
	UserID int    `json:"id"`
	Role   string `json:"role_type"`
	Email  string `json:"email"`
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
	StatusCode int        `json:"error_code"`
	Message    string     `json:"error_description"`
	CreateUser CreateUser `json:"data,omitempty"`
}

type UserListResponse struct {
	StatusCode int        `json:"error_code"`
	Message    string     `json:"error_description"`
	UserList   []UserList `json:"data,omitempty"`
}

type UpdateUserResponse struct {
	StatusCode  int        `json:"error_code"`
	Message     string     `json:"error_description"`
	UpdatedUser UpdateUser `json:"data,omitempty"`
}

type GetUser struct {
	User_ID   int        `json:"id ,omitempty" db:"user_id"`
	FirstName string     `json:"firstname,omitempty" db:"first_name"`
	LastName  string     `json:"lastname,omitempty" db:"last_name"`
	Age       int        `json:"age,omitempty" db:"age"`
	Email     string     `json:"email,omitempty" db:"email"`
	Password  string     `json:"password,omitempty" db:"password"`
	Address   string     `json:"address,omitempty" db:"address"`
	Role      string     `json:"role_type,omitempty" db:"role_id"`
	CreatedBy string     `json:"createdb_y,omitempty" db:"created_by"`
	CreatedAT *time.Time `json:"created_at,omitempty" db:"created_at"`
	UpdatedBy string     `json:"updated_by,omitempty" db:"updated_by"`
	UpdatedAt *time.Time `json:"updated_at,omitempty" db:"updated_at"`
}

type GetUserResponse struct {
	StatusCode int     `json:"error_code"`
	Message    string  `json:"error_description"`
	GotUser    GetUser `json:"data,omitempty"`
}

type AuthenticationResponse struct {
	StatusCode int    `json:"status_code"`
	Message    string `json:"message"`
}

type DeleteUser struct {
	UserID    int    `json:"id,omitempty" db:"user_id"`
	FirstName string `json:"firstname,omitempty" db:"first_name"`
	Email     string `json:"email,omitempty" db:"email"`
}

type UserDeleteResponse struct {
	StatusCode     int              `json:"error_code"`
	Message        string           `json:"error_description"`
	BookReportList []BookReportList `json:"book_issued"`
	DeletedUser    DeleteUser       `json:"data,omitempty"`
}
