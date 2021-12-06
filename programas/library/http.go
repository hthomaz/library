package library

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"strconv"
)

//https://www.golangprograms.com/example-to-handle-get-and-post-request-in-golang.html

func Library() {
	http.HandleFunc("/books", booksHandler)
	http.HandleFunc("/book", bookHandler)
	http.ListenAndServe(":3000", nil)
}

func bookHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != "GET" {
		fmt.Fprintf(w, "Sorry, only GET method is supported.")
		return
	}

	bookId, idOk := r.URL.Query()["id"]
	bookIdInt, err := strconv.Atoi(bookId[0])

	if !idOk || bookIdInt <= 0 {
		fmt.Fprintf(w, "Url parameter id is missing")
		return
	}
	if err != nil {
		fmt.Fprintf(w, "Check is id is a valid int")
		return
	}

	requestedBook, bookFound := getRequestedBook(bookIdInt)

	if bookFound {
		bufferOutput, _ := json.Marshal(requestedBook)
		w.Write(bufferOutput)
		return
	}

	fmt.Fprintf(w, "Book not found")
	return

}

func booksHandler(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case "GET":
		bufferOutput, _ := json.Marshal(RegisteredBooks)
		w.Write(bufferOutput)
		//getBooks(&w, r)

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

		response := RegisterBook(*newBook)
		w.Write([]byte(response))

	default:
		fmt.Fprintf(w, "Sorry, only GET and POST methods are supported.")
	}

}

/* Como criar esse tipo de funcao??
func getBooks(w *http.ResponseWriter, r *http.Request) {

	var output = RegisteredBooks
	bufferOutput, _ := json.Marshal(output)
	w.Write(bufferOutput)

	return
}
*/
