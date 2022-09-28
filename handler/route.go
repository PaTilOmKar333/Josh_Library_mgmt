package handler

import (
	"github.com/gorilla/mux"
)

func Router() *mux.Router {
	//r := mux.NewRouter()

	//r.HandleFunc("/users", LoginHandler(depnd.Userservice)).Methods("GET")
	router := mux.NewRouter()

	router.HandleFunc("/login", LoginHandler(depnd.Userservice)).Methods("POST")
	router.HandleFunc("/users", IsAuthorized([]string{"superadmin", "admin"}, depnd.AuthTokenService, AllUsersHandler(depnd.Userservice))).Methods("GET")
	router.HandleFunc("/user", IsAuthorized([]string{"superadmin", "admin"}, depnd.AuthTokenService, CreateUsersHandler(depnd.Userservice))).Methods("POST")
	//router.HandleFunc("/user", IsAuthorized([]string{"superadmin", "admin", "user"}, depnd.AuthTokenService, GetUsersByEmailOrIDHandler(depnd.Userservice))).Methods("GET")
	router.HandleFunc("/user/{user_id}", IsAuthorized([]string{"superadmin", "admin", "user"}, depnd.AuthTokenService, GetUsersByIDHandler(depnd.Userservice))).Methods("GET")
	//router.HandleFunc("/user/{email_id}", IsAuthorized([]string{"superadmin", "admin", "user"}, depnd.AuthTokenService, UpdateUserHandler(depnd.Userservice))).Methods("PUT")
	router.HandleFunc("/user/{user_id}", IsAuthorized([]string{"superadmin", "admin", "user"}, depnd.AuthTokenService, UpdateUserWithIDHandler(depnd.Userservice))).Methods("PUT")
	router.HandleFunc("/user/{user_id}", IsAuthorized([]string{"superadmin", "admin"}, depnd.AuthTokenService, DeleteUserHandler(depnd.Userservice))).Methods("DELETE")
	router.HandleFunc("/book", IsAuthorized([]string{"superadmin", "admin", "user"}, depnd.AuthTokenService, AllBooksHandler(depnd.Bookservice))).Methods("GET")
	router.HandleFunc("/book", IsAuthorized([]string{"superadmin", "admin"}, depnd.AuthTokenService, CreateBooksHandler(depnd.Bookservice))).Methods("POST")
	router.HandleFunc("/book/{book_id}", IsAuthorized([]string{"superadmin", "admin"}, depnd.AuthTokenService, UpdateBookWithIDHandler(depnd.Bookservice))).Methods("PUT")
	router.HandleFunc("/book/{book_id}", IsAuthorized([]string{"superadmin", "admin"}, depnd.AuthTokenService, DeleteBookHandler(depnd.Bookservice))).Methods("DELETE")
	router.HandleFunc("/user/{user_id}/book/{book_id}", IsAuthorized([]string{"superadmin", "admin"}, depnd.AuthTokenService, IssueBookHandler(depnd.BookReportService))).Methods("POST")
	router.HandleFunc("/user/{user_id}/book", IsAuthorized([]string{"superadmin", "admin", "user"}, depnd.AuthTokenService, GetBookReportHandler(depnd.BookReportService))).Methods("GET")
	router.HandleFunc("/bookreport", IsAuthorized([]string{"superadmin", "admin"}, depnd.AuthTokenService, GetAllBookReportHandler(depnd.BookReportService))).Methods("GET")
	router.HandleFunc("/user/{user_id}/book/{book_id}/return", IsAuthorized([]string{"superadmin", "admin"}, depnd.AuthTokenService, ReturnBookHandler(depnd.BookReportService))).Methods("POST")
	return router
}

// router.HandleFunc("/login/", LoginHandler(depnd.Userservice)).Methods("GET")
// router.HandleFunc("/users/", IsAuthorized([]string{"admin"}, AllUsersHandler(depnd.Userservice))).Methods("GET")
// router.HandleFunc("/user/", CreateUsersHandler(depnd.Userservice)).Methods("POST")
// //	router.HandleFunc("/user/{user_id}", GetUsersByIDHandler(depnd.Userservice)).Methods("GET")
// router.HandleFunc("/user", GetUsersByEmailOrIDHandler(depnd.Userservice)).Methods("GET")
// //router.HandleFunc("/userr/{email_id}/", GetUsersByEmailHandler(depnd.Userservice)).Methods("GET")
// router.HandleFunc("/user/{email_id}/", UpdateUserHandler(depnd.Userservice)).Methods("PUT")
// router.HandleFunc("/user/{user_id}/", DeleteUserHandler(depnd.Userservice)).Methods("DELETE")
// router.HandleFunc("/book/", AllBooksHandler(depnd.Bookservice)).Methods("GET")
// router.HandleFunc("/book/", CreateBooksHandler(depnd.Bookservice)).Methods("POST")
// router.HandleFunc("/book/{book_id}/", DeleteBookHandler(depnd.Bookservice)).Methods("DELETE")
// router.HandleFunc("/user/{user_id}/book/{book_id}/", IssueBookHandler(depnd.BookReportService)).Methods("POST")
// router.HandleFunc("/user/{user_id}/book/", GetBookReportHandler(depnd.BookReportService)).Methods("GET")
// router.HandleFunc("/user/{user_id}/book/{book_id}/return", ReturnBookHandler(depnd.BookReportService)).Methods("POST")
