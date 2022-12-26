package routes

import (
	"bookstore/pkg/controllers" //absolute paths

	"github.com/gorilla/mux"
)

var RegisterBookStoreRoutes = func(r *mux.Router) {
	r.HandleFunc("/books/", controllers.CreateBook).Methods("POST")
	r.HandleFunc("/books/", controllers.GetBooks).Methods("GET")
	r.HandleFunc("/books/{bookId}", controllers.GetBookById).Methods("GET")
	r.HandleFunc("/books/{bookId}", controllers.UpdateBook).Methods("PUT")
	r.HandleFunc("/books/{bookId}", controllers.DeleteBook).Methods("DELETE")
}
