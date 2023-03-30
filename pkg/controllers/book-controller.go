package controllers

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"

	"github.com/dev-hack95/bookstore_api_backend/pkg/models"
	"github.com/dev-hack95/bookstore_api_backend/pkg/utils"
	"github.com/gorilla/mux"
)

var NewBook models.Book

func GetBook(w http.ResponseWriter, r *http.Request) {
	newBooks := models.GetALLBooks()
	// Marshelling and Unmarshelling data
	// Marshal takes any value ( any is a wrapper around interface{} ) and converts it into a byte slice.
	res, _ := json.Marshal(newBooks)
	w.Header().Set("Content-Type", "pkglication/json")
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

func GetBookById(w http.ResponseWriter, r *http.Request) {
	// vars return the route varibles
	bookId := mux.Vars(r)["id"]
	ID, err := strconv.ParseInt(bookId, 0, 0)
	if err != nil {
		fmt.Println(err)
	}
	bookDetails, _ := models.GetBookById(ID)
	res, _ := json.Marshal(bookDetails)
	w.Header().Set("Content-Type", "pkglication/json")
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

func CreateBook(w http.ResponseWriter, r *http.Request) {
	CreateBook := &models.Book{}
	// Send data to decode/Unmarshell to utils
	utils.ParseBody(r, CreateBook)
	// ParseBody(r *http.Request, x interface{})
	// CreateBook interface{} // body is stored into interface until processed is finshed
	// Reads body and convert it into json
	b := CreateBook.CreateBook()
	res, _ := json.Marshal(b)
	w.Header().Set("Content-Type", "pkglication/json")
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

func DeleteBook(w http.ResponseWriter, r *http.Request) {
	bookId := mux.Vars(r)["id"]
	ID, err := strconv.ParseInt(bookId, 0, 0)
	if err != nil {
		fmt.Println("error while parsing")
	}
	deleteBook := models.DeleteBook(ID)
	// Marshelling and Unmarshelling data
	// Marshal takes any value ( any is a wrapper around interface{} ) and converts it into a byte slice.
	res, _ := json.Marshal(deleteBook)
	w.Header().Set("Content-Type", "pkglication/json")
	w.WriteHeader(http.StatusOK)
	w.Write(res)

}

func UpdateBook(w http.ResponseWriter, r *http.Request) {
	// Create a varible updateBook which is directing to schema
	var updateBook = &models.Book{}
	// Send data to decode/Unmarshell to utils
	utils.ParseBody(r, updateBook)
	// ParseBody(r *http.Request, x interface{})
	// updateBook interface{}
	// Reads body and convert it into json
	bookId := mux.Vars(r)["id"]
	ID, err := strconv.ParseInt(bookId, 0, 0)
	if err != nil {
		fmt.Println(err)
	}
	bookDetails, db := models.GetBookById(ID)

	if updateBook.Name != "" {
		bookDetails.Name = updateBook.Name
	}

	if updateBook.Author != "" {
		bookDetails.Author = updateBook.Author
	}

	if updateBook.Publication != "" {
		bookDetails.Publication = updateBook.Publication
	}
	db.Save(&bookDetails)
	res, _ := json.Marshal(bookDetails)
	w.Header().Set("Content-Type", "pkglication/json")
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}
