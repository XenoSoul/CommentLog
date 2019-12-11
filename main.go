// Package main is the base of your local server.
package main

import (
	"log"
	"net/http"

	"./handlers"
)

func main() {
	log.Println("Starting app")

	// http.HandleFunc listens to a GET request and handles it.
	http.HandleFunc("/delete", handlers.HandDelete)
	http.HandleFunc("/", handlers.HandComment)

	// http.Handle allow the use of files from the spesified destination within your code.
	http.Handle("/html/", http.StripPrefix("/html/", http.FileServer(http.Dir("./html"))))
	http.Handle("/images/", http.StripPrefix("/images/", http.FileServer(http.Dir("./images"))))

	// http.ListenAndServe listens to get requests send to port 8080.
	log.Fatal(http.ListenAndServe(":8080", nil))
}

// http://localhost:8080/laptops
// godoc -http=:6060
