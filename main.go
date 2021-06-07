package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/arpesam/go-book-api/src/routes"
	"github.com/gorilla/mux"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

func main() {
	r := mux.NewRouter()
	fmt.Println(r)
	routes.RegisterBookStoreRoutes(r)
	http.Handle("/", r)
	fmt.Println("Runnning server")
	log.Fatal(http.ListenAndServe("localhost:8080", r))
}
