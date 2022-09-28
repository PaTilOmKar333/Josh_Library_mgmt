package handler

import (
	"encoding/json"
	"net/http"
	"project/models"
	"project/service"
	"strconv"

	"github.com/gorilla/mux"
)

func IssueBookHandler(bookReportService service.BookReportServiceInterface) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		//	var issueBookResponse models.BookReportResponse
		var issueBookResponse models.IssueBookResponse
		params := mux.Vars(r)

		// convert the id type from string to int
		uid, err := strconv.Atoi(params["user_id"])
		if err != nil {
			issueBookResponse.Message = "Unable to convert the string userid into int userid"
			issueBookResponse.StatusCode = http.StatusBadRequest
			res, _ := json.Marshal(issueBookResponse)
			w.WriteHeader(http.StatusBadRequest)
			w.Write(res)

			return
		}
		bid, err := strconv.Atoi(params["book_id"])
		if err != nil {
			issueBookResponse.Message = "Unable to convert the string userid into int userid"
			issueBookResponse.StatusCode = http.StatusBadRequest
			res, _ := json.Marshal(issueBookResponse)
			w.WriteHeader(http.StatusBadRequest)
			w.Write(res)
			return
		}
		issueBook, err := bookReportService.IssueBook(r.Context(), uid, bid)

		if err != nil {
			issueBookResponse.Message = err.Error()
			issueBookResponse.StatusCode = http.StatusInternalServerError
			res, _ := json.Marshal(issueBookResponse)
			w.WriteHeader(http.StatusInternalServerError)
			w.Write(res)

			return
		}
		issueBookResponse.IssueBook = issueBook
		issueBookResponse.StatusCode = http.StatusOK

		w.WriteHeader(http.StatusOK)

		res, _ := json.Marshal(issueBookResponse)
		w.Write(res)
	}
}

func GetAllBookReportHandler(bookReportService service.BookReportServiceInterface) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var bookReportListResponse models.BookReportListResponse
		reports, err := bookReportService.GetAllBookReport()

		if err != nil {
			bookReportListResponse.Message = err.Error()
			bookReportListResponse.StatusCode = http.StatusInternalServerError
			res, _ := json.Marshal(bookReportListResponse)
			w.WriteHeader(http.StatusInternalServerError)
			w.Write(res)

			return
		}
		bookReportListResponse.BookReportList = reports
		bookReportListResponse.StatusCode = http.StatusOK

		w.WriteHeader(http.StatusOK)

		res, _ := json.Marshal(bookReportListResponse)
		w.Write(res)
	}
}

func GetBookReportHandler(bookReportService service.BookReportServiceInterface) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var bookReportListResponse models.BookReportListResponse
		params := mux.Vars(r)

		// convert the id type from string to int
		uid, err := strconv.Atoi(params["user_id"])
		if err != nil {
			bookReportListResponse.Message = "Unable to convert the string userid into int userid"
			bookReportListResponse.StatusCode = http.StatusBadRequest
			res, _ := json.Marshal(bookReportListResponse)
			w.WriteHeader(http.StatusBadRequest)
			w.Write(res)
			return
		}

		reports, err := bookReportService.GetBookReport(r.Context(), uid)

		if err != nil {
			bookReportListResponse.Message = err.Error()
			bookReportListResponse.StatusCode = http.StatusInternalServerError
			res, _ := json.Marshal(bookReportListResponse)
			w.WriteHeader(http.StatusInternalServerError)
			w.Write(res)

			return
		}
		bookReportListResponse.BookReportList = reports
		bookReportListResponse.StatusCode = http.StatusOK

		w.WriteHeader(http.StatusOK)

		res, _ := json.Marshal(bookReportListResponse)
		w.Write(res)
	}
}

func ReturnBookHandler(bookReportService service.BookReportServiceInterface) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var returnBookResponse models.ReturnBookResponse
		params := mux.Vars(r)

		// convert the id type from string to int
		uid, err := strconv.Atoi(params["user_id"])
		if err != nil {
			returnBookResponse.Message = "Unable to convert the string userid into int userid"
			returnBookResponse.StatusCode = http.StatusBadRequest
			res, _ := json.Marshal(returnBookResponse)
			w.WriteHeader(http.StatusBadRequest)
			w.Write(res)
			return
		}
		bid, err := strconv.Atoi(params["book_id"])
		if err != nil {
			returnBookResponse.Message = "Unable to convert the string userid into int userid"
			returnBookResponse.StatusCode = http.StatusBadRequest
			res, _ := json.Marshal(returnBookResponse)
			w.WriteHeader(http.StatusBadRequest)
			w.Write(res)
			return
		}
		book, err := bookReportService.ReturnBook(uid, bid)

		if err != nil {
			returnBookResponse.Message = err.Error()
			returnBookResponse.StatusCode = http.StatusInternalServerError
			res, _ := json.Marshal(returnBookResponse)
			w.WriteHeader(http.StatusInternalServerError)
			w.Write(res)
			return
		}
		returnBookResponse.BookReportList = book
		returnBookResponse.StatusCode = http.StatusOK

		w.WriteHeader(http.StatusOK)

		res, _ := json.Marshal(returnBookResponse)
		w.Write(res)
	}
}
