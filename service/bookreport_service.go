//struct and methods

package service

import (
	"context"
	"errors"
	"project/models"
	"project/repo"
)

type BookReportServiceInterface interface {
	IssueBook(ctx context.Context, uid, bid int) (issueBook models.IssueBook, err error)
	GetBookReport(ctx context.Context, uid int) (bookReportLists []models.BookReportList, err error)
	GetAllBookReport() (bookReportLists []models.BookReportList, err error)
	ReturnBook(uid, bid int) (BookReport models.BookReportList, err error)
}

type bookreportService struct {
	repo repo.BookReportRepoInterface
}

func InitBookReportService(r repo.BookReportRepoInterface) BookReportServiceInterface {

	//initialies
	//repo.InitUserRepo()
	return &bookreportService{
		repo: r,
	}
}

func (brs *bookreportService) IssueBook(ctx context.Context, uid, bid int) (issueBook models.IssueBook, err error) {
	val, _ := ctx.Value("ClaimsToVerify").(*models.Claims)

	issuedBy := val.Email

	issueBook, err = brs.repo.IssueBook(uid, bid, issuedBy)
	if err != nil {
		return
	}
	return
}

func (brs *bookreportService) GetBookReport(ctx context.Context, uid int) (bookReportLists []models.BookReportList, err error) {

	val, _ := ctx.Value("ClaimsToVerify").(*models.Claims)
	if val.ID == uid || val.RoleType == "admin" || val.RoleType == "superadmin" {
		bookReportLists, err = brs.repo.GetBookReport(uid)
		if err != nil {
			return
		}
	} else {
		err = errors.New("unauthorized user")

	}

	return
}

func (brs *bookreportService) GetAllBookReport() (bookReportLists []models.BookReportList, err error) {
	bookReportLists, err = brs.repo.GetAllBookReport()
	if err != nil {
		return
	}

	return
}

func (brs *bookreportService) ReturnBook(uid, bid int) (BookReport models.BookReportList, err error) {
	BookReport, err = brs.repo.ReturnBook(uid, bid)
	if err != nil {
		return
	}
	return
}
