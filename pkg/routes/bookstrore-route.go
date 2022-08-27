package routes

import (
	"github.com/gorilla/mux"
	"github.com/rizgeek/go-bookstore/pkg/controllers"
)

var RegisterBooksRoutes = func(router *mux.Router) {
	router.Path("/book/").HandlerFunc(controllers.CreateBook).Methods("POST")
	router.Path("/book/").HandlerFunc(controllers.GetBook).Methods("GET")
	router.Path("/book/{bookId}").HandlerFunc(controllers.GetBookById).Methods("GET")
	router.Path("/book/{bookId}").HandlerFunc(controllers.UpdateBook).Methods("PUT")
	router.Path("/book/{bookId}").HandlerFunc(controllers.DeleteBook).Methods("DELETE")
}
