package main

import (
	// "fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
	// "github.com/ginzhu/gorm/dialects/mysql"

	"github.com/LidoHon/goBookStore/pkg/routes"
	

)


func main(){
	r := mux.NewRouter()
	routes.RegisterBookStoreRoutes(r)
	http.Handle("/", r)
	log.Fatal(http.ListenAndServe("localhost:9010", r))
	


}