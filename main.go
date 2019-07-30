package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

// Book entity
type Book struct {
	ID     uint    `json:"id"`
	Isbn   string  `json:"isbn"`
	Title  string  `json:"title"`
	Author *Author `json:"author"`
}

// Author entity
type Author struct {
	Firstname string `json:"firstname"`
	Lastname  string `json:"lastname"`
}

// Init books var as a slice Book struct
var books []Book

func generateBookID(data []Book) uint {
	var id uint = 0
	var booksID []int

	fmt.Println(data)

	for _, item := range data {
		booksID = append(booksID, int(item.ID))
	}

	max := booksID[0]

	for _, value := range booksID {
		if value > max {
			max = value
		}
	}

	id = uint(max + 1)

	return id
}

// Get All Books
func getBooks(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(books)
}

// Get a Book
func getBook(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)
	x, err := strconv.ParseUint(params["id"], 0, 64)
	if err != nil {
		fmt.Println(err)
	}
	// fmt.Println(books)
	// fmt.Println(x)
	// fmt.Printf("x is typeof %T\n", x)
	for _, item := range books {
		if item.ID == uint(x) {
			json.NewEncoder(w).Encode(item)
			return
		}
	}
	json.NewEncoder(w).Encode(&Book{})
}

// Create a book
func createBook(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var book Book
	_ = json.NewDecoder(r.Body).Decode(&book)
	book.ID = generateBookID(books)
	fmt.Println(book.ID)
	generateBookID(books)
	books = append(books, book)
	json.NewEncoder(w).Encode(book)
}

// Update a book
func updateBook(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var book Book
	_ = json.NewDecoder(r.Body).Decode(&book)
	book.ID = generateBookID(books)
	fmt.Println(book.ID)
	generateBookID(books)
	books = append(books, book)
	json.NewEncoder(w).Encode(book)
}

func main() {
	// Init Router
	r := mux.NewRouter()

	// Mock data - @todo - implement DB
	books = append(books, Book{ID: 1, Isbn: "1234", Title: "Book One", Author: &Author{Firstname: "Ridoan", Lastname: "Saleh"}})
	books = append(books, Book{ID: 2, Isbn: "1235", Title: "Book Two", Author: &Author{Firstname: "Parlindungan", Lastname: "Nasution"}})

	// Route Handlers / Endpoints
	r.HandleFunc("/api/books", getBooks).Methods("GET")
	r.HandleFunc("/api/books/{id}", getBook).Methods("GET")
	r.HandleFunc("/api/books", createBook).Methods("POST")
	r.HandleFunc("/api/books/{id}", updateBook).Methods("PUT")
	// r.HandleFunc("/api/books/{id}", deleteBook).Methods("DELETE")

	log.Fatal(http.ListenAndServe(":8000", r))
}
