package models

func BookToBookList(b Book, bs BookStatus) (bl BookList) {

	bl.BookID = b.BookID
	bl.BookName = b.BookName
	bl.AuthorName = b.AuthorName
	bl.Status = bs.Status
	bl.AvailableCopies = b.AvailableCopies
	return
}

func IssueBookMapping(br BookReport, b Book, u User) (ib IssueBook) {
	ib.BookReportID = br.BookReportID
	ib.BookName = b.BookName
	ib.Category = b.Category
	ib.IssuedBy = br.IssuedBy
	ib.IssuedTo = u.FirstName
	ib.Price = b.Price
	ib.Status = "issued"
	ib.ReturnDate = &br.ReturnDate
	return
}

func UserToUserList(u User, ur UserRole) (ul UserList) {
	ul.User_ID = u.User_ID
	ul.FirstName = u.FirstName
	ul.LastName = u.LastName
	ul.Email = u.Email
	ul.Age = u.Age
	ul.Address = u.Address
	ul.Role = ur.Role
	ul.CreatedBy = u.CreatedBy
	ul.CreatedAT = u.CreatedAT
	ul.UpdatedBy = u.UpdatedBy
	ul.UpdatedAt = u.UpdatedAt
	return
}

func UpdateBookMapping(ub Book) (nb UpdateBook) {

	nb.AuthorName = ub.AuthorName
	nb.AvailableCopies = ub.AvailableCopies
	nb.BookID = ub.BookID
	nb.BookName = ub.BookName
	if nb.AvailableCopies > 0 {
		nb.BookStatus = "available"
	} else {
		nb.BookStatus = "unavailable"
	}
	nb.Category = ub.Category
	nb.CreatedAT = ub.CreatedAT
	nb.CreatedBy = ub.CreatedBy
	nb.Price = ub.Price
	nb.UpdatedAt = ub.UpdatedAt
	nb.UpdatedBy = ub.UpdatedBy
	return
}

func UserToUserAuth(u User, ur UserRole) (ul UserAuth) {
	ul.User_ID = u.User_ID
	ul.FirstName = u.FirstName
	ul.LastName = u.LastName
	ul.Email = u.Email
	ul.Age = u.Age
	ul.Address = u.Address
	ul.Password = u.Password
	ul.Role = ur.Role
	return
}

func BookReportToBookReportList(br BookReport, b Book, u User) (brl BookReportList) {

	actualReturnDate := br.ActualReturnDate

	//	actualReturnDate.IsZero)

	//	actualReturnDate.String()

	brl.BookName = b.BookName
	brl.UserName = u.FirstName
	brl.BookReportID = br.BookReportID
	brl.IssueDate = br.IssueDate
	brl.ReturnDate = br.ReturnDate
	brl.IssuedBy = br.IssuedBy
	brl.Category = b.Category
	brl.Price = b.Price

	//	issuedate := br.IssueDate

	// fmt.Println(issuedate)
	// if issuedate.Valid {
	// 	brl.IssueDate = issuedate.Time
	// 	fmt.Println(brl.IssueDate)
	// }
	// returndate := br.ReturnDate
	// if returndate.Valid {
	// 	brl.ReturnDate = returndate.Time
	// }
	if actualReturnDate != nil {
		brl.BookStatus = "Book Returned"
		brl.ActualReturnDate = *actualReturnDate
	} else {
		brl.BookStatus = "Book Issued"
	}

	return
}
func ReturnBookReportfunc(u User, b Book, br BookReport) (rbr BookReportList) {
	//actualReturnDate := br.ActualReturnDate

	actualReturnDate := br.ActualReturnDate

	rbr.BookReportID = br.BookReportID
	rbr.BookName = b.BookName
	rbr.UserName = u.FirstName

	// issuedate := br.IssueDate
	// if issuedate.Valid {
	// 	rbr.IssueDate = issuedate.Time
	// 	fmt.Println(rbr.IssueDate)
	// }
	// returndate := br.ReturnDate
	// if returndate.Valid {
	// 	rbr.ReturnDate = returndate.Time
	// }
	// if actualReturnDate.Valid {
	// 	rbr.BookStatus = "Book Returned"
	// 	rbr.ActualReturnDate = *&actualReturnDate.Time
	// } else {
	// 	rbr.BookStatus = "Book Issued"
	// }

	rbr.IssueDate = br.IssueDate
	rbr.ReturnDate = br.ReturnDate
	if actualReturnDate != nil {
		rbr.ActualReturnDate = *br.ActualReturnDate
	}

	return
}

// func UpdatedUserDetails(ou User, nu User) (uu UpdateUser) {
// 	if ou.FirstName == nu.FirstName && ou.LastName == nu.LastName {
// 		fmt.Println("ou.FirstName == nu.FirstName && ou.LastName == nu.LastName")
// 		uu.OldPassword = ou.Password
// 		uu.NewPassword = nu.Password
// 	} else if ou.FirstName == nu.FirstName && ou.Password == nu.Password {
// 		fmt.Println("ou.FirstName == nu.FirstName && ou.Password == nu.Password")
// 		uu.OldLastName = ou.LastName
// 		uu.NewLastName = nu.LastName
// 	} else if ou.LastName == nu.LastName && ou.Password == nu.Password {
// 		fmt.Println("ou.LastName == nu.LastName && ou.Password == nu.Password")
// 		uu.OldFirstName = ou.FirstName
// 		uu.NewFirstName = nu.FirstName
// 	} else if ou.FirstName == nu.FirstName {
// 		fmt.Println("ou.FirstName == nu.FirstName")
// 		uu.OldLastName = ou.LastName
// 		uu.OldPassword = ou.Password
// 		uu.NewLastName = nu.LastName
// 		uu.NewPassword = nu.Password
// 	} else if ou.LastName == nu.LastName {
// 		fmt.Println("ou.LastName == nu.LastName")
// 		uu.OldFirstName = ou.FirstName
// 		uu.OldPassword = ou.Password
// 		uu.NewFirstName = nu.FirstName
// 		uu.NewPassword = nu.Password
// 	} else if ou.Password == nu.Password {
// 		fmt.Println("ou.Password == nu.Password")
// 		uu.OldFirstName = ou.FirstName
// 		uu.OldLastName = ou.LastName
// 		uu.NewFirstName = nu.FirstName
// 		uu.NewLastName = nu.LastName
// 	} else {
// 		fmt.Println("else")
// 		uu.OldFirstName = ou.FirstName
// 		uu.OldLastName = ou.LastName
// 		uu.OldPassword = ou.Password
// 		uu.NewFirstName = nu.FirstName
// 		uu.NewLastName = nu.LastName
// 		uu.NewPassword = nu.Password
// 	}

// // uu.OldFirstName = ou.FirstName
// // uu.OldLastName = ou.LastName
// // uu.OldPassword = ou.Password
// // uu.NewFirstName = nu.FirstName
// // uu.NewLastName = nu.LastName
// // uu.NewPassword = nu.Password
// 	return
// }
