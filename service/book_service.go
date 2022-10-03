//struct and methods

package service

import (
	"context"
	"project/models"
	"project/repo"
)

type BookServiceInterface interface {
	ListBooks() (bookslist []models.Book, err error)
	CreateBook(ctx context.Context, book models.Book) (createdBook models.Book, err error)
	DeleteBook(bid int) (id int, bookReportLists []models.BookReportList, err error)
	UpdateBookWithID(ctx context.Context, bid int, book models.Book) (updateBook models.Book, err error)
}

type bookService struct {
	repo repo.BookRepoInterface
}

func InitBookService(r repo.BookRepoInterface) BookServiceInterface {

	//initialies
	//repo.InitUserRepo()
	return &bookService{
		repo: r,
	}
}

func (bs *bookService) ListBooks() (bookslist []models.Book, err error) {
	bookslist, err = bs.repo.ListBooks()
	//fmt.Println("service layer: ", users)
	if err != nil {
		return
	}
	return
}

func (bs *bookService) CreateBook(ctx context.Context, book models.Book) (createdBook models.Book, err error) {
	val, _ := ctx.Value("ClaimsToVerify").(*models.Claims)
	createdBy := val.Email
	updatedBy := val.Email
	createdBook, err = bs.repo.CreateBook(book, createdBy, updatedBy)
	if err != nil {
		return
	}
	switch createdBook.BookStatus {
	case "1":
		createdBook.BookStatus = "Available"
	case "2":
		createdBook.BookStatus = "Not available"
	}
	return
}

func (bs *bookService) DeleteBook(bid int) (id int, bookReportLists []models.BookReportList, err error) {
	id, bookReportLists, err = bs.repo.DeleteBook(bid)
	if err != nil {
		return
	}
	return
}

func (bs *bookService) UpdateBookWithID(ctx context.Context, bid int, book models.Book) (updateBook models.Book, err error) {
	val, _ := ctx.Value("ClaimsToVerify").(*models.Claims)
	updatedBy := val.Email
	updateBook, err = bs.repo.UpdateBookWithID(bid, book, updatedBy)
	if err != nil {
		return
	}
	return
}

//(id int, bookReportLists []models.BookReportList, err error)
