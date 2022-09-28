package handler

import (
	"encoding/json"
	"errors"
	"log"
	"net/http"
	"project/models"
	"project/service"
	"strconv"

	"github.com/gorilla/mux"
)

func AllBooksHandler(bookService service.BookServiceInterface) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var listBooks models.BookListResponse

		books, err := bookService.ListBooks()
		if err != nil {
			listBooks.Message = err.Error()
			listBooks.StatusCode = http.StatusInternalServerError
			res, _ := json.Marshal(listBooks)
			w.WriteHeader(http.StatusInternalServerError)
			w.Write(res)

			return
		}
		listBooks.BookList = books
		listBooks.StatusCode = http.StatusOK

		w.WriteHeader(http.StatusOK)

		res, _ := json.Marshal(listBooks)
		w.Write(res)

	}
}

func CreateBooksHandler(bookService service.BookServiceInterface) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var book models.Book
		var createBookResponse models.BookCreatedResponse

		err := json.NewDecoder(r.Body).Decode(&book)

		if err != nil {
			createBookResponse.Message = "unable to decode the request body."
			createBookResponse.StatusCode = http.StatusBadRequest
			res, _ := json.Marshal(createBookResponse)
			w.WriteHeader(http.StatusBadRequest)
			w.Write(res)
			return
		}

		createdbook, err := bookService.CreateBook(r.Context(), book)
		if err != nil {
			createBookResponse.Message = err.Error()
			createBookResponse.StatusCode = http.StatusInternalServerError
			res, _ := json.Marshal(createBookResponse)
			w.WriteHeader(http.StatusInternalServerError)
			w.Write(res)
			return
		}
		createBookResponse.CreatedBook = createdbook
		createBookResponse.StatusCode = http.StatusOK

		w.WriteHeader(http.StatusOK)

		res, _ := json.Marshal(createBookResponse)
		w.Write(res)
	}
}

func DeleteBookHandler(bookService service.BookServiceInterface) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var DeleteBookResponse models.BookDeleteResponse
		params := mux.Vars(r)

		// convert the id type from string to int
		bid, err := strconv.Atoi(params["book_id"])

		if err != nil {
			log.Print(err)
			DeleteBookResponse.Message = "Unable to convert the string bookid into int bookid"
			DeleteBookResponse.StatusCode = http.StatusBadRequest
			res, _ := json.Marshal(DeleteBookResponse)
			w.WriteHeader(http.StatusBadRequest)
			w.Write(res)
			return
		}
		id, bookReportLists, err := bookService.DeleteBook(bid)

		if err != errors.New("book is already assign to user. to delete book,user need to return it") {

			DeleteBookResponse.BookReportList = bookReportLists
			DeleteBookResponse.Message = err.Error()
			DeleteBookResponse.StatusCode = http.StatusInternalServerError
			res, _ := json.Marshal(DeleteBookResponse)
			w.WriteHeader(http.StatusInternalServerError)
			w.Write(res)

			return
		}
		DeleteBookResponse.DeleteBook.BookID = id
		DeleteBookResponse.StatusCode = http.StatusOK

		w.WriteHeader(http.StatusOK)

		res, _ := json.Marshal(DeleteBookResponse)
		w.Write(res)
	}
}

func UpdateBookWithIDHandler(bookService service.BookServiceInterface) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

		var updateBookResponse models.UpdateBookResponse
		var book models.Book

		err := json.NewDecoder(r.Body).Decode(&book)

		if err != nil {
			updateBookResponse.Message = "unable to decode the request body."
			updateBookResponse.StatusCode = http.StatusBadRequest
			w.WriteHeader(http.StatusBadRequest)
			res, _ := json.Marshal(updateBookResponse)
			w.Write(res)
			return
		}
		params := mux.Vars(r)
		bid, err := strconv.Atoi(params["book_id"])

		if err != nil {
			updateBookResponse.Message = "unable to convert userid in int."
			updateBookResponse.StatusCode = http.StatusBadRequest
			w.WriteHeader(http.StatusBadRequest)
			res, _ := json.Marshal(updateBookResponse)
			w.Write(res)
			return
		}
		updatedBook, err := bookService.UpdateBookWithID(r.Context(), bid, book)

		if err != nil {
			//fmt.Sprintln("error....")
			updateBookResponse.Message = err.Error()
			updateBookResponse.StatusCode = http.StatusInternalServerError
			res, _ := json.Marshal(updateBookResponse)
			w.WriteHeader(http.StatusInternalServerError)
			w.Write(res)

			return
		}
		updateBookResponse.UpdateBook = updatedBook
		updateBookResponse.StatusCode = http.StatusOK

		w.WriteHeader(http.StatusOK)

		res, _ := json.Marshal(updateBookResponse)
		w.Write(res)

	}
}
