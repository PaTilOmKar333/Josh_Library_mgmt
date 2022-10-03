//struct and methods

package repo

import (
	"database/sql"
	"errors"
	"fmt"
	"log"
	"project/app"
	"project/models"
	"strings"
	"time"

	"github.com/jmoiron/sqlx"
)

type BookRepoInterface interface {
	ListBooks() (bookLists []models.Book, err error)
	CreateBook(book models.Book, createdBy, updatedBy string) (createdBook models.Book, err error)
	// DeleteBook(bid int) (id int, err error)
	DeleteBook(bid int) (id int, bookReportLists []models.BookReportList, err error)
	UpdateBookWithID(bid int, book models.Book, updatedBy string) (updatedBook models.Book, err error)
}

type bookRepo struct {
	db *sqlx.DB
}

func InitBookRepo() BookRepoInterface {

	var br bookRepo
	br.db = app.GetDB()
	return &br
}

func (br *bookRepo) ListBooks() (bookLists []models.Book, err error) {
	sqlStatement1 := `select * from books`

	err = br.db.Select(&bookLists, sqlStatement1)
	if err != nil {
		log.Println(err)
		err = errors.New("sorry for inconvenience, there is error in fetching list of books. we are working on this")
		return
	}

	return
}

func (br *bookRepo) CreateBook(book models.Book, createdBy, updatedBy string) (createdBook models.Book, err error) {

	sqlStatement := `INSERT INTO books( book_name, Status, price, category, created_by, created_at, updated_by, updated_at) VALUES ($1, $2, $3, $4, $5, $6, $7, $8) RETURNING book_id`
	//var id int
	err = br.db.Get(&book, sqlStatement, book.BookName, "available", book.Price, book.Category, createdBy, time.Now().Format("01-02-2006"), updatedBy, time.Now().Format("01-02-2006"))
	if err != nil {
		if err != nil {

			errorstring := err.Error()

			if strings.Contains(errorstring, "constraint_book_name") {
				fmt.Println("errorstring:", errorstring)
				err = errors.New("book is already exist.please enter new book")
				return
			} else {
				log.Println(err)
				err = errors.New("sorry for inconvenience, there is error in creating new of book. we are working on this")
				return
			}

		}
	}
	id := book.BookID

	sqlStatement1 := `select * from books WHERE book_id=$1`
	err = br.db.Get(&createdBook, sqlStatement1, id)

	if err == sql.ErrNoRows {
		log.Println(err)
		err = errors.New("book with provided ID is not present in database")
		return
	} else if err != nil {
		log.Println(err)
		err = errors.New("sorry for inconvenience, there is error in creating book. we are working on this")
		return
	}
	//fmt.Printf("inserted single record %v", id)

	return
}

func (br *bookRepo) DeleteBook(bid int) (id int, bookReportLists []models.BookReportList, err error) {

	var getbook models.Book
	var selectbook models.BookReport

	sqlStatement1 := `select * from books WHERE book_id=$1`
	err = br.db.Get(&getbook, sqlStatement1, bid)
	id = bid

	if err == sql.ErrNoRows {
		log.Println(err)
		err = errors.New("book with provided ID is not present in database")
		return
	} else if err != nil {
		log.Println(err)
		err = errors.New("sorry for inconvenience, there is error in deleting book. we are working on this")
		return
	}

	queryToCheckBookReport := `SELECT * FROM book_report WHERE book_id=$1 and actual_retun_date is null`
	err = br.db.Get(&selectbook, queryToCheckBookReport, bid)

	if err == sql.ErrNoRows {
		queryToDeleteBookReport := `DELETE FROM book_report WHERE book_id=$1`
		_, err = br.db.Exec(queryToDeleteBookReport, bid)
		if err != nil {
			log.Println(err)
			err = errors.New("sorry for inconvenience, there is error in deleteing book. we are working on this")
			return
		}

		queryToDeleteBook := `DELETE FROM books WHERE book_id=$1`
		_, err = br.db.Exec(queryToDeleteBook, bid)
		if err != nil {
			log.Println(err)
			err = errors.New("sorry for inconvenience, there is error in deleteing book. we are working on this")
			return
		}
		err = errors.New("book deleted successfully")
		return

	} else if err != nil {
		log.Println(err)
		err = errors.New("sorry for inconvenience, there is error in deleteing book. we are working on this")
		return
	} else {
		var selectbook []models.BookReport
		queryToCheckBookReport := `SELECT * FROM book_report WHERE book_id=$1 and actual_retun_date is null`
		err = br.db.Select(&selectbook, queryToCheckBookReport, bid)
		if err != nil {
			log.Println(err)
			err = errors.New("sorry for inconvenience, there is error in deleting book. we are working on this")
			return
		}

		for _, bookReport := range selectbook {
			var book models.Book
			var user models.User

			quertToGetBook := `select * from books where book_id=$1`
			err = br.db.Get(&book, quertToGetBook, bookReport.BookID)
			if err != nil {
				log.Println(err)
				err = errors.New("sorry for inconvenience, there is error in deleting book. we are working on this")
				return
			}

			quertToGetUser := `select * from users where user_id=$1`
			err = br.db.Get(&user, quertToGetUser, bookReport.UserID)
			if err != nil {
				log.Println(err)
				err = errors.New("sorry for inconvenience, there is error in deleting book. we are working on this")
				return
			}

			bookReportList := models.BookReportToBookReportList(bookReport, book, user)
			bookReportLists = append(bookReportLists, bookReportList)

		}

		err = errors.New("book is already assign to user. to delete book,user need to return it")
		return
	}

}

func (br *bookRepo) UpdateBookWithID(bid int, book models.Book, updatedBy string) (updatedBook models.Book, err error) {
	var oldbook models.Book

	sqlStatement1 := `select * FROM books WHERE book_id=$1`
	err = br.db.Get(&oldbook, sqlStatement1, bid)
	if err == sql.ErrNoRows {
		log.Println(err)
		err = errors.New("book with Provided ID is not present in database")
		return
	} else if err != nil {
		log.Println(err)
		err = errors.New("sorry for inconvenience, there is error in updating book. we are working on this")
		return
	}

	if book.BookName == "" {
		book.BookName = oldbook.BookName

	} else if book.Price == 0 {
		book.Price = oldbook.Price

	} else if book.Category == "" {
		book.Category = oldbook.Category
	}

	sqlStatement2 := `UPDATE books SET book_name=$2, price=$3, category=$4, updated_by=$5, updated_at=$6 WHERE book_id=$1 RETURNING book_id`

	err = br.db.Get(&book, sqlStatement2, bid, book.BookName, book.Price, book.Category, updatedBy, time.Now().Format("01-02-2006"))
	if err != nil {

		log.Println(err)
		err = errors.New("sorry for inconvenience, there is error in updating book. we are working on this")
		return
	}

	id := book.BookID
	var updateBook models.Book
	sqlStatement3 := `select * FROM books where book_id=$1 `
	err = br.db.Get(&updateBook, sqlStatement3, id)
	if err != nil {
		log.Println(err)
		err = errors.New("sorry for inconvenience, there is error in updating book. we are working on this")
		return
	}

	return
}
