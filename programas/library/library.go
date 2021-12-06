package library

import (
	"fmt"
	"time"
)

type Book struct {
	Id              int     `json:"id"`
	Name            string  `json:"name"`
	Price           float64 `json:"price"`
	Author          string  `json:"author"`
	PublicationYear string  `json:"publicationYear"`
}

var RegisteredBooks = []Book{
	{
		Id:              1,
		Name:            "Harry Potter",
		Price:           20.50,
		Author:          "JK Rowling",
		PublicationYear: "1995",
	},
	{
		Id:              2,
		Name:            "Hannibal",
		Price:           45.00,
		Author:          "Thomas Harris",
		PublicationYear: "1989",
	},
	{
		Id:              3,
		Name:            "Letters from an Astrophysicist",
		Price:           13.00,
		Author:          "Neil deGrasse Tyson",
		PublicationYear: "2021",
	},
	{
		Id:              4,
		Name:            "Brumas de Avalon",
		Price:           32.10,
		Author:          "Marion Zimmer",
		PublicationYear: "1915",
	},
}

func RegisterBook(newBook Book) string {
	isNewBookValid, registerMessage := checkBookParameters(newBook)
	if !isNewBookValid {
		return registerMessage
	}
	newBook.Id = bookIdGenerator(newBook)
	RegisteredBooks = append(RegisteredBooks, newBook)
	return fmt.Sprintf("Book %v registered", newBook)
}

func checkBookParameters(newBook Book) (bool, string) {

	publicationYearLayout := "2006"
	publicationYearDate, _ := time.Parse(publicationYearLayout, newBook.PublicationYear)
	switch {
	case newBook.Name == "":
		return false, "Please inform book Name"
	case newBook.Author == "":
		return false, "Please inform book Author"
	case newBook.Price < 0.0:
		return false, "Please specify a valid price"
	case publicationYearDate.Year() <= 1900:
		return false, "Please specify a valid publication year"
	default:
		return true, ""
	}
}

func getRequestedBook(id int) (Book, bool) {

	var book Book
	// Forma menos eficiente pois copia cada elemento visitado do slice
	// for _, book := range RegisteredBooks {
	// 	if book.Id == id {
	// 		return book, true
	// 	}
	// }

	//Um pouco mais enficiente, porem criar um map talvez seja melhor
	for index := range RegisteredBooks {
		if RegisteredBooks[index].Id == id {
			return RegisteredBooks[index], true
		}
	}

	// Teria que inicialiar o mapa ao rodar o codigo
	// booksMap := map[int]Book{}
	// for _, book := range RegisteredBooks {
	// 	booksMap[book.Id] = book
	// }
	// if book, found := booksMap[id]; found {
	// 	return book, true
	// }

	return book, false

}

//Not optimal yet?
func bookIdGenerator(newBook Book) int {
	nextId := len(RegisteredBooks) + 1
	return nextId
}
