package routes

import (
	"github.com/gorilla/mux"
	"github.com/mtrosin/go-bookstore/pkg/controllers"
)

var RegisterBookStoreRoutes = func(router *mux.Router) {
	Router.HandleFunc("/book/", controllers.CreateBook).Methods("POST")
	Router.HandleFunc("/book/", controllers.GetBook).Methods("GET")
	Router.HandleFunc("/book/{bookId}", controllers.GetBookById).Methods("GET")
	Router.HandleFunc("/book/{bookId}", controllers.UpdateBook).Methods("PUT")
	Router.HandleFunc("/book/{bookId}", controllers.DeleteBook).Methods("DELETE")
}