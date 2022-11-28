package main

import (
	"CRMBackend/Udacity/routes"
	"log"
	"net/http"
)

func main() {
	r := routes.NewRouter()
	r.PathPrefix("/").Handler(http.StripPrefix("/", http.FileServer(http.Dir("./static"))))
	log.Fatal(http.ListenAndServe(":8080", r))
}
