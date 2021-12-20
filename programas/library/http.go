package library

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
	"gorm.io/gorm"
)

//https://www.golangprograms.com/example-to-handle-get-and-post-request-in-golang.html

func Library(db *gorm.DB) {
	r := mux.NewRouter()
	r.HandleFunc("/books", booksHandler(db))
	r.HandleFunc("/book/{id}", bookHandler(db))
	http.Handle("/", r)
	http.ListenAndServe(":3000", nil)
}

func bookHandler(db *gorm.DB) func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		if r.Method != "GET" {
			fmt.Fprintf(w, "Sorry, only GET method is supported.")
			return
		}

		parametersMap := mux.Vars(r)
		bookId, idOk := parametersMap["id"]
		//bookId, idOk := r.URL.Query()["id"]
		bookIdInt, err := strconv.Atoi(bookId)

		if !idOk || bookIdInt <= 0 {
			fmt.Fprintf(w, "Url parameter id is missing")
			return
		}
		if err != nil {
			fmt.Fprintf(w, "Check is id is a valid int")
			return
		}

		//requestedBook, bookFound := getRequestedBook(bookIdInt)
		var requestedBook Book
		result := db.Where("bookid = ?", bookIdInt).Find(&requestedBook)

		if result.RowsAffected != 0 {
			bufferOutput, _ := json.Marshal(requestedBook)
			w.Header().Add("content-type", "application/json")
			w.Write(bufferOutput)
			return
		}

		w.WriteHeader(404)
		return

	}
}

func booksHandler(db *gorm.DB) func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		switch r.Method {
		case "GET":
			// bufferOutput, _ := json.Marshal(RegisteredBooks)
			// w.Write(bufferOutput)
			getBooks(db, w, r)

		case "POST":
			buffer, err := ioutil.ReadAll(r.Body)

			if err != nil {
				w.Header().Add("status", "400")
				return
			}

			newBook := &Book{}
			err = json.Unmarshal(buffer, newBook)

			if err != nil {
				w.Header().Add("status", "400")
				return
			}

			response := RegisterBook(*newBook, db)
			w.Write([]byte(response))

		default:
			fmt.Fprintf(w, "Sorry, only GET and POST methods are supported.")
		}
	}
}

func getBooks(db *gorm.DB, w http.ResponseWriter, r *http.Request) {

	var books []Book

	db.Find(&books)

	var output = books
	bufferOutput, _ := json.Marshal(output)
	w.Header().Add("content-type", "application/json")
	w.Write(bufferOutput)

	return
}
