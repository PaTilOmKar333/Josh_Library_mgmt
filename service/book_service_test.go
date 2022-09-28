package service

// func TestCreateBook(t *testing.T) {
// 	mockRepo := new(mock_l.MockRepositry)
// 	book := models.Book{BookName: "chava", AuthorName: "omkar patil", AvailableCopies: 10, BookStatusID: 1}

// 	mockRepo.On("CreateBook", book).Return(book, nil)

// 	testService := InitBookService(mockRepo)

// 	result, err := testService.CreateBook(book) //.CreateBook(&book)

// 	mockRepo.AssertExpectations(t)

// 	assert.Equal(t, 0, result.BookID)
// 	assert.Equal(t, "chava", result.BookName)
// 	assert.Equal(t, "omkar patil", result.AuthorName)
// 	assert.Equal(t, 10, result.AvailableCopies)
// 	assert.Equal(t, 1, result.BookStatusID)
// 	assert.Nil(t, err)

// }
