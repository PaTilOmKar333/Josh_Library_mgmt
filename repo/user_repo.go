//struct and methods

package repo

import (
	"database/sql"
	"errors"
	"fmt"
	"log"
	"math/rand"
	"project/app"
	"project/models"
	"strings"
	"time"

	"github.com/jmoiron/sqlx"
)

type UserRepoInterface interface {
	//user methods
	Login(authdetails models.Authentication) (err error)
	ListUser() (usersList []models.User, err error)
	CreateUser(cratedByEmail, updatedByEmail string, user models.User) (createduser models.User, err error)
	GetUserByID(uid int) (user models.User, err error)
	GetUser(param interface{}) (user models.User, err error)
	//	UpdateUser(email string, user models.User) (updateUser models.UpdateUser, err error)
	GetUserByEmail(email string) (userdetails models.UserAuth, err error)
	//DeleteUser(uid int) (deleteuser models.User, err error)
	DeleteUser(uid int) (selectuser models.User, bookReportLists []models.BookReportList, err error)
	UpdateUserWithID(uid int, user models.User, updatedBy string) (updateUser models.User, err error)
}

type userRepo struct {
	db *sqlx.DB
}

func InitUserRepo() UserRepoInterface {
	//var err error
	var ur userRepo
	ur.db = app.GetDB()
	// return object
	return &ur
}

func (ur *userRepo) Login(authdetails models.Authentication) (err error) {
	return
}

func (ur *userRepo) ListUser() (usersList []models.User, err error) {
	sqlStatement := `SELECT * FROM users`

	err = ur.db.Select(&usersList, sqlStatement)
	if err != nil {
		log.Println(err)
		err = errors.New("sorry for inconvenience, there is error in fetching list of users. we are working on this")
		return
	}

	return
}

func (ur *userRepo) CreateUser(cratedByEmail, updatedByEmail string, user models.User) (createduser models.User, err error) {

	password := func() int {
		return 10000 + rand.Intn(99999-10000)
	}()

	sqlStatement := `INSERT INTO users(name, email, password, role_type, created_by, created_at, updated_by, updated_at) VALUES ($1, $2, $3, $4, $5, $6, $7, $8) RETURNING id`
	err = ur.db.Get(&user, sqlStatement, user.Name, user.Email, password, user.RoleType, cratedByEmail, time.Now().Format("01-02-2006"), updatedByEmail, time.Now().Format("01-02-2006"))
	if err != nil {
		errorstring := err.Error()

		if strings.Contains(errorstring, "constraint_email_unique") {
			fmt.Println("errorstring:", errorstring)
			err = errors.New("user is already exists with same emailid. please use another email to create user")
			return
		} else {
			log.Println(err)
			err = errors.New("sorry for inconvenience, there is error in creating user. we are working on this")
			return
		}

	}
	user_id := user.ID
	sqlStatement1 := `select name, email, role_type, created_by, created_at, updated_by, updated_at FROM users WHERE id=$1`
	err = ur.db.Get(&createduser, sqlStatement1, user_id)
	if err != nil {
		log.Println(err)
		err = errors.New("sorry for inconvenience, there is error in creating user. we are working on this")
		return
	}
	return
}

func (ur *userRepo) GetUser(variable interface{}) (user models.User, err error) {
	//variable, err = strconv.Atoi(variable)

	switch variable.(type) {
	case int:
		sqlStatement := `select id, name, email FROM users where id=$1 `

		err = ur.db.Get(&user, sqlStatement, variable)
		if err == sql.ErrNoRows {
			log.Println(err)
			err = errors.New("no user in database with this id")
			return
		} else if err != nil {
			log.Println(err)
			err = errors.New("sorry for inconvenience, there is error in fetching user. we are working on this")
			return
		}
		return
	case string:

		sqlStatement := `select id, name, email FROM users where email=$1 `
		//var id int
		//err = ur.db.QueryRow(sqlStatement, user.User_ID, user.FirstName, user.LastName, user.Age, user.Email, user.Password, user.Address, user.Role_ID).Scan(&id)
		err = ur.db.Get(&user, sqlStatement, variable)
		if err == sql.ErrNoRows {
			log.Println(err)
			err = errors.New("no user in database with this email")
			return
		} else if err != nil {
			log.Println(err)
			err = errors.New("sorry for inconvenience, there is error in fetching user. we are working on this")
			return
		}
		return
	}

	return
}

// func (ur *userRepo) UpdateUser(email string, user models.User) (updateUser models.UpdateUser, err error) {
// 	var newuser, olduser models.User

// 	sqlStatement1 := `select * FROM users WHERE email=$1`
// 	err = ur.db.Get(&olduser, sqlStatement1, email)
// 	if err != nil {
// 		log.Println(err)
// 		err = errors.New("sorry for inconvenience, there is error in updating user. we are working on this")
// 		return
// 	}

// 	if user.Password == "" {
// 		user.Password = olduser.Password

// 	} else if user.FirstName == "" {
// 		user.FirstName = olduser.FirstName

// 	} else if user.LastName == "" {
// 		user.LastName = olduser.LastName
// 	}

// 	sqlStatement2 := `UPDATE users SET first_name=$2, last_name=$3, password=$4 WHERE email=$1 RETURNING user_id`

// 	err = ur.db.Get(&user, sqlStatement2, email, user.FirstName, user.LastName, user.Password)
// 	if err != nil {
// 		log.Println(err)
// 		err = errors.New("sorry for inconvenience, there is error in updating user. we are working on this")
// 		return
// 	}

// 	id := user.User_ID
// 	sqlStatement3 := `select * FROM users where user_id=$1 `
// 	err = ur.db.Get(&newuser, sqlStatement3, id)
// 	if err != nil {
// 		log.Println(err)
// 		err = errors.New("sorry for inconvenience, there is error in updating user. we are working on this")
// 		return
// 	}

// 	updateUser = models.UpdatedUserDetails(olduser, newuser)
// 	return
// }

func (ur *userRepo) UpdateUserWithID(uid int, user models.User, updatedBy string) (updateUser models.User, err error) {
	var olduser models.User

	sqlStatement1 := `select * FROM users WHERE id=$1`
	err = ur.db.Get(&olduser, sqlStatement1, uid)
	if err == sql.ErrNoRows {
		log.Println(err)
		err = errors.New("user with provided ID is not present in database")
		return
	} else if err != nil {
		log.Println(err)
		err = errors.New("sorry for inconvenience, there is error in updating user. we are working on this")
		return
	}

	if user.Password == "" {
		user.Password = olduser.Password

	} else if user.Name == "" {
		user.Name = olduser.Name

	}

	sqlStatement2 := `UPDATE users SET name=$2, password=$3, updated_by=$4, updated_at=$5 WHERE id=$1 RETURNING id`

	err = ur.db.Get(&user, sqlStatement2, uid, user.Name, user.Password, updatedBy, time.Now().Format("01-02-2006"))
	if err != nil {

		log.Println(err)
		err = errors.New("sorry for inconvenience, there is error in updating user. we are working on this")
		return
	}

	id := user.ID
	sqlStatement3 := `select * FROM users where id=$1 `
	err = ur.db.Get(&updateUser, sqlStatement3, id)
	if err != nil {
		log.Println(err)
		err = errors.New("sorry for inconvenience, there is error in updating user. we are working on this")
		return
	}

	return
}

func (ur *userRepo) DeleteUser(uid int) (selectuser models.User, bookReportLists []models.BookReportList, err error) {

	sqlStatement := `select id, name, email FROM users where id=$1 `
	err = ur.db.Get(&selectuser, sqlStatement, uid)

	if err == sql.ErrNoRows {
		log.Println(err)
		err = errors.New("user with provided ID is not present in database")
		return
	} else if err != nil {
		log.Println(err)
		err = errors.New("sorry for inconvenience, there is error in fetching user. we are working on this")
		return
	}

	var selectBookReport models.BookReport

	queryToCheckBookReport := `SELECT * FROM book_report WHERE user_id=$1 and return_date is null`
	err = ur.db.Get(&selectBookReport, queryToCheckBookReport, uid)
	if err == sql.ErrNoRows {
		queryToDeleteBookReport := `DELETE FROM book_report WHERE user_id=$1`
		_, err = ur.db.Exec(queryToDeleteBookReport, uid)
		if err != nil {
			log.Println(err)
			err = errors.New("sorry for inconvenience, there is error in deleteing User. we are working on this")
			return
		}
		queryToDeleteUser := `DELETE FROM users WHERE id=$1`
		_, err = ur.db.Exec(queryToDeleteUser, uid)
		if err != nil {
			log.Println(err)
			err = errors.New("sorry for inconvenience, there is error in deleteing User. we are working on this")
			return
		}
		return

	} else if err != nil {
		log.Println(err)
		err = errors.New("sorry for inconvenience, there is error in deleteing User. we are working on this")
		return
	} else {
		var selectbook []models.BookReport
		queryToCheckBookReport := `SELECT * FROM book_report WHERE user_id=$1 and return_date is null`
		err = ur.db.Select(&selectbook, queryToCheckBookReport, uid)
		if err != nil {
			log.Println(err)
			err = errors.New("sorry for inconvenience, there is error in deleting user. we are working on this")
			return
		}

		for _, bookReport := range selectbook {
			var book models.Book
			var user models.User

			quertToGetBook := `select * from books where book_id=$1`
			err = ur.db.Get(&book, quertToGetBook, bookReport.BookID)
			if err != nil {
				log.Println(err)
				err = errors.New("sorry for inconvenience, there is error in deleting user. we are working on this")
				return
			}

			quertToGetUser := `select * from users where id=$1`
			err = ur.db.Get(&user, quertToGetUser, bookReport.UserID)
			if err != nil {
				log.Println(err)
				err = errors.New("sorry for inconvenience, there is error in deleting user. we are working on this")
				return
			}

			bookReportList := models.BookReportToBookReportList(bookReport, book, user)
			bookReportLists = append(bookReportLists, bookReportList)

		}

		err = errors.New("book is already assign to user. to delete user,user need to return all books")
		return
	}

}

func (ur *userRepo) GetUserByEmail(email string) (userdetails models.UserAuth, err error) {

	//var user models.User

	sqlStatement := `select id, name, email, role_type, password FROM users where email=$1 `
	err = ur.db.Get(&userdetails, sqlStatement, email)

	if err == sql.ErrNoRows {
		log.Println(err)
		err = errors.New("user not found")
		return
	} else if err != nil {
		log.Println(err)
		err = errors.New("sorry for inconvenience, there is error in fetching user. we are working on this")
		return
	}

	return
}

func (ur *userRepo) GetUserByID(uid int) (user models.User, err error) {

	sqlStatement := `select id, name, email, role_type, created_by, created_at, updated_by, updated_at FROM users where id=$1 `
	err = ur.db.Get(&user, sqlStatement, uid)
	if err != nil {
		log.Println(err)
		err = errors.New("user with provided ID is not present in database")
		return
	}
	return
}
