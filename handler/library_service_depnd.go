package handler

import (
	"project/repo"
	"project/service"
)

type dependancies struct {
	Userservice       service.UserServiceInterface
	Bookservice       service.BookServiceInterface
	BookReportService service.BookReportServiceInterface
	AuthTokenService  service.AuthTokenInterface
}

var depnd dependancies

func InitDependancies() {
	userrepo := repo.InitUserRepo()
	authService := service.InitAuthService()
	userservice := service.InitUserService(userrepo, authService)
	depnd.Userservice = userservice
	depnd.AuthTokenService = authService

	bookrepo := repo.InitBookRepo()
	bookservice := service.InitBookService(bookrepo)
	depnd.Bookservice = bookservice

	bookreportrepo := repo.InitBookReportRepo()
	bookreportservice := service.InitBookReportService(bookreportrepo)
	depnd.BookReportService = bookreportservice
}
