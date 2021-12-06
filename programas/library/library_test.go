package library

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestBookRegisterAllParameters(t *testing.T) {
	var newBook = Book{
		Name:            "Teste",
		Price:           26.00,
		Author:          "MockAuthor",
		PublicationYear: "1945",
	}
	expectedBook := len(RegisteredBooks) + 1
	expectedOutput := fmt.Sprintf("Book {%d Teste 26 MockAuthor 1945} registered", expectedBook)
	output := RegisterBook(newBook)
	assert.Equal(t, expectedOutput, output)
}

func TestBookRegisterMissingName(t *testing.T) {
	var newBook = Book{
		Id:              4,
		Name:            "",
		Price:           26.00,
		Author:          "MockAuthor",
		PublicationYear: "1945",
	}
	expectedOutput := "Please inform book Name"
	output := RegisterBook(newBook)
	assert.Equal(t, expectedOutput, output)
}

func TestBookRegisterMissingAuthor(t *testing.T) {
	var newBook = Book{
		Id:              4,
		Name:            "MobbyDick",
		Price:           26.00,
		Author:          "",
		PublicationYear: "1945",
	}
	expectedOutput := "Please inform book Author"
	output := RegisterBook(newBook)
	assert.Equal(t, expectedOutput, output)
}

func TestBookRegisterInvalidPrice(t *testing.T) {
	var newBook = Book{
		Id:              4,
		Name:            "MobbyDick",
		Price:           -8.00,
		Author:          "MockAuthor",
		PublicationYear: "1945",
	}
	expectedOutput := "Please specify a valid price"
	output := RegisterBook(newBook)
	assert.Equal(t, expectedOutput, output)
}
func TestBookRegisterInvalidDate(t *testing.T) {
	var newBook = Book{
		Id:              4,
		Name:            "MobbyDick",
		Price:           -8.00,
		Author:          "MockAuthor",
		PublicationYear: "1945",
	}
	expectedOutput := "Please specify a valid price"
	output := RegisterBook(newBook)
	assert.Equal(t, expectedOutput, output)
}
func TestBookRegisterInvalidPublicationYear(t *testing.T) {
	var newBook = Book{
		Id:              4,
		Name:            "MobbyDick",
		Price:           8.00,
		Author:          "MockAuthor",
		PublicationYear: "1800",
	}
	expectedOutput := "Please specify a valid publication year"
	output := RegisterBook(newBook)
	assert.Equal(t, expectedOutput, output)
}

func TestGetBookExistingId(t *testing.T) {
	var newBook = Book{
		Id:              1,
		Name:            "Harry Potter",
		Price:           20.50,
		Author:          "JK Rowling",
		PublicationYear: "1995",
	}
	expectedBookOutput := newBook
	expectedBoolOutput := true
	bookOutput, found := getRequestedBook(newBook.Id)
	assert.Equal(t, expectedBookOutput, bookOutput)
	assert.Equal(t, expectedBoolOutput, found)
}

func TestGetBookNotExistingId(t *testing.T) {
	var newBook Book
	expectedBookOutput := newBook
	expectedBoolOutput := false
	bookOutput, found := getRequestedBook(15)
	assert.Equal(t, expectedBookOutput, bookOutput)
	assert.Equal(t, expectedBoolOutput, found)
}
