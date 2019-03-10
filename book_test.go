package api

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestBookToJSON(t *testing.T) {
	book := Book{Title: "Cloud Native Go", Author: "Tanver Hasan", ISBN: "SSSS"}
	json := book.ToJSON()

	assert.Equal(t, `{"title":"Cloud Native Go","author":"Tanver Hasan","isbn":"SSSS"}`,
		string(json), "Book json marshaling went wrong")
}

func TestBookFromJson(t *testing.T) {
	json := []byte(`{"title":"Cloud Native Go","author":"Tanver Hasan","isbn":"SSSS"}`)
	book := FromJSON(json)
	
	assert.Equal(t, Book{Title:"Cloud Native Go", Author: "Tanver Hasan", ISBN: "SSSS"}, book, "Unmarshaling wen wrong")
}
