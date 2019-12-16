// Package handlers handles all the GET request and responses send to the webserver.
package handlers

import (
	"html/template"
	"log"
	"net/http"

	"../repository"
)

// HandRaw handles the /laptops GET request.
// In doing so it also parses the values from laptops.json file.
func HandRaw(response http.ResponseWriter, request *http.Request) {
	log.Println("Printing raw comments")
	request.ParseForm()

	comments, err := repository.LoadComments()
	if err != nil {
		log.Println(err)
		http.Redirect(response, request, "../html/error.html", http.StatusBadRequest)
		return
	}

	render, err1 := template.ParseFiles("templates/raw.html")
	if err1 != nil {
		log.Println(err1)
		http.Redirect(response, request, "../html/error.html", http.StatusBadRequest)
		return
	}

	// render.Execute sends a response to the user while parsing the laptops variables to the webpage.
	render.Execute(response, comments)
}

// type ByTag []Comment

// func (a ByTag) Len() int           { return len(a) }
// func (a ByTag) Swap(i, j int)      { a[i], a[j] = a[j], a[i] }
// func (a ByTag) Less(i, j int) bool { return a[i].Tag > a[j].Tag }
