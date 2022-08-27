package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/rizgeek/go-bookstore/pkg/routes"
)

func main() {
	r := mux.NewRouter()
	routes.RegisterBooksRoutes(r)
	fmt.Println("Run server")
	http.Handle("/", r)
	log.Fatal(http.ListenAndServe("localhost:9010", r))
}
