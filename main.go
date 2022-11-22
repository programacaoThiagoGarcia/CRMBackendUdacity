package main

import (
	"CRMBackend/Udacity/routes"
	"log"
	"net/http"
)

func main() {
	r := routes.NewRouter()
	log.Fatal(http.ListenAndServe(":8080", r))
}
