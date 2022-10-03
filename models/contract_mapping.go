package models

func IssueBookMapping(br BookReport, b Book, u User) (ib IssueBook) {
	ib.BookReportID = br.BookReportID
	ib.BookName = b.BookName
	ib.Category = b.Category
	ib.IssuedBy = br.IssuedBy
	ib.IssuedTo = u.Name
	ib.Price = b.Price
	ib.Status = "issued"
	return
}

func BookReportToBookReportList(br BookReport, b Book, u User) (brl BookReportList) {

	brl.BookName = b.BookName
	brl.UserName = u.Name
	brl.BookReportID = br.BookReportID
	brl.IssueDate = br.CreatedAt
	brl.ReturnDate = br.ReturnDate
	brl.IssuedBy = br.IssuedBy
	brl.Category = b.Category
	brl.Price = b.Price
	return
}
func ReturnBookReportfunc(u User, b Book, br BookReport) (rbr BookReportList) {

	rbr.BookReportID = br.BookReportID
	rbr.BookName = b.BookName
	rbr.UserName = u.Name
	rbr.ReturnDate = br.ReturnDate
	return
}
