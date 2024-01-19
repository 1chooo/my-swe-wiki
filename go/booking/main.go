

package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"sync"

	"github.com/gorilla/mux"
)

type Book struct {
	ID     string `json:"id"`
	Title  string `json:"title"`
	Author string `json:"author"`
}

var Library map[string]Book
var mu sync.Mutex

func init() {
	Library = make(map[string]Book)
	Library["1"] = Book{ID: "1", Title: "The Catcher in the Rye", Author: "J.D. Salinger"}
	Library["2"] = Book{ID: "2", Title: "To Kill a Mockingbird", Author: "Harper Lee"}
}

func main() {
	r := mux.NewRouter()

	r.HandleFunc("/books", GetBooks).Methods("GET")
	r.HandleFunc("/books/{id}", GetBook).Methods("GET")
	r.HandleFunc("/books", AddBook).Methods("POST")
	r.HandleFunc("/books/{id}", RemoveBook).Methods("DELETE")

	// 設定靜態檔案伺服器
	r.PathPrefix("/static/").Handler(http.StripPrefix("/static/", http.FileServer(http.Dir("./static"))))

	// 如果使用者訪問根目錄，重新導向到靜態檔案的 index.html
	r.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		http.ServeFile(w, r, "./static/index.html")
	}).Methods("GET")

	http.Handle("/", r)
	fmt.Println("Server is running on port 8080...")
	http.ListenAndServe(":8080", nil)
}

func GetBooks(w http.ResponseWriter, r *http.Request) {
	mu.Lock()
	defer mu.Unlock()

	books := make([]Book, 0, len(Library))
	for _, v := range Library {
		books = append(books, v)
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(books)
}

func GetBook(w http.ResponseWriter, r *http.Request) {
	mu.Lock()
	defer mu.Unlock()

	params := mux.Vars(r)
	book, found := Library[params["id"]]
	if !found {
		http.NotFound(w, r)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(book)
}

func AddBook(w http.ResponseWriter, r *http.Request) {
	mu.Lock()
	defer mu.Unlock()

	var book Book
	err := json.NewDecoder(r.Body).Decode(&book)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	Library[book.ID] = book

	w.WriteHeader(http.StatusCreated)
}

func RemoveBook(w http.ResponseWriter, r *http.Request) {
	mu.Lock()
	defer mu.Unlock()

	params := mux.Vars(r)
	_, found := Library[params["id"]]
	if !found {
		http.NotFound(w, r)
		return
	}

	delete(Library, params["id"])

	w.WriteHeader(http.StatusNoContent)
}
