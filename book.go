package api

import (
	"encoding/json"
	"net/http"
)

// Book type
type Book struct {
	Title  string `json:"title"`
	Author string `json:"author"`
	ISBN   string `json:"isbn"`
}

// ToJSON method
func (b Book) ToJSON() []byte {
	toJSON, err := json.Marshal(b)
	if err != nil {
		panic(err)
	}
	return toJSON
}

// FromJSON method
func FromJSON(data []byte) Book {
	book := Book{}
	err := json.Unmarshal(data, &book)
	if err != nil {
		panic(err)
	}
	return book
}

var books = []Book{
	Book{Title: "Himu", Author: "Humayan ahmed", ISBN: "555"},
	Book{Title: "Bidrohi", Author: "Kazi Nazrul", ISBN: "55S5"},
}
// BooksHandlerFunc method
func BooksHandlerFunc(w http.ResponseWriter, r *http.Request) {
	b, err := json.Marshal(books)
	if err !=nil{
		panic(err)
	}
	w.Header().Add("Content-Type","application/json;")
	w.Write(b)
}
