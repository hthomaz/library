package library

import (
	"fmt"
	"time"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type Book struct {
	ID              int     `json:"id" gorm:"column:bookid"`
	Name            string  `json:"name" gorm:"column:bookname"`
	Price           float64 `json:"price"`
	Author          string  `json:"author"`
	PublicationYear string  `json:"publicationYear" gorm:"column:publicationyear"`
}

func RegisterBook(newBook Book, db *gorm.DB) string {
	isNewBookValid, registerMessage := checkBookParameters(newBook)
	if !isNewBookValid {
		return registerMessage
	}

	db.Create(&newBook)
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

func SetEnviroment(host, user, dbName, password, dbport, databaseUrl string) *gorm.DB {

	var db *gorm.DB
	var err error

	//db connection string
	dbURI := fmt.Sprintf("host=%s user=%s dbname=%s sslmode=disable password=%s port=%s", host, user, dbName, password, dbport)
	if len(databaseUrl) != 0 {
		dbURI = databaseUrl
	}
	// opennig connectio to db
	db, err = gorm.Open(postgres.Open(dbURI))

	if err != nil {
		err.Error()
	} else {
		db.Exec("select 1")
		fmt.Printf("Conectado com sucesso!")
		return db
	}

	return nil
}
